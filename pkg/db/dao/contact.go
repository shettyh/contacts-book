package dao

import (
	"fmt"

	"github.com/shettyh/contacts-book/pkg/db"
	"github.com/shettyh/contacts-book/pkg/db/model"
)

type ContactDao struct {
}

func (*ContactDao) Add(contact *model.Contact) error {
	dbSession := db.GetSession()
	err := dbSession.Create(&contact).Error
	return err
}

func (*ContactDao) Update(contact *model.Contact) error {
	dbSession := db.GetSession()
	err := dbSession.Model(&model.Contact{}).Update(&contact).Error
	return err
}

func (*ContactDao) Delete(id, userId string) error {
	dbSession := db.GetSession()
	err := dbSession.Delete(&model.Contact{Email: id, UserId: userId}).Error
	return err
}

func (*ContactDao) GetAll(userId string, offset, limit int) ([]model.Contact, error) {
	dbSession := db.GetSession()

	var contacts []model.Contact
	err := dbSession.Where(&model.Contact{UserId: userId}).Offset(offset).Limit(limit).Find(&contacts).Error
	return contacts, err
}

func (*ContactDao) Search(userId, emailId, name string,
	offset, limit int) ([]model.Contact, error) {

	dbSession := db.GetSession()

	var contacts []model.Contact
	err := dbSession.Where("user_id = ? AND name LIKE ? AND email like ?",
		userId,
		fmt.Sprintf("%%%s%%", name),
		fmt.Sprintf("%%%s%%", emailId)).Offset(offset).Limit(limit).Find(&contacts).Error
	return contacts, err
}
