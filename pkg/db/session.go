package db

import (
	"fmt"
	"sync"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/shettyh/contacts-book/pkg/config"
	"github.com/shettyh/contacts-book/pkg/db/model"
)

// Supported DB types
const (
	dbTypeMysql  = "mysql"
	dbTypeSqlite = "sqlite3"
)

// Errors
const (
	errUnsupportedDb = "unsupported database type: %s"
)

// dbInstance singleton holder for Session
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
		dbType, connectionString, err := getConnectionDetails()
		if err != nil {
			log.Fatalf("failed to form the db connection string, %v", err)
		}
		// Connect to MySQL server
		db, err := gorm.Open(dbType, connectionString)
		if err != nil {
			log.Fatalf("failed to open database connection, %s", err.Error())
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

// getConnectionDetails will return the database type and connection string.
// If the provided DB type is not supported then it will return error.
func getConnectionDetails() (dbType, connectionString string, err error) {
	conf := config.GetInstance()
	switch conf.DbType {
	case dbTypeMysql:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
		dbType = dbTypeMysql
		return
	case dbTypeSqlite:
		connectionString = conf.DbName // Path for the sqlite data file
		dbType = dbTypeSqlite
		return
	default:
		err = fmt.Errorf(errUnsupportedDb, conf.DbType)
		return
	}
}
