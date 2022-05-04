package RpcService

import (
	"BWKV1/DB"
	"BWKV1/MiddleWare"
	"BWKV1/Model"
	"BWKV1/Services"
	BWKErr "BWKV1/error"
	"context"
	"encoding/json"
	"gorm.io/gorm"
)

type WorkerServer struct {
	Services.UnimplementedWorkerServiceServer
}

func (w WorkerServer) GetBooksByAuthor(ctx context.Context, req *Services.GetBooksByAuthorRequest) (rsp *Services.GetBooksByAuthorResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	name := req.AuthorName
	if name != "" {
		rsp.Ret = BWKErr.PARAM_INVALID
		return
	}

	var bookList []Model.Book
	err = DB.MysqlClient().MysqlDB().Model(&Model.Book{}).Where("author_name = ?", name).Find(&bookList).Error
	if err != nil {
		MiddleWare.GLog.Errorf("Mysql get author:%s book err", name)
		rsp.Ret = BWKErr.MYSQL_GET_AUTHOR_BOOKS_FAIL
		return
	}
	if len(bookList) == 0 {
		MiddleWare.GLog.Debugf("Mysql get author:%s book succ, the author has no book", name)
		rsp.Ret = BWKErr.MYSQL_GET_AUTHOR_BOOKS_FAIL
		return
	}
	rsp.AuthorInfo.Name = name
	rsp.AuthorInfo.Nation = bookList[0].CountryName
	rsp.BookInfo = []*Services.BookInfo{}
	for _, v := range bookList {
		rsp.BookInfo = append(rsp.BookInfo, &Services.BookInfo{
			Isbn:     v.Isbn,
			Author:   v.AuthorName,
			Nation:   v.CountryName,
			BookName: v.BookName,
			BookNum:  uint64(v.Num),
			Status:   Services.BookSaleStatus(v.Status),
			BookType: v.BookType,
		})
	}
	return
}

func (w WorkerServer) GetAllBook(ctx context.Context, req *Services.GetAllBookRequest) (rsp *Services.GetAllBookResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	rsp = &Services.GetAllBookResponse{}
	bookList := &[]Model.Book{}
	if err = DB.MysqlClient().MysqlDB().Model(&Model.Book{}).Find(bookList).Error; err != nil {
		rsp.Ret = BWKErr.MYSQL_ERR_GET_ALL_BOOK
		return
	}

	// 填充书籍信息
	rsp.BookList = make([]*Services.BookInfo, 0)
	for _, v := range *bookList {
		rsp.BookList = append(rsp.BookList, &Services.BookInfo{
			Isbn:     v.Isbn,
			Author:   v.AuthorName,
			Nation:   v.CountryName,
			BookName: v.BookName,
			BookNum:  uint64(v.Num),
			Status:   Services.BookSaleStatus(v.Status),
		})
	}
	rsp.Ret = BWKErr.SUCCESS
	return
}

func (w WorkerServer) GetBookNum(ctx context.Context, req *Services.GetBookNumRequest) (rsp *Services.GetBookNumResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	isbn := req.Isbn
	book := &Model.Book{}

	// 先从redis中查找
	ret, val := DB.RedisClient().GetBookCache(isbn)
	// redis中存在
	if ret == BWKErr.SUCCESS {
		err = json.Unmarshal([]byte(val), book)
		if err != nil {
			MiddleWare.GLog.Errorf("JSON unmarshal err:%s", err.Error())
			rsp.Ret = BWKErr.JSON_UNMARSHAL_ERROR
			return
		} else {
			rsp.Num = book.Num
			rsp.Ret = BWKErr.SUCCESS
			err = nil
		}
		return
	}

	// redis中不存在，在mysql中查找
	if err = DB.MysqlClient().MysqlDB().Model(&Model.Book{}).Where("isbn = ?", isbn).Take(book).Error; err != nil {
		MiddleWare.GLog.Errorf("Mysql book isbn:%s found fail", isbn)
		rsp.Ret = BWKErr.MYSQL_ERR_BOOK_NOT_FOUND
		return
	}

	// 添加缓存
	go func() {
		// 加入缓存
		ret = DB.RedisClient().SetBookCache(book)
		if ret != 0 {
			MiddleWare.GLog.Errorf("Redis set book:%s fail", isbn)
		}
		return
	}()

	rsp.Ret = BWKErr.SUCCESS
	rsp.Num = book.Num
	return
}

func (w WorkerServer) BuyBook(ctx context.Context, req *Services.BuyBookRequest) (rsp *Services.BuyBookResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	bookList := req.BookList
	for _, data := range bookList {
		isbn := data.Isbn
		buyNum := data.BuyNum
		book := &Model.Book{}
		ret, str := DB.RedisClient().GetBookCache(isbn)
		if ret != BWKErr.SUCCESS {
			// redis中没有从mysql数据库中查找
			err = DB.MysqlClient().MysqlDB().Model(&Model.Book{}).Where("isbn = ?", isbn).Find(book).Error
			if err != nil {
				MiddleWare.GLog.Errorf("mysql get isbn:%s book fail", isbn)
				rsp.Ret = BWKErr.MYSQL_ERR_BOOK_NOT_FOUND
				return rsp, err
			}
			if book.Num < buyNum {
				MiddleWare.GLog.Errorf("mysql get isbn:%s num:%d not enough", isbn, book.Num)
				rsp.Ret = BWKErr.BOOK_NUM_NOT_ENOUGH
				return rsp, err
			}
			// 设置缓存
			DB.RedisClient().SetBookCache(book)
		} else {
			err := json.Unmarshal([]byte(str), book)
			if err != nil {
				MiddleWare.GLog.Errorf("redis get cache isbn:%s json unmarshal fail", isbn)
				rsp.Ret = BWKErr.JSON_UNMARSHAL_ERROR
				return rsp, err
			}
			if book.Num < buyNum {
				MiddleWare.GLog.Errorf("redis get cache isbn:%s book not enough", isbn)
				rsp.Ret = BWKErr.BOOK_NUM_NOT_ENOUGH
				return rsp, err
			}
		}
	}

	if err = DB.MysqlClient().MysqlDB().Transaction(func(tx *gorm.DB) error {
		for _, data := range bookList {
			book := &Model.Book{}
			isbn := data.Isbn
			buyNum := data.BuyNum
			err = DB.MysqlClient().MysqlDB().Model(&Model.Book{}).Where("isbn = ?", isbn).Find(book).Error
			if err != nil {
				return err
			}
			if book.Num < buyNum {
				return err
			}
		}
		return nil
	}); err != nil {
		MiddleWare.GLog.Errorf("mysql BuyBook tx error")
		rsp.Ret = BWKErr.MYSQL_ERR_UPDATE_BOOK
		return rsp, err
	}

	// 数据库删除成功时，清除所有相关缓存
	go func() {
		for _, data := range bookList {
			isbn := data.Isbn
			DB.RedisClient().DelBookCache(isbn)
		}
	}()
	return
}
