Dalam setiap pekerjaan di atas anda diwajibkan untuk melengkapi hal-hal berikut

1. Buatkan Dokumen Teknis dalam bentuk

   a. Entity Relationship Diagram (ERD) - Bobot 10%

  ![ERD Merchant Service](https://user-images.githubusercontent.com/37678093/140065206-e2e36184-cd9e-430c-bfcb-77f040f44f45.png)

   b. Data Manipulation Language (DML) - Bobot 5%
   c. Activity Diagrams - Bobot 10%
   ![Activity Diagrams drawio](https://user-images.githubusercontent.com/37678093/140091648-590171fb-c8cd-4ea9-a2fd-704ef134539c.png)

   d. Use Case Diagrams - Bobot 5%
   ![Usecase Diagram drawio 2](https://user-images.githubusercontent.com/37678093/140083486-32fa928f-1a6f-4da4-a0af-cbb5610d4b85.png)


2. Coding pekerjaan Poin A, B, dan C menggunakan REST API - Bobot 40%
   a. Gunakan database Mysql atau PostgreSQL
   b. Dikerjakan dengan menggunakan Golang (Recommended)
   c. Untuk PHP menggunakan framework CodeIgniter/Laravel
   d. Gunakan best practice Rest API untuk menentukan Response Code

3. Poin Plus :
   a. Optimasi Database dengan memperhatikan indexing dan tipe data - Bobot 5%

   b. Untuk authorization gunakan JWT - Bobot 5%
   c. Untuk produk bisa melakukan upload gambar produk - Bobot 5%
   d. Menggunakan Form Validation - Bobot 5%
   e. Gunakan prinsip-prinsip dalam Object Oriented Programing - Bobot 10%



Hal-hal lain yang perlu saya tambahkan:

a. File Postman

b. Link Deploy

c. Api Spec

# CANGKOEL RESTFUL API

```
https://cangkoel.herokuapp.com/
```

CANGKOEL adalah sebuah platform yang menyediakan solusi bagi petani/perseorangan yang mempunyai lahan/komoditas namun tidak punya biaya untuk produksi kemudian mengekspor komoditas mereka! tidak punya buyer? kami yang mencarikan!



# List of Available Endpoints

### Users Petani
- `GET /users/all`
- `POST /users/register`
- `POST /users/login`
- `GET /users/:user_id` 
- `PUT /users/:user_id`
- `DELETE /users/:user_id`



## RESTful Endpoints User
### GET /users/all

> Get All users

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "status": 200,
  "meta" : {
    "total_page": 1,
    "current_page": 1,
    "page_size": 10,
    "total": 4
  }, 
  "data" : [
      {
          "id" : 01522ac9-ff4c-4306-8ecf-7d6f64d7b074,
          "full_name" : "Joko Hitler",
          "email" : "jokohitler@mail.com"
      }, {
          "id" : 01522ac9-ff4c-4306-8ecf-7d6f64d7b075,
          "full_name" : "Abdul Franklin",
          "email" : "abdulfranklin@mail.com"
      }
      , {
          "id" : 01522ac9-ff4c-4306-8ecf-7d6f64d7b076,
          "full_name" : "Siti Isabella",
          "email" : "sitiisabella@mail.com"
      }
      , {
          "id" : 01522ac9-ff4c-4306-8ecf-7d6f64d7b077,
          "full_name" : "Michael Parto",
          "email" : "michaelparto@mail.com"
      }
  ],
  "message": "success"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "status": 500,
  "meta" : {
    "total_page": 0,
    "current_page": 0,
    "page_size": 0,
    "total": 0
  }, 
  "data" : {
      "error" : ""
  },
  "message" : "Internal server error",
}
```
---

### POST /users/register

> Create new user

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "full_name" : "<full name to get insert into>",
  "email": "<email to get insert into>",
  "password": "<password to get insert into>"
}
```

_Response (201)_
```json
{
  "status" : "201"
  "data" : 
      {
        "id" : <given id by system>,
        "full_name" : "<posted full name>",
        "email" : "<posted email>"
      },
  "message" : "success create user"
}
```

_Response (400 - Bad Request)_
```json
{
  "status" : 400,
  "data" : 
      {
        "errors" : []
      },
  "message" : "input data required"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "status" : 500,
  "data" : 
      {
        "error" : ""
      },
   "message" : "Internal Server error"
}
```
---

### POST /users/login

> Compare data login on database with data inputed

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "email": "<email to get compare>",
  "password": "<password to get compare>"
}
```

_Response (200)_
```json
{
  "status" : 200,
  "data" : 
      {
        "token" : "<your access token>"
      },
  "message" : "success login user"
}
```

_Response (400 - Bad Request)_
```json
{
  "status" : "400",
  "data" : 
      {
        "errors" : []
      },
  "message" : "input data required"
}
```

_Response (401 - Unauthorized)_
```json
{
  "status" : "401"
  "data" : 
      {
        "error" : ""
      },
   "message" : "input data error"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "status" : 500,
  "data" : 
      {
        "error" : ""
      },
  "message" : "Internal Server error"
}
```
---

### GET /users/:user_id

> Get user by  ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "status" : 200,
  "data" :
      {
        "id" : 458e6a13-a641-481b-af9f-fd1575a721d3,
        "full_name" : "Michael",
        "email" : "michael@mail.com",
        "outlets" : [
           {
               "id" : 666a1c48-e63d-4ad0-ab19-70b05d8a7cc5,
               "nama_outlet" : "Toko Roti cabang Jakarta",
               "nomor_hp" : 6283213231232,
               "users_id" : 458e6a13-a641-481b-af9f-fd1575a721d3,
           },
            {
               "id" : 7c169f9f-38ed-42ef-b0a2-3f4e49170b7e,
               "nama_outlet" : "Toko Roti cabang Surabaya",
               "nomor_hp" : 6283213231232,
               "users_id" : 458e6a13-a641-481b-af9f-fd1575a721d3,
           }
        ]
      },
   "message" : "success get user by Id"
}
```
_Response (400 - Bad Request)_
```json
{
  "status" : "400"
  "data" : 
      {
        "error" : "user id <id? not found"
      },
  "message" : "error bad request user Id"
}
```

_Response (500 - Internal Server Error)_

```json
{
  "status" : 500,
  "data" : {
      "error" : ""
  },
  "message" : "Internal server error"
}
```
---

### PUT /users/:user_id

> Update user by Id

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
    "full_name" : "Michael baru",
    "email" : "michaelbaru@mail.com"
}
```

_Response (200)_
```json
{
  "status" : 200,
  "data" :
      {
        "id" : 458e6a13-a641-481b-af9f-fd1575a721d3,
        "full_name" : "Michael baru",
        "email" : "michaelbaru@mail.com"
      },
  "message" : "success update user by Id"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "status" : 500,
  "data" : {
      "error" : ""
  },
  "message" : "Internal server error"
}
```
---

### DELETE /users/:id

> Delete user by ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "status" : 200,
  "data" : "",
  "message" : "success delete user by Id",
}
```

_Response (500 - Internal Server Error)_
```json
{
  "status" : 500,
  "data" : {
      "error" : ""
  },
  "message" : "Internal server error"
}
```
---

