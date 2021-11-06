Dalam setiap pekerjaan di atas anda diwajibkan untuk melengkapi hal-hal berikut

1. Buatkan Dokumen Teknis dalam bentuk

   a. Entity Relationship Diagram (ERD) - Bobot 10%
![ERD](https://user-images.githubusercontent.com/37678093/140611942-ed5b7de1-94f9-486b-9a79-ae90cbdbf6f6.png)


   b. Data Manipulation Language (DML) - Bobot 5%

### Table Users

**Get all users**

   ```
SELECT id, full_name , email
FROM users
ORDER BY created_at DESC
   ```
**Register User**

   ```
INSERT INTO users(id, full_name, email, password, created_at, updated_at)
VALUES(?,?,?,?,?,?)
   ```

**Login User**

   ```
SELECT * 
FROM users 
WHERE email = ?
   ```

**Find User By Id**

   ```
SELECT * 
FROM users 
WHERE id = ?
   ```

**Delete User By Id**

   ```
DELETE FROM users 
WHERE id = ?
   ```

**Update User By Id**

   ```
UPDATE users 
SET full_name= ?, email= ?, password= ?,updated_at = ?
WHERE id = ?
   ```

**Find Outlet User By Id**

   ```
SELECT * FROM outlets 
WHERE id = ?
   ```

**Create Outlet by User**

   ```
INSERT INTO outlets(id, outlet_name, picture, user_id, created_at, updated_at)
VALUES(?,?,?,?,?,?)
   ```

**Get All Outlets**

   ```
SELECT id, outlet_name, picture, user_id
FROM outlets
ORDER BY created_at DESC
   ```



### Table Products

**Create Product**

   ```
INSERT INTO products(id, product_name, price, sku, picture, created_at, updated_at, outlet_id)
VALUES(?,?,?,?,?,?,?,?)
   ```

**Find All Product**

   ```
SELECT id, product_name, price, sku, picture, outlet_id
FROM products
ORDER BY created_at DESC
   ```

**Find Product By Id**

   ```
SELECT id, product_name, price, sku, picture, outlet_id 
FROM products 
WHERE id = ?
   ```

**Update Product By Id**

   ```
UPDATE products 
SET product_name = ?, price = ?, sku = ?, picture = ?, outlet_id = ?, updated_at = ?
WHERE id = ?
   ```

**Delete Product By Id**

   ```
DELETE FROM products 
WHERE id = ?
   ```



   c. Activity Diagrams - Bobot 10%
![Untitled Diagram-Page-4 drawio](https://user-images.githubusercontent.com/37678093/140614126-5bb39609-a65f-4f33-873d-87dcbecac1bf.png)


   d. Use Case Diagrams - Bobot 5%
![usecase](https://user-images.githubusercontent.com/37678093/140612864-ac546bd4-90ff-4031-806d-a4fd2f975c97.png)



2. Coding pekerjaan Poin A, B, dan C menggunakan REST API - Bobot 40%

   a. Gunakan database Mysql atau PostgreSQL (using MySQL) :heavy_check_mark:
   
   b. Dikerjakan dengan menggunakan Golang (using Golang) :heavy_check_mark:
   
   c. Untuk PHP menggunakan framework CodeIgniter/Laravel
   
   d. Gunakan best practice Rest API untuk menentukan Response Code (using http respone code) :heavy_check_mark:

3. Poin Plus :

   a. Optimasi Database dengan memperhatikan indexing dan tipe data - Bobot 5% :heavy_check_mark:

   b. Untuk authorization gunakan JWT - Bobot 5% :heavy_check_mark:
   
   c. Untuk produk bisa melakukan upload gambar produk - Bobot 5% 
   
   d. Menggunakan Form Validation - Bobot 5% :heavy_check_mark:
   
   e. Gunakan prinsip-prinsip dalam Object Oriented Programing - Bobot 10% (golang has no OOP)



Hal-hal lain yang perlu saya tambahkan:

a. Download File Postman (compress this zip file, extract and import to your insomnia software)
[Insomnia_2021-11-06.zip](https://github.com/mohammedazez/merchant-service/files/7490319/Insomnia_2021-11-06.zip)

b. Link Deploy

c. Api Spec

# POST MINI APP RESTFUL API

```
deploy= https://merchant.herokuapp.com/
localhost= http://localhost:8010

example :
http://localhost:8010/api/users/all
https://merchant.herokuapp.com/api/users/all
```

# List of Endpoints

### Users
- `GET /api/users`
- `POST /api/users/register`
- `POST /api/users/login`
- `GET /api/users/:user_id` 
- `PUT /api/users/:user_id`
- `DELETE /api/users/:user_id`
- `POST /api/users/outlet`
- `GET /api/users/outlet`

### Products
- `POST /api/products`
- `GET /api/products/:outlet_id`
- `GET /api/products` 
- `GET /api/products/detail/:product_id`
- `PUT /api/products/:product_id`
- `DELETE /api/products/:product_id`



## RESTful Endpoints User
### GET /api/users/all

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
  "meta" : {
    "message": "success get all user User"
    "code": 200,
    "status": "status OK"
  }, 
  "data" : [
      {
          "id" : "b989074a-c603-4fa5-8b16-45620de8c538",
          "full_name" : "robert",
          "email" : "robert@mail.com"
      }, 
      {
          "id" : "01522ac9-ff4c-4306-8ecf-7d6f64d7b075",
          "full_name" : "Abdul Franklin",
          "email" : "abdulfranklin@mail.com"
      }, 
      {
          "id" : "01522ac9-ff4c-4306-8ecf-7d6f64d7b076",
          "full_name" : "Siti Isabella",
          "email" : "sitiisabella@mail.com"
      }, 
      {
          "id" : 01522ac9-ff4c-4306-8ecf-7d6f64d7b077,
          "full_name" : "Michael",
          "email" : "michael@mail.com"
      }
  ]
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
    "message": "internal server error"
    "code": 500,
    "status": "error"
  }, 
  "data" : {
      "error" : ""
  },
}
```
---

### POST /api/users/register

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
  "meta": {
    "message": "success create new user",
    "code": 201,
    "status": "Status Created"
  },
  "data": {
    "id": "98c636a3-519d-44bf-b7d5-f15b39b7be3a",
    "full_name": "robobot",
    "email": "robobot@mail.com"
  }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta": {
    "message": "input data required",
    "code": 400,
    "status": "bad request"
  },
  "data": {
    "errors": [
      "Key: 'UserInput.FullName' Error:Field validation for 'FullName' failed on the 'required' tag"
    ]
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
    "message": "internal server error"
    "code": 500,
    "status": "error"
  }, 
  "data" : {
      "error" : ""
  },
}
```
---

### POST /api/users/login

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
  "meta": {
    "message": "success login user",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": "b989074a-c603-4fa5-8b16-45620de8c538",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYjk4OTA3NGEtYzYwMy00ZmE1LThiMTYtNDU2MjBkZThjNTM4In0.qBBnpzkYsBJGWVb2RzOX51Upkw3qW9TItbj2VkcwKcA"
  }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta": {
    "message": "input data required",
    "code": 400,
    "status": "bad request"
  },
  "data": {
    "errors": "user id not found"
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
    "message": "internal server error"
    "code": 500,
    "status": "error"
  }, 
  "data" : {
      "error" : ""
  },
}
```
---

### GET /api/users/:user_id

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
  "meta": {
    "message": "success get user by ID",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111",
    "full_name": "ironman",
    "email": "ironman@mail.com",
    "created_at": "2021-11-06T17:02:36.138+07:00",
    "updated_at": "2021-11-06T17:02:36.138+07:00",
    "Outlet": [
      {
        "id": "065d9cbf-4b4f-4e2b-bb0a-7c91521e784b",
        "outlet_name": "Toko Mainan Terbaik",
        "picture": "link picture",
        "created_at": "2021-11-06T17:03:55.826+07:00",
        "updated_at": "2021-11-06T17:03:55.826+07:00",
        "product": null,
        "user_id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111"
      },
      {
        "id": "27c907e8-e5db-4fb8-a4cb-b4893212e49e",
        "outlet_name": "Toko Mainan Terbaik",
        "picture": "link picture",
        "created_at": "2021-11-06T17:03:54.778+07:00",
        "updated_at": "2021-11-06T17:03:54.778+07:00",
        "product": null,
        "user_id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111"
      },
      {
        "id": "31db31b3-de8c-4def-8e07-55f1e4bbf9d5",
        "outlet_name": "Toko Mainan Terbaik",
        "picture": "link picture",
        "created_at": "2021-11-06T17:03:55.714+07:00",
        "updated_at": "2021-11-06T17:03:55.714+07:00",
        "product": null,
        "user_id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111"
      },
    ]
  }
}
```
_Response (400 - Bad Request)_
```json
{
  "meta": {
    "message": "input params error",
    "code": 400,
    "status": "bad request"
  },
  "data": {
    "errors": "user id acb7be4f-bcb5-4727-bd50-e3a3de5f611 not found"
  }
}
```

