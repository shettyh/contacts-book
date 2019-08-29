package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/shettyh/contacts-book/pkg/config"

	"github.com/stretchr/testify/assert"

	"github.com/shettyh/contacts-book/pkg/db/model"
)

func init() {
	// Set test database details
	os.Setenv("CB_DBTYPE", "sqlite3")
	os.Setenv("CB_DBHOST", "localhost")
	os.Setenv("CB_DBNAME", "../../contactsbook")
}

func TestGetSession(t *testing.T) {
	// Validate if the DB connection works fine
	session := GetSession()

	// Check DB operations

	// Create user
	newUser := model.User{
		Email:    "shetty@live.com",
		Name:     "Shetty",
		Phone:    "8977820092",
		Password: "shetty@123",
	}
	err := session.Create(&newUser).Error
	assert.Nil(t, err)

	// Get user
	var savedUser model.User
	err = session.Model(&model.User{}).Where(model.User{Email: "shetty@live.com"}).First(&savedUser).Error
	assert.Nil(t, err)
	assert.NotNil(t, savedUser)

	// Delete User
	err = session.Model(model.User{}).Delete(model.User{Email: "shetty@live.com"}).Error
	assert.Nil(t, err)

	// Get user
	var deletedUser model.User
	err = session.Model(&model.User{}).Where(model.User{Email: "shetty@live.com"}).First(&deletedUser).Error
	assert.NotNil(t, err)
}

func TestGetConnectionDetails(t *testing.T) {
	testcases := []struct {
		EnvMap                   map[string]string
		expectedDbType           string
		expectedConnectionString string
		err                      error
	}{
		{
			EnvMap: map[string]string{
				"CB_DBTYPE":     "mysql",
				"CB_DBHOST":     "localhost",
				"CB_DBUSER":     "shettyh",
				"CB_DBPASSWORD": "shettyh@123",
				"CB_DBNAME":     "contactsbook",
				"CB_DBPORT":     "3306",
			},
			expectedDbType:           "mysql",
			expectedConnectionString: "shettyh:shettyh@123@tcp(localhost:3306)/contactsbook",
			err:                      nil,
		},
		{
			EnvMap: map[string]string{
				"CB_DBTYPE": "sqlite3",
				"CB_DBNAME": "contactsbook",
			},
			expectedDbType:           "sqlite3",
			expectedConnectionString: "contactsbook",
			err:                      nil,
		},
		{
			EnvMap: map[string]string{
				"CB_DBTYPE": "invalidDB",
			},
			err: fmt.Errorf(errUnsupportedDb, "invalidDB"),
		},
	}

	// Run test cases
	for _, tc := range testcases {
		// Set envs
		for k, v := range tc.EnvMap {
			os.Setenv(k, v)
		}
		err := config.GetInstance().Reload()
		assert.Nil(t, err)

		dbType, connectionString, err := getConnectionDetails()
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.expectedConnectionString, connectionString)
		assert.Equal(t, tc.expectedDbType, dbType)
	}
}
