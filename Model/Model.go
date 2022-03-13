package Model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	// isbn号码
	Isbn string `gorm:"isbn" gorm:"not null"`
	// 作者
	Author string `gorm:"author" gorm:"not null"`
	// 作者国籍
	Nation string `gorm:"nation" gorm:"not null"`
	// 书名
	BookName string `gorm:"bookname" gorm:"not null"`
}

type BookNum struct {
	Isbn string `gorm:"isbn"`
	Num  uint64 `gorm:"num"`
}
