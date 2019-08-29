//+build integration

package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/shettyh/contacts-book/pkg/db/model"
)

func TestRegister(t *testing.T) {
	requestBody := model.User{
		Email:    "shettyh@live.com",
		Name:     "Shetty H",
		Phone:    "8972000091",
		Password: "shetty@123",
	}

	jsonRequest, err := json.Marshal(&requestBody)
	assert.Nil(t, err)

	client := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest("PUT", "http://localhost/api/v1/register", bytes.NewBuffer(jsonRequest))
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAddContact(t *testing.T) {
	requestBody := model.Contact{
		Email: "shettyh@live.com",
		Name:  "Shetty H",
		Phone: "8972000091",
	}

	jsonRequest, err := json.Marshal(&requestBody)
	assert.Nil(t, err)

	client := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest("PUT", "http://localhost/api/v1/user/contacts/add", bytes.NewBuffer(jsonRequest))
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth("shettyh@live.com", "shetty@123")

	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpdateContact(t *testing.T) {
	requestBody := model.Contact{
		Email: "shettyh@live.com",
		Name:  "Manjunath Shetty H",
		Phone: "8972000094",
	}

	jsonRequest, err := json.Marshal(&requestBody)
	assert.Nil(t, err)

	client := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest("POST", "http://localhost/api/v1/user/contacts/update", bytes.NewBuffer(jsonRequest))
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth("shettyh@live.com", "shetty@123")

	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetAllContacts(t *testing.T) {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest("GET", "http://localhost/api/v1/user/contacts", nil)
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth("shettyh@live.com", "shetty@123")

	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var contacts []model.Contact
	err = json.NewDecoder(resp.Body).Decode(&contacts)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(contacts))
}

func TestSearchContactsByName(t *testing.T) {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest("GET", "http://localhost/api/v1/user/contacts/search?name=Manj", nil)
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth("shettyh@live.com", "shetty@123")

	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var contacts []model.Contact
	err = json.NewDecoder(resp.Body).Decode(&contacts)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(contacts))
}

func TestSearchContactsByEmail(t *testing.T) {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest("GET", "http://localhost/api/v1/user/contacts/search?emailId=live.com", nil)
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth("shettyh@live.com", "shetty@123")

	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var contacts []model.Contact
	err = json.NewDecoder(resp.Body).Decode(&contacts)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(contacts))
}

func TestDeleteContact(t *testing.T) {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest("DELETE", "http://localhost/api/v1/user/contacts/shettyh@live.com", nil)
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth("shettyh@live.com", "shetty@123")

	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
