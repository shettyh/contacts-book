package dao

import (
	"errors"

	"github.com/shettyh/contacts-book/pkg/db"
	"github.com/shettyh/contacts-book/pkg/db/model"
	"golang.org/x/crypto/bcrypt"
)

const (
	errPasswordLength = "password should be of atleast 5 characters"
)

type UserDao struct {
}

func (*UserDao) Add(user *model.User) error {
	if len(user.Password) < 5 {
		return errors.New(errPasswordLength)
	}

	// convert the password to hash before storing
	bytePassword := []byte(user.Password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user.Password = string(passwordHash)

	// Get the DB session
	dbSession := db.GetSession()

	// Create user
	err := dbSession.Create(user).Error
	return err
}

func (*UserDao) Get(userId string) (*model.User, error) {
	dbSession := db.GetSession()

	var user model.User
	err := dbSession.Where(&model.User{Email: userId}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
