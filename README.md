# Contacts Book

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
- [ ] Add unit tests.
- [ ] Horizontal scaling support with New SQL databases like Cockroach or TiDB.
- [ ] Add docker file.
- [ ] Add Makefile for build and test.

