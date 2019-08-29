# Contacts Book
[![Build Status](https://travis-ci.org/shettyh/contacts-book.svg?branch=master)](https://travis-ci.org/shettyh/contacts-book)
[![codecov](https://codecov.io/gh/shettyh/contacts-book/branch/master/graph/badge.svg)](https://codecov.io/gh/shettyh/contacts-book)
[![GoDoc](https://godoc.org/github.com/shettyh/contacts-book?status.svg)](https://godoc.org/github.com/shettyh/contacts-book)
[![Go Report Card](https://goreportcard.com/badge/github.com/shettyh/contacts-book)](https://goreportcard.com/report/github.com/shettyh/contacts-book)

## API Docs

### Register
User needs to register to the APP to create their contacts book. Using register API they can register.

* **URL** : `/api/v1/register`
* **Request Type** : `PUT`
* **Content Type** : `application/json`
* **Request Body**: 
    ```
    {
        "email" : "email",
        "name" : "name",
        "phone" : "contact number",
        "password" : "password"
    }
     ```
* **Response Body** : ( On status code 200 )
    ```
    {
        "status": "Registered successfully"
    }
    ```
 
### Add Contact
User can add a new contact using this API. This API needs authentication.

* **URL** : `/api/v1/user/contacts/add`
* **Request Type** : `PUT`
* **Content Type** : `application/json`
* **Authorization** : `Basic <base64(username:password)>`
* **Request Body**: 
    ```
    {
        "email" : "emailid of the contact",
        "name" : "name of the contact",
        "phone" : "contact number"
    }
     ```

### Update Contact
User can update an existing contact using this API. This API needs authentication.

* **URL** : `/api/v1/user/contacts/update`
* **Request Type** : `POST`
* **Content Type** : `application/json`
* **Authorization** : `Basic <base64(username:password)>`
* **Request Body**: 
    ```
    {
        "email" : "existing emailid of the contact",
        "name" : "name of the contact",
        "phone" : "contact number"
    }
     ```
  
### Delete Contact
User can delete an existing contact using this API. This API needs authentication.

* **URL** : `/api/v1/user/contacts/<contact email id>`
* **Request Type** : `DELETE`
* **Authorization** : `Basic <base64(username:password)>`

### Get all contacts
User can get all the contacts using this API. This API can be paginated using the below mentioned Query params and the results will be Sorted based on EmailId of the contact.
* **URL** : `/api/v1/user/contacts?pageNo=<page no>&pageSize=<page size>`
* **Request Type** : `GET`
* **Content Type** : `application/json`
* **Response Body** : ( On status code 200 )
    ```
    [
        {
            "email": "email of the contact",
            "name": "Name of the contact",
            "phone": "contacts phone number",
            "user_id": "Users email ID"
        },
        ...
    ]
    ```
### Search Contacts
User can search for contacts using name or email or both using this API. This API can be paginated using the below mentioned Query params and the results will be Sorted based on EmailId of the contact.
* **URL** : `/api/v1/user/contacts/search?emailId=<emailId>&name=<name>&pageNo=<page no>&pageSize=<page size>`
* **Request Type** : `GET`
* **Content Type** : `application/json`
* **Response Body** : ( On status code 200 )
    ```
    [
        {
            "email": "email of the contact",
            "name": "Name of the contact",
            "phone": "contacts phone number",
            "user_id": "Users email ID"
        },
        ...
    ]
    ```
  
 
## Enhancements (TODO)
- [ ] Use separate models for API requests and have DTO's for database model conversion. So that API requests and DB models will be independent of each other.
- [ ] Have cache/session in authentication middleware so that no need to query DB for user credentials validation.
- [ ] Add Integration tests.
- [ ] Add more unit tests.
- [ ] Horizontal scaling support with New SQL databases like Cockroach or TiDB.
- [ ] Add docker file.
- [x] Add Makefile for build and test.
- [ ] Input fields validation for all fields like email, phone etc
- [ ] Dont take DB password in plain text from config.
- [x] Add travis Build.
- [ ] Better error messages, like Email already exists etc.
