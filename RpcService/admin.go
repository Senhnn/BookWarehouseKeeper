package RpcService

import (
	"BWKV1/DB"
	"BWKV1/MiddleWare"
	"BWKV1/Model"
	"BWKV1/Services"
	BWKErr "BWKV1/error"
	"context"
	"database/sql"
	"gorm.io/gorm"
	"sync"
	"sync/atomic"
	"time"
)

type AdminServer struct {
	Services.UnimplementedAdminServiceServer
}

// AddBookNum 增加书籍库存
func (a *AdminServer) AddBookNum(ctx context.Context, req *Services.AddBookNumRequest) (rsp *Services.AddBookNumResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	rsp = &Services.AddBookNumResponse{Ret: -1}
	isbn := req.GetISBN()
	addNum := req.GetAddNum()
	book := Model.Book{}

	// 先更新数据库
	// 数据库事务：增对应书籍数量
	err = DB.MysqlClient().MysqlDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Model.Book{}).Where("isbn = ?", isbn).First(&book).Error; err != nil {
			return err
		}
		if err := tx.Model(&book).Update("num", book.Num+uint32(addNum)); err != nil {
			return nil
		}
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		rsp.Ret = BWKErr.MYSQL_ERR_UPDATE_BOOK
		return
	}

	// 再删除缓存
	errCode, _ := DB.RedisClient().DelBookCache(isbn)
	if errCode == BWKErr.SUCCESS {
		rsp.Ret = errCode
		MiddleWare.GLog.Errorf("Redis del succ, errCode:%d", errCode)
		return
	}

	// 定时器，延迟双删
	t := time.NewTimer(time.Second)
	go func() {
		select {
		case <-t.C:
			DB.RedisClient().DelBookCache(isbn)
		}
		return
	}()

	rsp.Ret = BWKErr.SUCCESS
	return
}

// DecBookNum 减少书籍库存
func (a *AdminServer) DecBookNum(ctx context.Context, req *Services.DecBookRequest) (rsp *Services.DecBookResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	rsp = &Services.DecBookResponse{Ret: -1}
	isbn := req.GetISBN()
	decNum := req.GetDecNum()
	book := Model.Book{}

	// 先更新数据库
	// 数据库事务：减少应书籍数量
	err = DB.MysqlClient().MysqlDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("BookNum").Model(&Model.Book{}).Where("isbn = ?", isbn).First(&book).Error; err != nil {
			return err
		}
		if err := tx.Model(&book).Update("num", book.Num-uint32(decNum)); err != nil {
			return nil
		}
		return nil
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		rsp.Ret = BWKErr.MYSQL_ERR_UPDATE_BOOK
		return
	}

	// 再删除缓存
	errCode, _ := DB.RedisClient().DelBookCache(isbn)
	if errCode == BWKErr.SUCCESS {
		rsp.Ret = BWKErr.SUCCESS
		MiddleWare.GLog.Errorf("Redis del succ, errCode:%d", errCode)
		return
	}

	// 定时器，延迟双删
	t := time.NewTimer(time.Second)
	go func() {
		select {
		case <-t.C:
			DB.RedisClient().DelBookCache(isbn)
		}
		return
	}()

	rsp.Ret = BWKErr.SUCCESS
	return
}

