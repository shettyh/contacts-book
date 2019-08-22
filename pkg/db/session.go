package db

import (
	"fmt"
	"sync"

	"github.com/shettyh/contacts-book/pkg/db/model"

	"log"

	"github.com/shettyh/contacts-book/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbType = "mysql"
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
		log.Print("Creating the database connection...")
		// Form the mysql connection URL
		conf := config.GetInstance()
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)

		// Connect to MySQL server
		db, err := gorm.Open(dbType, connectionString)
		if err != nil {
			panic(fmt.Sprintf("failed to open database connection, %s", err.Error()))
		}
		dbInstance.instance = new(Session)
		dbInstance.instance.DB = db

		log.Print("Database connection initialized successfully.")

		// Create and migrate database schema using GORM
		// Add all the models for the migration here.
		dbInstance.instance.AutoMigrate(
			&model.User{},
			model.Contact{},
		)
	})
	return dbInstance.instance
}
