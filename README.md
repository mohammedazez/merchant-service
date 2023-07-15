1. Buatlah struktur database (ERD dan/atau Table Relationship). (20)

![ERD Aplikasi Receipt Luscious drawio](https://github.com/mohammedazez/merchant-service/assets/37678093/52585e4f-9fc9-4f80-84e2-e2fdbde82417)

![Example data drawio](https://github.com/mohammedazez/merchant-service/assets/37678093/72261e78-c902-4f4c-9aab-ee4f90038d2d)

![User Interface](https://github.com/mohammedazez/merchant-service/assets/37678093/cf5a582d-1ead-4e1a-bfa3-d75da5c8f3ac)


2. Membuat API CRUD (create, read / index, update, delete) data master (bahan dan kategori). (20)

   

   ### Ingredients

   - Endpoint: `GET /api/v1/ingredient	`

     Method ini akan digunakan untuk menampilkan bahan makanan.

     > Get All Ingredients

     

     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/ingredient
     ```

     _Response (200)_

     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "ingredient_name": "Bawang Merah",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "2 siung",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8ccba789904bbfd7c3b0",
           "ingredient_name": "Bawang Putih",
           "recipe_id": "643f8a7ebc2985b6227584ee",
           "quantity": "2 siung",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8c6b0f2d077622e36292",
           "ingredient_name": "Garam",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "secukupnya,
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8c4d0f2d077622e36291",
           "ingredient_name": "Cabe",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "secukupnya",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8c3a0f2d077622e36290",
           "ingredient_name": "Ketumbar",
           "recipe_id": "643f8a7ebc2985b6227584ee",
           "quantity": "secukupnya",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8bd85823cf72b856117c5",
           "ingredient_name": "Telur",
           "recipe_id": "643f8d13a789904bbfd7c3b3",
           "quantity": "2 buah",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         }
       ]
     }
     ```
   
     _Response (500 - Internal Server Error)_
   
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
   
     ---
   
   - `POST /api/v1/ingredient`
     Method ini akan digunakan untuk menambahkan data bahan makanan.

     > Add new ingredient
   
     
   
     _Request Body_
   
     ```json
     {
       "ingredient_name": "Bawang Merah",
       "recipe_id": "643f8d13a789904bbfd7c3b3",
       "quantity": "2 siung"
     }
     ```
     
     _Response (201)_
     
     ```json
     {
       "status": true,
       "message": "created",
       "code": 201,
       "data": {
         "id": "643f8d13a789904bbfd7c3b3",
         "ingredient_name": "Bawang Merah",
         "recipe_id": "643f8b755823cf72b856117b",
         "quantity": "2 siung",
         "created_at": "2023-07-15T06:41:23.451Z",
         "updated_at": null
       }
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     Bad Request (400 - bad request)
     
     ```json
     {
       "status": false,
       "message": "bad request",
       "code": 400,
       "error": {
         "ingredient_name": "cannot be blank",
         "recipe_id": "cannot be blank",
         "quantity": "cannot be blank"
       }
     }
     ```
     
     ---
     
     
     
   - `PUT /api/v1/ingredient/:ingredient_id`
     Method ini akan digunakan untuk update data bahan makanan.
   
     > Edit ingredient
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/ingredient/643f8d13a789904bbfd7c3b3
     ```
     
     _Request Body_
     
     ```json
     {
       "ingredient_name": "Bawang Merah",
       "recipe_id": "643f8d13a789904bbfd7c3b3",
       "quantity": "4 siung"
     }
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": {
         "id": "643f8d13a789904bbfd7c3b3",
         "ingredient_name": "Bawang Merah",
         "recipe_id": "643f8b755823cf72b856117b",
         "quantity": "4 siung",
         "created_at": "2023-07-15T06:41:23.451Z",
         "updated_at": "2023-07-16T06:41:23.451Z"
       }
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---
     
   - `DELETE /api/v1/ingredients/:ingredient_id`
     Method ini akan digunakan untuk menghapus data bahan makanan.
   
     > Delete ingredient
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/ingredient/643f8d13a789904bbfd7c3b3
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---
     
     
   
   ### Category
   
   - `GET /api/v1/category	`
     Method ini akan digunakan untuk menampilkan kategori makanan.
   
     > Get All category
   
     
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category
     ```
   
     _Response (200)_
   
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "category_name": "Menu Utama",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8ccba789904bbfd7c3b0",
           "category_name": "Hidangan Pembuka",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8c6b0f2d077622e36292",
           "category_name": "Masakan Sayuran",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8c4d0f2d077622e36291",
           "category_name": "Masakan Ikan",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8c3a0f2d077622e36290",
           "category_name": "Minuman",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8bd85823cf72b856117c5",
           "category_name": "Hidangan Penutup",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         }
       ]
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---
     
   - `GET /api/v1/category/:category_id 	`
     Method ini akan digunakan untuk menampilkan detail kategori makanan.
   
     > Get detail category
   
     
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category/643f8d13a789904bbfd7c3b3
     ```
   
     _Response (200)_
   
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": {
         "id": "643f8d13a789904bbfd7c3b3",
         "category_name": "Masakan Ikan",
         "created_at": "2023-07-15T06:41:23.451Z",
         "updated_at": null,
         "recipe": [
           {
             "id": "643f8d13a789904bbfd7c3b3",
             "recipe_name": "Ikan Mas Bakar",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8ccba789904bbfd7c3b0",
             "recipe_name": "Ikan Gurame Goreng",
             "created_at": "2023-07-15T06:41:23.451Z"
           },
           {
             "id": "643f8c6b0f2d077622e36292",
             "recipe_name": "Pepes Ikan",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8c4d0f2d077622e36291",
             "recipe_name": "Ikan Tuna Bakar",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8c3a0f2d077622e36290",
             "recipe_name": "Ikan Asam Manis",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8bd85823cf72b856117c5",
             "recipe_name": "Ikan Tuna Bakar",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           }
         ]
       }
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---


   - `POST /api/v1/category`
       Method ini akan digunakan untuk menambahkan data kategori makanan.

     > Add new category

     

     _Request Body_

     ```json
     {
       "category_name": "Menu Utama"
     }
     ```
     
     _Response (201)_
     
     ```json
     {
       "status": true,
       "message": "created",
       "code": 201,
       "data": {
         "id": "643f8d13a789904bbfd7c3b3",
         "category_name": "Menu Utama",
         "created_at": "2023-07-15T06:41:23.451Z",
         "updated_at": null
       }
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     Bad Request (400 - bad request)
     
     ```json
     {
       "status": false,
       "message": "bad request",
       "code": 400,
       "error": {
         "category_name": "cannot be blank"
       }
     }
     ```
     
     ---



   - `PUT /api/v1/category/:category_id`
     Method ini akan digunakan untuk update data category.

     > Edit category

     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category/643f8d13a789904bbfd7c3b3
     ```
     
     _Request Body_
     
     ```json
     {
       "category_name": "Menu Utama Baru"
     }
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": {
         "id": "643f8d13a789904bbfd7c3b3",
         "category_name": "Menu Utama Baru",
         "created_at": "2023-07-15T06:41:23.451Z",
         "updated_at": "2023-07-16T06:41:23.451Z"
       }
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---


   - `DELETE /api/v1/category/:category_id`

     Method ini akan digunakan untuk menghapus data category.

     > Delete category
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category/643f8d13a789904bbfd7c3b3
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---




2. Membuat API CRUD data resep dan beserta menampilkan bahan-bahan dan kategori yang digunakan. (20)

### 	Recipe

- `GET /api/v1/recipe	`
     Method ini akan digunakan untuk menampilkan resep makanan.

     > Get All recipe

     
   
     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/recipe
     ```
   
     _Response (200)_
   
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "recipe_name": "Ikan Mas Bakar",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8ccba789904bbfd7c3b0",
           "recipe_name": "Ikan Gurame Goreng",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8c6b0f2d077622e36292",
           "recipe_name": "Pepes Ikan",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8c4d0f2d077622e36291",
           "recipe_name": "Ikan Tuna Bakar",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8c3a0f2d077622e36290",
           "recipe_name": "Ikan Asam Manis",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         },
         {
           "id": "643f8bd85823cf72b856117c5",
           "recipe_name": "Ikan Tuna Bakar",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         }
       ]
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---


- `GET /api/v1/recipe/:recipe_id`
     Method ini akan digunakan untuk menampilkan detail resep.
   
     > Get detail recipe
   
     
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/recipe/643f8d13a789904bbfd7c3b3
     ```
   
     _Response (200)_
   
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": {
         "id": "643f8d13a789904bbfd7c3b3",
         "recipe_name": "Ikan Tuna Bakar",
         "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
         "created_at": "2023-07-15T06:41:23.451Z",
         "updated_at": null,
         "ingredient": [
           {
             "id": "643f8d13a789904bbfd7c3b3",
             "ingredient_name": "Bawang Merah",
             "recipe_id": "643f8d13a789904bbfd7c3b3",
             "quantity": "2 siung",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8ccba789904bbfd7c3b0",
             "ingredient_name": "Bawang Putih",
             "recipe_id": "643f8d13a789904bbfd7c3b3",
             "quantity": "2 siung",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8c6b0f2d077622e36292",
             "ingredient_name": "Garam",
             "recipe_id": "643f8d13a789904bbfd7c3b3",
             "quantity": "secukupnya",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           }
         ]
       }
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---
     
