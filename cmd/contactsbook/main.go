package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/shettyh/contacts-book/pkg/db"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shettyh/contacts-book/pkg/db/model"
)

//type User struct {
//	Email    string `gorm:"type:varchar(255);primary_key"`
//	Name     string `gorm:"not null"`
//	Phone    string
//	Password string
//}

func main() {
	//db, err := gorm.Open("sqlite3", "test.db")
	//db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//db.AutoMigrate(&model.User{}, &model.Contact{})

	//conf := config.GetInstance()
	//log.Print(conf)

	os.Setenv("CB_DBHOST", "localhost")
	//os.Setenv("CB_DBPORT", "3306")
	//os.Setenv("CB_DBUSER", "root")
	os.Setenv("CB_DBNAME", "test")
	session := db.GetSession()

	//db.Create(&model.User{Email: "shetty@live.com", Name: "Shetty", Phone: "8970820090"})
	//
	//db.Create(&model.User{Email: "chetty@live.com", Name: "Chetty", Phone: "8970820090"})
	//
	//db.Create(&model.Contact{
	//	Email:  "abc@live.com",
	//	Name:   "ABC",
	//	Phone:  "909i0i932",
	//	UserId: "shetty@live.com",
	//})
	//
	//db.Create(&model.Contact{
	//	Email:  "abc@live.com",
	//	Name:   "ABC",
	//	Phone:  "909i0i932",
	//	UserId: "chetty@live.com",
	//})

	var contacts []model.Contact
	var user = model.User{Email: "shetty@live.com"}
	session.Find(&user).Related(&contacts)
	fmt.Println(contacts[0].UserId)

	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	api := r.Group("/api/v1")

	// No auth endpoints
	{
		api.POST("/register", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	// basic auth endpoints
	{
		basicAuth := api.Group("/")
		basicAuth.Use(AuthHandler)
		{
			basicAuth.GET("/test", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "oho",
				})
			})
		}
	}

	r.Run()
}
func AuthHandler(ctx *gin.Context) {
	auth := strings.SplitN(ctx.GetHeader("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		ctx.JSON(http.StatusUnauthorized, "Authorization failed")
		ctx.Abort()
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !validateCredentials(pair[0], pair[1]) {
		ctx.JSON(http.StatusUnauthorized, "Authorization failed")
		ctx.Abort()
		return
	}

	ctx.Next()
}

func validateCredentials(username, password string) bool {
	fmt.Printf("user and password is : %s, %s", username, password)
	if username == "shetty" && password == "1234" {
		return true
	}
	return false
}
