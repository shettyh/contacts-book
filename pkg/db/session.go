package db

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbInstance struct {
	once     sync.Once
	instance *Session
}

type Session struct {
	*gorm.DB
}

func GetSession() *Session {
	dbInstance.once.Do(func() {
		db, err := gorm.Open("mysql", "root@/contactsbook?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic(fmt.Sprintf("failed to open database connection, %s", err.Error()))
		}
		dbInstance.instance = new(Session)
		dbInstance.instance.DB = db
	})
	return dbInstance.instance
}
