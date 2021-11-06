Dalam setiap pekerjaan di atas anda diwajibkan untuk melengkapi hal-hal berikut

1. Buatkan Dokumen Teknis dalam bentuk

   a. Entity Relationship Diagram (ERD) - Bobot 10%

  ![ERD Merchant Service](https://user-images.githubusercontent.com/37678093/140065206-e2e36184-cd9e-430c-bfcb-77f040f44f45.png)

   b. Data Manipulation Language (DML) - Bobot 5%

### Table Users

- **Get all users**

   ```
SELECT id, full_name , email
FROM users
ORDER BY created_at DESC
   ```

- **Register User**

   ```
INSERT INTO users(id, full_name, email, password, created_at, updated_at)
VALUES(?,?,?,?,?,?)
   ```

- **Login User**

   ```
SELECT * 
FROM users 
WHERE email = ?
   ```

- **Find User By Id**

   ```
SELECT * 
FROM users 
WHERE id = ?
   ```

- **Delete User By Id**

   ```
DELETE FROM users 
WHERE id = ?
   ```

- **Update User By Id**

   ```
UPDATE users 
SET full_name= ?, email= ?, password= ?,updated_at = ?
WHERE id = ?
   ```

- **Find Outlet User By Id**

   ```
SELECT * FROM outlets 
WHERE id = ?
   ```

- **Create Outlet by User**

   ```
INSERT INTO outlets(id, outlet_name, picture, user_id, created_at, updated_at)
VALUES(?,?,?,?,?,?)
   ```

- **Get All Outlets**

   ```
SELECT id, outlet_name, picture, user_id
FROM outlets
ORDER BY created_at DESC
   ```



### Table Products

- **Create Product**

   ```
INSERT INTO products(id, product_name, price, sku, picture, created_at, updated_at, outlet_id)
VALUES(?,?,?,?,?,?,?,?)
   ```

- **Find All Product**

   ```
SELECT id, product_name, price, sku, picture, outlet_id
FROM products
ORDER BY created_at DESC
   ```

- **Find Product By Id**

   ```
SELECT id, product_name, price, sku, picture, outlet_id 
FROM products 
WHERE id = ?
   ```

- **Update Product By Id**

   ```
UPDATE products 
SET product_name = ?, price = ?, sku = ?, picture = ?, outlet_id = ?, updated_at = ?
WHERE id = ?
   ```

- **Delete Product By Id**

   ```
DELETE FROM products 
WHERE id = ?
   ```



   c. Activity Diagrams - Bobot 10%
   ![Activity Diagrams drawio](https://user-images.githubusercontent.com/37678093/140091648-590171fb-c8cd-4ea9-a2fd-704ef134539c.png)

   d. Use Case Diagrams - Bobot 5%
   ![Usecase Diagram drawio 2](https://user-images.githubusercontent.com/37678093/140083486-32fa928f-1a6f-4da4-a0af-cbb5610d4b85.png)


2. Coding pekerjaan Poin A, B, dan C menggunakan REST API - Bobot 40%

   a. Gunakan database Mysql atau PostgreSQL (using MySQL) :heavy_check_mark:

   b. Dikerjakan dengan menggunakan Golang (using Golang) :heavy_check_mark:

   c. Untuk PHP menggunakan framework CodeIgniter/Laravel

   d. Gunakan best practice Rest API untuk menentukan Response Code (using http respone code) :heavy_check_mark:

3. Poin Plus :

   a. Optimasi Database dengan memperhatikan indexing dan tipe data - Bobot 5% :heavy_check_mark:

   b. Untuk authorization gunakan JWT - Bobot 5% :heavy_check_mark:

   c. Untuk produk bisa melakukan upload gambar produk - Bobot 5% :heavy_check_mark:

   d. Menggunakan Form Validation - Bobot 5% :heavy_check_mark:

   e. Gunakan prinsip-prinsip dalam Object Oriented Programing - Bobot 10% (golang has no OOP)



Hal-hal lain yang perlu saya tambahkan:

a. File Postman

b. Link Deploy

c. Api Spec

# POST MINI APP RESTFUL API

```
https://merchant.herokuapp.com/
```


# List of Available Endpoints

### Users

- `GET /users/all`
- `POST /users/register`
- `POST /users/login`
- `GET /users/:user_id` 
- `PUT /users/:user_id`
- `DELETE /users/:user_id`

### Outlets

- `GET /outlets/all`
- `POST /outlets/create`
- `GET /outlets/:outlet_id` 
- `PUT /outlets/:outlet_id`
- `DELETE /outlets/:outlet_id`

### Products

- `GET /products/all`
- `POST /products/create`
- `GET /products/:product_id` 
- `PUT /products/:product_id`
- `DELETE /products/:product_id`



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
          "full_name" : "robert",
          "email" : "robert@mail.com"
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
          "full_name" : "Michael",
          "email" : "michael@mail.com"
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

## RESTful Endpoints Outlets

### GET /outlets/all

> Get All outlets

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
          "id" : "01522ac9-ff4c-4306-8ecf-7d6f64d7b074",
          "outlet_name" : "Toko Roti Cabang Jakarta",
          "user_id" : "59de8f50-7ccc-4bfc-8f65-f496be0f7033"
      }, {
          "id" : "b0bdeb76-bdcc-40a4-8d15-cfdc1aa573b3",
          "outlet_name" : "Toko Roti Cabang Bandung",
          "user_id" : "59de8f50-7ccc-4bfc-8f65-f496be0f7033"
      }
      , {
          "id" : "d575db68-ff86-4ed8-ad09-f9a056e103fc",
          "outlet_name" : "Toko Roti Cabang Surabaya",
          "user_id" : "59de8f50-7ccc-4bfc-8f65-f496be0f7033"
      }
      , {
          "id" : "1d1d7cbd-4508-4ab8-aa1b-913a0ea0f71a",
          "outlet_name" : "Toko Roti Cabang Semarang",
          "user_id" : "59de8f50-7ccc-4bfc-8f65-f496be0f7033"
      }
  ],
  "message": "success"
}
```

*Response (401 - Unauthorized)*

```
{
  "status" : "401",
  "data" : 
      {
        "error" : ""
      },
  "message" : "Unauthorize"
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

### POST /outlets/create

> Create new outlets

_Request Header_

```
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```json
{
  "outlet_name": "Toko Roti cabang Yogyakarta",
}
```

_Response (201)_

```json
{
  "status" : 201,
  "data" : 
      {
         "id" : <given id by system>,
		   "outlet_name" : "<posted outlet name>"
      },
   "message" : "success create new outlet",
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

### GET /outlets/:outlet_id

> Get outlet by Id

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
        "id" : "1d1d7cbd-4508-4ab8-aa1b-913a0ea0f71a",
        "outlet_name" : "Toko Roti cabang Semarang",
        "outlets" : [
           {
               "id" : "e5bd18e5-d797-464c-8674-341d1f818c2b",
               "product_name" : "Roti Bagel",
               "price" : 250000,
               "sku" : "RTI-BGL-001",
               "display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
               "outlet_id" : "1d1d7cbd-4508-4ab8-aa1b-913a0ea0f71a"
           },
            {
               "id" : "d2d0c82d-23cc-4db0-ad26-f563580da197",
               "product_name" : "Roti Biali",
               "price" : 300000,
               "sku" : "RTI-BLI-002",
               "display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
               "outlet_id" : "1d1d7cbd-4508-4ab8-aa1b-913a0ea0f71a"
           }
        ]
      },
   "message" : "success get outlet by Id"
}
```

_Response (400 - Bad Request)_

```json
{
  "status" : "400"
  "data" : 
      {
        "error" : "outlet id <id? not found"
      },
  "message" : "error bad request outlet Id"
}
```

*Response (401 - Unauthorized)*

```
{
  "status" : "401",
  "data" : 
      {
        "error" : ""
      },
  "message" : "Unauthorize"
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

### PUT /outlets/:outlet_id

> Update outlet by Id

_Request Header_

```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```json
{
    "outlet_name" : "Toko Roti cabang Semarang Terbaik",
}
```

_Response (200)_

```json
{
  "status" : 200,
  "data" :
      {
        "id" : "1d1d7cbd-4508-4ab8-aa1b-913a0ea0f71a",
        "outlet_name" : "Toko Roti cabang Semarang Terbaik"
      },
  "message" : "success update outlet by Id"
}
```

*Response (401 - Unauthorized)*

```
{
  "status" : "401",
  "data" : 
      {
        "error" : ""
      },
  "message" : "Unauthorize"
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

### DELETE /outlets/:outlets_id

> Delete outlet by ID

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
  "message" : "success delete outlet by Id",
}
```

*Response (401 - Unauthorized)*

```
{
  "status" : "401",
  "data" : 
      {
        "error" : ""
      },
  "message" : "Unauthorize"
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

## RESTful Endpoints Products

### GET /products/all

> Get All products

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
          "id" : "e5bd18e5-d797-464c-8674-341d1f818c2b",
          "product_name" : "Roti Bagel",
          "price" : 250000,
          "sku" : "RTI-BGL-001",
          "display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
          "outlet_id" : "1d1d7cbd-4508-4ab8-aa1b-913a0ea0f71a"
      }, {
          "id" : "0e57af3d-c33d-431b-abfd-0d653a2bad5f",
          "product_name" : "Roti Bolilo",
          "price" : 150000,
          "sku" : "RTI-BLL-001",
          "display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
          "outlet_id" : "0a9f05b9-9847-47b0-b50c-8dc41c41cf8d"
      }
      , {
          "id" : "d7ba8992-4c5a-41cb-944f-08d8d3bbffea",
          "product_name" : "Roti Grissini",
          "price" : 250000,
          "sku" : "RTI-GRI-001",
          "display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
          "outlet_id" : "66d33754-2c22-46ac-bde8-3b102c1ba0ec"
      }
      , {
          "id" : "16d5a573-79bb-4bde-a7a3-e764424eaa23",
          "product_name" : "Roti Ciabatta",
          "price" : 450000,
          "sku" : "RTI-CBT-001",
          "display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
          "outlet_id" : "21df4b09-85fd-4a86-ac12-3cf2a61bb39b"
      }
  ],
  "message": "success"
}
```

*Response (401 - Unauthorized)*

```
{
  "status" : "401",
   "meta" : {
    "total_page": 0,
    "current_page": 0,
    "page_size": 0,
    "total": 0
  }, 
  "data" : 
      {
        "error" : ""
      },
  "message" : "Unauthorize"
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

### POST /products/create

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
  "product_name" : "Roti Grissini",
  "price" : 250000,
  "sku" : "RTI-GRI-001",
  "display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
  "outlet_id" : "66d33754-2c22-46ac-bde8-3b102c1ba0ec"
}
```

_Response (201)_

```json
{
  "status" : 201,
  "data" : 
      {
         "id" : <given id by system>,
	 "product_name" : "<posted product name>",
	 "price" : "<posted price>",
	 "sku" : "<posted sku>",
	 "display_image" : "<posted display image>",
	 "outlet_id" : "<posted outlet id>",
      },
   "message" : "success create new product",
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

*Response (401 - Unauthorized)*

```
{
  "status" : "401",
  "data" : 
      {
        "error" : ""
      },
  "message" : "Unauthorize"
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

### GET /products/:product_id

> Get product by Id

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
         "id" : "e5bd18e5-d797-464c-8674-341d1f818c2b",
         "product_name" : "Roti Bagel",
         "price" : 250000,
         "sku" : "RTI-BGL-001",
         "display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
         "outlet_id" : "1d1d7cbd-4508-4ab8-aa1b-913a0ea0f71a"
      },
   "message" : "success get product by Id"
}
```

_Response (400 - Bad Request)_

```json
{
  "status" : "400"
  "data" : 
      {
        "error" : "product id <id? not found"
      },
  "message" : "error bad request product Id"
}
```

*Response (401 - Unauthorized)*

```
{
  "status" : "401",
  "data" : 
      {
        "error" : ""
      },
  "message" : "Unauthorize"
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
    "product_name" : "Roti Bagel good quality",
    "price" : 300000,
    "sku" : "RTI-BGL-001",
    "display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
    "outlet_id" : "1d1d7cbd-4508-4ab8-aa1b-913a0ea0f71a"
}
```

_Response (200)_

```json
{
  "status" : 200,
  "data" :
      {
        "id" : "e5bd18e5-d797-464c-8674-341d1f818c2b",
    	"product_name" : "Roti Bagel good quality",
    	"price" : 300000,
    	"sku" : "RTI-BGL-001",
    	"display_image" : "https://merchants.herokuapp.com/products/foto-product-baru.jpg",
    	"outlet_id" : "1d1d7cbd-4508-4ab8-aa1b-913a0ea0f71a"
      },
  "message" : "success update product by Id"
}
```

*Response (401 - Unauthorized)*

```
{
  "status" : "401",
  "data" : 
      {
        "error" : ""
      },
  "message" : "Unauthorize"
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

### DELETE /products/:products_id

> Delete product by Id

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
  "message" : "success delete product by Id",
}
```

*Response (401 - Unauthorized)*

```
{
  "status" : "401",
  "data" : 
      {
        "error" : ""
      },
  "message" : "Unauthorize"
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