_Response (500 - Internal Server Error)_

```json
{
  "meta" : {
    "message": "internal server error"
    "code": 500,
    "status": "error"
  }, 
  "data" : {
      "error" : ""
  },
}
```
---

### PUT /api/users/:user_id

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
	"full_name": "azizbaru",
  	"email": "azizbaru@mail.com",
	"password": "passwordbaru"
}
```

_Response (200)_
```json
{
  "meta": {
    "message": "success update user by ID",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111",
    "full_name": "azizbaru",
    "email": "azizbaru@mail.com"
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "error": "user id 589d2205-862b-4d25-b192-150c19eb654b not found"
  }
}
```
*Response (401 - Unauthorized)*

```
{
  "meta": {
    "message": "Unauthorize",
    "code": 401,
    "status": "error"
  },
  "data": {
    "error": "user ID not authorize"
  }
}
```



---

### DELETE /api/users/:user_id

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
  "meta": {
    "message": "success delete user by ID",
    "code": 200,
    "status": "success"
  },
  "data": {
    "message": "success delete user ID : 98c636a3-519d-44bf-b7d5-f15b39b7be3a",
    "time_delete": "2021-11-06T19:18:28.312614213+07:00"
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "error": ""
  }
}
```
_Response (401 - Unauthorize)_

