package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Email string `gorm:"type:varchar(255);primary_key"`
	Name string `gorm:"not null"`
	Phone string
	Password string
}

func main() {
	//db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("mysql", "root@/contactsbook?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	db.Create(&User{Email:"shetty@live.com", Name: "Shetty", Phone: "8970820090" })


}