- `POST /api/v1/recipe`
       Method ini akan digunakan untuk menambahkan resep makanan.

     > Add new recipe

     

     _Request Body_

     ```json
     {
       "recipe_name": "Menu Utama",
       "category_id": "643f8c4d0f2d077622e36291",
       "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
       "ingredient": [
         {
           "id": "643f8d13a789904bbfd7c3b3"
         },
         {
           "id": "643f8ccba789904bbfd7c3b0"
         },
         {
           "id": "643f8c6b0f2d077622e36292"
         }
       ]
     }
     ```
   
     _Response (201)_
   
     ```json
     {
       "status": true,
       "message": "created",
       "code": 201,
       "data": {
         "recipe_name": "Menu Utama",
         "category_id": "643f8c4d0f2d077622e36291",
         "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
         "ingredient": [
           {
             "id": "643f8d13a789904bbfd7c3b3"
           },
           {
             "id": "643f8ccba789904bbfd7c3b0"
           },
           {
             "id": "643f8c6b0f2d077622e36292"
           }
         ]
       }
     }
     ```
   
     _Response (500 - Internal Server Error)_
   
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
   
     Bad Request (400 - bad request)
   
     ```json
     {
       "status": false,
       "message": "bad request",
       "code": 400,
       "error": {
         "recipe_name": "cannot be blank",
     	"category_id": "cannot be blank",
      	"how_to_make": "cannot be blank"
       }
     }
     ```
   
   ---
   