```json
{
  "meta": {
    "message": "Unauthorize",
    "code": 401,
    "status": "error"
  },
  "data": {
    "error": "unexpected end of JSON input"
  }
}
```

_Response (400 - Bad Request)_

```json
{
  "meta": {
    "message": "error bad request delete user",
    "code": 400,
    "status": "error"
  },
  "data": {
    "error": "user id 589d2205-862b-4d25-b192-150c19eb654b not found"
  }
}
```

---

### POST /api/outlets/create

> User create new outlets

_Request Header_

```
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```json
{
	"outlet_name": "Toko Mainan Terbaik",
	"picture": "link picture",
	"user_id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111"
}
```

_Response (201)_

```json
{
  "meta": {
    "message": "success create new Outlet",
    "code": 201,
    "status": "Status OK"
  },
  "data": {
    "id": "c162adff-6239-4e21-9276-02a7985c815f",
    "outlet_name": "Toko Mainan Terbaik",
    "picture": "link picture",
    "created_at": "2021-11-06T19:24:50.372885909+07:00",
    "updated_at": "2021-11-06T19:24:50.372886009+07:00",
    "product": null,
    "user_id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111"
  }
}
```

_Response (400 - Bad Request)_

```json
{
  "meta": {
    "message": "input data required",
    "code": 400,
    "status": "bad request"
  },
  "data": {
    "errors": [
      "Key: 'OutletInput.OutletName' Error:Field validation for 'OutletName' failed on the 'required' tag"
    ]
  }
}
```

_Response (500 - Internal Server Error)_

```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "errors": "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`merchat_service_2`.`outlets`, CONSTRAINT `fk_users_outlet` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)); Error 1452: Cannot add or update a child row: a foreign key constraint fails (`merchat_service_2`.`outlets`, CONSTRAINT `fk_users_outlet` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))"
  }
}
```

Response (401 - Unauthorize)_

```json
{
  "meta": {
    "message": "Unauthorize",
    "code": 401,
    "status": "error"
  },
  "data": {
    "error": "token contains an invalid number of segments"
  }
}
```



---

### GET /api/users/outlet

> Get All users outlet

_Request Header_

```json
not needed
```

_Request Body_

```
not needed
```

_Response (200)_

```json
{
  "meta": {
    "message": "success get all Outlet",
    "code": 200,
    "status": "status OK"
  },
  "data": [
    {
      "id": "c162adff-6239-4e21-9276-02a7985c815f",
      "outlet_name": "Toko Mainan Terbaik",
      "picture": "link picture",
      "user_id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111"
    },
    {
      "id": "8cd78110-3822-4573-ab4f-b040c7e9aeb1",
      "outlet_name": "Toko Mainan Terbaik",
      "picture": "link picture",
      "user_id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111"
    },
    {
      "id": "62bbd67b-1ecb-46ea-81df-a5834cb8d380",
      "outlet_name": "Toko Mainan Terbaik",
      "picture": "link picture",
      "user_id": "acb7be4f-bcb5-4727-bd50-e3a3de5f6111"
    }
  ]
}
```

_Response (500 - Internal Server Error)_

```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "error": ""
  }
}
```

---



## RESTful Endpoints Products


### GET /api/product/:outlet_id

> Get product outlet by Id

_Request Header_

```json
not needed
```

_Request Body_

```
not needed
```

_Response (200)_

```json

  "meta": {
    "message": "success get outlet by ID",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27",
    "outlet_name": "Toko Sembako Terbaik",
    "picture": "link picture",
    "created_at": "2021-11-06T19:40:27.168+07:00",
    "updated_at": "2021-11-06T19:40:27.168+07:00",
    "product": [
      {
        "id": "12a62f4a-8c3d-42b1-ae22-bcc0583e2790",
        "product_name": "Bawang merah",
        "price": 250000,
        "sku": "sku",
        "picture": "linkPicture",
        "created_at": "2021-11-06T19:41:12.96+07:00",
        "updated_at": "2021-11-06T19:41:12.96+07:00",
        "outlet_id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27"
      },
      {
        "id": "5206464e-1839-4db2-a21e-ca17632ac00b",
        "product_name": "Bawang merah",
        "price": 250000,
        "sku": "sku",
        "picture": "linkPicture",
        "created_at": "2021-11-06T19:41:11.808+07:00",
        "updated_at": "2021-11-06T19:41:11.808+07:00",
        "outlet_id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27"
      },
      {
        "id": "57dc4409-1f49-4fc5-a5a9-a546720b78ad",
        "product_name": "Bawang merah",
        "price": 250000,
        "sku": "sku",
        "picture": "linkPicture",
        "created_at": "2021-11-06T19:41:13.357+07:00",
        "updated_at": "2021-11-06T19:41:13.357+07:00",
        "outlet_id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27"
      }
    ],
    "user_id": "b989074a-c603-4fa5-8b16-45620de8c538"
  }
}
```

_Response (400 - Bad Request)_

```json
{
  "meta": {
    "message": "input params error",
    "code": 400,
    "status": "bad request"
  },
  "data": {
    "errors": "Outlet id 37b7988e-da04-45cc-9f24-0ff8ec4b1e2 not found"
  }
}
```

_Response (500 - Internal Server Error)_

```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "error": ""
  }
}
```

---



### GET /api/products

> Get All products

_Request Header_
```json
not needed
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta": {
    "message": "success get all Product",
    "code": 200,
    "status": "status OK"
  },
  "data": [
    {
      "id": "57dc4409-1f49-4fc5-a5a9-a546720b78ad",
      "product_name": "Bawang merah",
      "price": 250000,
      "sku": "sku",
      "picture": "linkPicture",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "outlet_id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27"
    },
    {
      "id": "fcb8bd19-b5d6-43d0-9e73-b76cd6a0aaf3",
      "product_name": "Bawang merah",
      "price": 250000,
      "sku": "sku",
      "picture": "linkPicture",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "outlet_id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27"
    },
    {
      "id": "12a62f4a-8c3d-42b1-ae22-bcc0583e2790",
      "product_name": "Bawang merah",
      "price": 250000,
      "sku": "sku",
      "picture": "linkPicture",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "outlet_id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27"
    }
  ]
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "error": ""
  }
}
```


---

### POST /api/products

> Create new products

_Request Header_
```
{
   "Authorization": "<your Authorization>"
}
```
_Request Body_

```json
{
	"product_name": "Bawang merah",
	"price": 250000,
	"sku": "sku",
	"picture": "linkPicture",
	"outlet_id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27"
}
```

_Response (201)_
```json
{
  "meta": {
    "message": "success create new Product",
    "code": 201,
    "status": "Status OK"
  },
  "data": {
    "id": "57dc4409-1f49-4fc5-a5a9-a546720b78ad",
    "product_name": "Bawang merah",
    "price": 250000,
    "sku": "sku",
    "picture": "linkPicture",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "outlet_id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27"
  }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta": {
    "message": "input data required",
    "code": 400,
    "status": "bad request"
  },
  "data": {
    "errors": [
      "Key: 'ProductInput.ProductName' Error:Field validation for 'ProductName' failed on the 'required' tag"
    ]
  }
}
```

*Response (401 - Unauthorized)*

```
{
  "meta": {
    "message": "Unauthorize",
    "code": 401,
    "status": "error"
  },
  "data": {
    "error": "unauthorize"
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "error": ""
  }
}
```
---

### GET /api/products/:product_id

> Get product by Id

_Request Header_
```json
not needed
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta": {
    "message": "success get product by ID",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": "57dc4409-1f49-4fc5-a5a9-a546720b78ad",
    "product_name": "Bawang merah",
    "price": 250000,
    "sku": "sku",
    "picture": "linkPicture",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "outlet_id": "37b7988e-da04-45cc-9f24-0ff8ec4b1e27"
  }
}
```
_Response (400 - Bad Request)_
```json
{
  "meta": {
    "message": "input params error",
    "code": 400,
    "status": "bad request"
  },
  "data": {
    "errors": "product id 57dc4409-1f49-4fc5-a5a9-a546720b78a not found"
  }
}
```

_Response (500 - Internal Server Error)_

```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "error": ""
  }
}
```
---

### PUT /products/:product_id

> Update product by Id

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
	"product_name": "Roti Tawar baru",
	"price": 4500000,
	"sku": "sku 001",
	"picture": "linkPicture2",
	"outlet_id": "8cd78110-3822-4573-ab4f-b040c7e9aeb1"
}
```

