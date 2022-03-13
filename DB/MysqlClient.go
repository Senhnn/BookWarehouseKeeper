package DB

import (
	"BWKV1/Model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var mysqlClient *Mysql

var mysqlInitOnce sync.Once

type Mysql struct {
	mysqlDb *gorm.DB
}

func (m *Mysql) MysqlDB() *gorm.DB {
	return m.mysqlDb
}

func MysqlClient() *Mysql {
	return mysqlClient
}

func MysqlInit() {
	mysqlInitOnce.Do(func() {
		mysqlClientInit()
		mysqlCreate()
	})
}

func mysqlCreate() {
	// 迁移
	MysqlClient().MysqlDB().AutoMigrate(&Model.BookNum{}, &Model.Book{})
	// 设置连接池的最大数量
	db, err := MysqlClient().mysqlDb.DB()
	if err != nil {
		panic(any(fmt.Errorf("Get DB err, %w \n", err.Error())))
	}
	db.SetMaxIdleConns(10)           // 连接池中最大的空闲连接数
	db.SetMaxOpenConns(100)          // 连接池中最多容纳的连接数量
	db.SetConnMaxLifetime(time.Hour) // 连接池中连接的最大可复用时间
}

func mysqlClientInit() {
	res := viper.GetStringMapString("mysql")
	var ip string = res["ip"]
	var port string = res["port"]
	var usr string = res["user"]
	var passwd string = res["password"]
	var dbname string = res["dbname"]
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", usr, passwd, ip, port, dbname)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               mysqlDsn,
		DefaultStringSize: 1024,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 前缀
			SingularTable: true, // 单数表名
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 跳过外键约束，尽量在代码里使用逻辑外键
	})
	if err != nil {
		panic(any("failed to connect mysql database"))
	}
	fmt.Println("Open mysql successful")
	mysqlClient = &Mysql{
		mysqlDb: db,
	}
}