// AddNewBook 添加新书
func (a *AdminServer) AddNewBook(ctx context.Context, req *Services.AddNewBookRequest) (rsp *Services.AddNewBookResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	rsp = &Services.AddNewBookResponse{Ret: -1}
	isbn := req.GetBookInfo().GetIsbn()
	bookInfo := req.GetBookInfo()

	wg := sync.WaitGroup{}
	// 书籍是否存在的标志
	var isExist uint32 = 0
	// 在数据库book表中查询书籍信息，书籍不存在。
	wg.Add(1)
	go func() {
		defer wg.Done()
		if DB.MysqlClient().MysqlDB().Where("isbn = ?", isbn).Take(&Model.Book{}).Error == nil {
			MiddleWare.GLog.Errorf("Book(isbn:%s) is exist!", isbn)
			atomic.AddUint32(&isExist, 1)
		}
	}()

	wg.Wait()
	// 此时说明书籍已经存在
	if isExist != 0 {
		rsp.Ret = BWKErr.MYSQL_ERR_BOOK_HAS_EXIST
		return
	}

	txErr := DB.MysqlClient().MysqlDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("BookNum").Create(&Model.Book{
			Model:       gorm.Model{},
			Isbn:        bookInfo.GetIsbn(),
			AuthorName:  bookInfo.Author,
			BookName:    bookInfo.GetBookName(),
			CountryName: bookInfo.Nation,
			BookType:    bookInfo.BookType,
			Num:         uint32(bookInfo.BookNum),
			Status:      uint8(Services.BookSaleStatus_CLOSE),
		}).Error; err != nil {
			MiddleWare.GLog.Errorf("Book(isbn:%s) Create fail!", isbn)
			return err
		}
		return nil
	})
	if txErr != nil {
		rsp.Ret = BWKErr.MYSQL_ERR_CREATE_BOOK_OR_BOOKNUM
	} else {
		rsp.Ret = BWKErr.SUCCESS
	}
	return
}

// RemoveBook 删除书籍
func (a *AdminServer) RemoveBook(ctx context.Context, req *Services.RemoveBookRequest) (rsp *Services.RemoveBookResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	rsp = &Services.RemoveBookResponse{Ret: -1}
	isbn := req.GetISBN()

	// Mysql事务
	if txErr := DB.MysqlClient().MysqlDB().Transaction(func(tx *gorm.DB) error {
		// mysql停止售卖
		book := &Model.Book{}
		if innerErr := tx.Model(&Model.Book{}).Where("isbn = ?", isbn).Take(book).Error; innerErr != nil {
			MiddleWare.GLog.Errorf("book(isbn:%s) not found", isbn)
			return innerErr
		}
		// 停止售卖
		book.Status = uint8(Services.BookSaleStatus_CLOSE)

		// 删除相关数据
		if innerErr := tx.Model(&Model.Book{}).Where(
			"isbn = ?", isbn).Delete(book).Error; innerErr != nil {
			MiddleWare.GLog.Errorf("book(isbn:%s) not found", isbn)
			return innerErr
		}
		return nil
	}); txErr != nil {
		rsp.Ret = BWKErr.MYSQL_ERR_REMOVE_BOOK
		return
	}

	// 删除缓存
	if redErr, _ := DB.RedisClient().DelBookCache(isbn); redErr == BWKErr.REDIS_ERR_DEL_DATA {
		rsp.Ret = BWKErr.REDIS_ERR_DEL_DATA
		return
	}

	// 延迟删除redis
	time.AfterFunc(time.Second, func() {
		DB.RedisClient().DelBookCache(isbn)
	})
	return
}

// StartOrCloseBookSale 开启或者关闭售卖
func (a *AdminServer) StartOrCloseBookSale(ctx context.Context, req *Services.BookSaleRequest) (rsp *Services.BookSaleResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	rsp = &Services.BookSaleResponse{Ret: -1}

	// 修改mysql的bookNum表状态
	// 更新数据
	if err1 := DB.MysqlClient().MysqlDB().Preload("BookNum").Model(
		&Model.Book{}).Where("isbn = ?", req.Isbn).Update("status", req.SetStatus).Error; err1 != nil {
		rsp.Ret = BWKErr.MYSQL_ERR_UPDATE_BOOK
		return
	}

	// 删除redis缓存
	errCode, _ := DB.RedisClient().DelBookCache(req.GetIsbn())
	if errCode == BWKErr.SUCCESS {
		rsp.Ret = BWKErr.SUCCESS
		return
	}
	// 延迟双删
	time.AfterFunc(time.Second, func() {
		DB.RedisClient().DelBookCache(req.GetIsbn())
	})
	return
}
