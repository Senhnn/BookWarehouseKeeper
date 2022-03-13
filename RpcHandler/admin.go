package RpcHandler

import (
	"BWKV1/DB"
	"BWKV1/Model"
	"BWKV1/Services"
	BWKErr "BWKV1/error"
	"database/sql"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"time"
)

type AdminServices struct{}

// AddBookNum 增加书籍库存
func (a *AdminServices) AddBookNum(ctx context.Context, req *Services.AddBookNumRequest,
	opts ...grpc.CallOption) (rsp *Services.AddBookNumResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	isbn := req.GetISBN()
	addNum := req.GetAddNum()
	bookNum := Model.BookNum{}

	// 先更新数据库
	// 数据库事务：增对应书籍数量
	err = DB.MysqlClient().MysqlDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("isbn = ?", isbn).First(&bookNum).Error; err != nil {
			return err
		}
		if err := tx.Model(&bookNum).Where("isbn = ?", isbn).Update("num", bookNum.Num+addNum); err != nil {
			return nil
		}
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		rsp.Ret = BWKErr.UPDATE_MYSQL_ERR
		return
	}

	// 再删除缓存
	ret, err := DB.RedisClient().HDel(context.Background(), DB.BookNumHash, isbn).Result()
	if err != nil {
		rsp.Ret = BWKErr.SUCCESS
		fmt.Println(ret)
		return
	}

	// 定时器，延迟双删
	t := time.NewTimer(time.Second)
	go func() {
		select {
		case <-t.C:
			DB.RedisClient().HDel(context.Background(), DB.BookNumHash, isbn)
		}
		return
	}()

	rsp.Ret = BWKErr.SUCCESS
	return
}

// DecBookNum 减少书籍库存
func (a *AdminServices) DecBookNum(ctx context.Context, req *Services.DecBookRequest,
	opts ...grpc.CallOption) (rsp *Services.DecBookResponse, err error) {
	defer func() { rsp.ErrorInfo = BWKErr.ErrMap[rsp.Ret] }()
	isbn := req.GetISBN()
	decNum := req.GetDecNum()
	bookNum := Model.BookNum{}

	// 先更新数据库
	// 数据库事务：减少应书籍数量
	err = DB.MysqlClient().MysqlDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("isbn = ?", isbn).First(&bookNum).Error; err != nil {
			return err
		}
		if err := tx.Model(&bookNum).Where("isbn = ?", isbn).Update("num", bookNum.Num-decNum); err != nil {
			return nil
		}
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		rsp.Ret = BWKErr.UPDATE_MYSQL_ERR
		return
	}

	// 再删除缓存
	ret, err := DB.RedisClient().HDel(context.Background(), DB.BookNumHash, isbn).Result()
	if err != nil {
		rsp.Ret = BWKErr.SUCCESS
		fmt.Println(ret)
		return
	}

	// 定时器，延迟双删
	t := time.NewTimer(time.Second)
	go func() {
		select {
		case <-t.C:
			DB.RedisClient().HDel(context.Background(), DB.BookNumHash, isbn)
		}
		return
	}()

	rsp.Ret = BWKErr.SUCCESS
	return
}

// AddNewBook 上架新书
func (a *AdminServices) AddNewBook(ctx context.Context, req *Services.AddNewBookRequest,
	opts ...grpc.CallOption) (rsp *Services.AddNewBookResponse, err error) {

}

// DelBook 下架书籍
func (a *AdminServices) DelBook(ctx context.Context, req *Services.DelBookRequest,
	opts ...grpc.CallOption) (rsp *Services.DelBookResponse, err error) {

}

// RemoveBook 删除书籍
func (a *AdminServices) RemoveBook(ctx context.Context, req *Services.RemoveBookRequest,
	opts ...grpc.CallOption) (rsp *Services.RemoveBookResponse, err error) {

}