_Response (200)_
```json
{
  "meta": {
    "message": "success update Product by ID",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": "64a63c63-b673-45da-815c-f92e3197561e",
    "product_name": "Roti Tawar baru",
    "price": 4500000,
    "sku": "sku 001",
    "picture": "linkPicture2",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "outlet_id": "8cd78110-3822-4573-ab4f-b040c7e9aeb1"
  }
}
```

*Response (401 - Unauthorized)*

```
{
  "meta": {
    "message": "Unauthorize",
    "code": 401,
    "status": "error"
  },
  "data": {
    "error": "illegal base64 data at input byte 33"
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "error": ""
  }
}
```
---

### DELETE /api/products/:products_id

> Delete product by Id

_Request Header_
```json
not needed
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta": {
    "message": "success delete Product by ID",
    "code": 200,
    "status": "success"
  },
  "data": {
    "message": "success delete Product ID : 4f4daacc-1f4b-4cfb-bf52-354001a8bb5f",
    "time_delete": "2021-11-06T17:15:11.090252815+07:00"
  }
}
```

*Response (401 - Unauthorized)*

```
{
  "meta": {
    "message": "Unauthorize",
    "code": 401,
    "status": "error"
  },
  "data": {
    "error": "unexpected end of JSON input"
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta": {
    "message": "internal server error",
    "code": 500,
    "status": "error"
  },
  "data": {
    "error": ""
  }
}
```
---

