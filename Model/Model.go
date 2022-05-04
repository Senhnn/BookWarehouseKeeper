package Model

import "gorm.io/gorm"

//type AuthorInfo struct {
//	gorm.Model
//	// 作家姓名
//	Name string
//
//	// 国籍
//	Nationality   Nationality `gorm:"foreignKey:CountryName"`
//	NationalityId uint
//
//	// 作品
//	Books []Book
//}

//type Nationality struct {
//	gorm.Model
//	// 国家名称
//	CountryName string
//	// 作家
//	Authors []AuthorInfo
//}

//type BookType struct {
//	gorm.Model
//	BookType string
//
//	// 书籍，多对多关系
//	Books []Book `gorm:"many2many:book_booktype;"`
//}

type Book struct {
	gorm.Model
	// isbn号码
	Isbn string `gorm:"not null"`
	// 书名
	BookName string `gorm:"not null"`

	AuthorName string

	CountryName string

	// 书籍类型
	BookType string

	// 书籍数量以及售卖状态
	Num    uint32 `gorm:"not null"`
	Status uint8  `gorm:"not null"`
}

//type BookNum struct {
//	gorm.Model
//	Num    uint32 `gorm:"not null"`
//	Status uint8  `gorm:"not null"`
//}