- `PUT /api/v1/recipe/:recipe_id`
     Method ini akan digunakan untuk update data resep makanan.

     > Edit resep makanan

     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/recipe/643f8d13a789904bbfd7c3b3
     ```
     
     _Request Body_
     
     ```json
     {
       "recipe_name": "Menu Utama",
       "category_id": "643f8c4d0f2d077622e36291",
       "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
       "ingredient": [
         {
           "id": "643f8d13a789904bbfd7c3b3"
         },
         {
           "id": "643f8ccba789904bbfd7c3b0"
         },
         {
           "id": "643f8c6b0f2d077622e36292"
         }
       ]
     }
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": {
         "recipe_name": "Menu Utama",
         "category_id": "643f8c4d0f2d077622e36291",
         "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
         "ingredient": [
           {
             "id": "643f8d13a789904bbfd7c3b3"
           },
           {
             "id": "643f8ccba789904bbfd7c3b0"
           },
           {
             "id": "643f8c6b0f2d077622e36292"
           }
         ]
       }
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---
     
- `DELETE /api/v1/recipe/:recipe_id`
     Method ini akan digunakan untuk menghapus data resep makanan.

     > Delete recipe
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/recipe/643f8d13a789904bbfd7c3b3
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---


4. Membuat API CRUD Menampilkan list / index / read resep berdasarkan search / filter bahan dan kategori yang digunakan. (20)

     Filter Bahan makanan

     _Path Example_
     
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/ingredient?ingredient_name=garam
     ```

     | Parameter       | Data Type | Location |
     | :-------------- | :-------: | -------: |
     | ingredient_name |  String   |    query |
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8c6b0f2d077622e36292",
           "ingredient_name": "Garam",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "secukupnya,
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         }
       ]
     }
     ```
     
     

     
     
     Filter kategori makanan
     
     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category?category_name=menu
     ```
     
     | Parameter     | Data Type | Location |
     | :------------ | :-------: | -------: |
     | category_name |  String   |    query |
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "category_name": "Menu Utama",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         }
       ]
     }
     ```
     


5. Buat unit test untuk setiap fungsi yang telah dikerjakan (20)
