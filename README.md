"# diary-backend" 
# Backend Diary
Panduan ini dibuat untuk dapat mengakses API 

Unduh repository ke dalam komputer menggunakan perintah `git clone`. Url
repository dapat dilihat di dalam tombol code.

```
git clone <url repository> <folder tujuan>
```

### Jalankan server
Untuk dapat menjalankan server kamu harus masuk ke dalam folder `Backend`, lalu menjalankan file `main.go`. Pastikan kamu sudah menginstall Golang pada komputer kamu, jika belum kamu dapat download di https://go.dev/dl/

```
cd Backend
go run main.go
...
```
Atau kamu dapat langsung mengakses API langsung pada project yang sudah di deploy online di URL :
```
https://quiet-woodland-87309.herokuapp.com/
```

### Database Design
<a href="https://ibb.co/sgzMmBf"><img src="https://i.ibb.co/51wfM70/Diary-Database-ERD.png" alt="Diary-Database-ERD" border="0"></a>

## API EndPoint
Note :
```
Domain Name
Local  : http://Localhost/
Online : https://quiet-woodland-87309.herokuapp.com/
```

### Register User
Endpoint `Method : POST`
```
'Domain'/api/v1/users
```

Parameter `Body : JSON`
```
{
    "name": "test name",
    "gender": "Male",
    "email" : "email@email.com",
    "password": "000123"
}
```
Response
```
{
    "meta": {
        "message": "Account has been registered",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 15,
        "name": "test name",
        "gender": "Male",
        "email": "email@email.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNX0.XDQKQijiPi7lN_r6PvKNwfcfXZpf0eSMNqSB9kWv8V0"
    }
}
```

### Login User
Endpoint `Method : POST`
```
'Domain'/api/v1/sessions
```

Parameter `Body : JSON`
```
{
    "email" : "email@email.com",
    "password": "000123"
}
```
Response
```
{
    "meta": {
        "message": "Login Successfully",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 15,
        "name": "tes",
        "gender": "Male",
        "email": "Tes@gmail.com",
        "role": "user",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNX0.XDQKQijiPi7lN_r6PvKNwfcfXZpf0eSMNqSB9kWv8V0"
    }
}
```

### Email Check
Endpoint ini berguna sebagai validasi pada saat registrasi user untuk mencegah adanya email yang double 

Endpoint `Method : POST`
```
'Domain'/api/v1/check-email
```

Parameter `Body : JSON`
```
{
    "email": "Tes@gmail.com"
}
```
Response
```
{
    "meta": {
        "message": "Email is available",
        "code": 200,
        "status": "success"
    },
    "data": {
        "is_email_available": true
    }
}
```


### Upload Avatar
Endpoint ini digunakan untuk upload photo profile user 

Endpoint `Method : POST`
```
'Domain'/api/v1/upload-avatar
```

Parameter `Body : form-data`
```
input type file dengan nama `avatar` dengan value type image
<input type="file" name="avatar" .....>
```
Response
```
{
    "meta": {
        "message": "Successfully upload image",
        "code": 200,
        "status": "success"
    },
    "data": {
        "is_uploaded": true
    }
}
```

##  Books

### Get All Active Books

Endpoint `Method : GET`
```
'Domain'/api/v1/books
```

Parameter `Body : none`
```
none
```
Response
```
{
    "meta": {
        "message": "List of Books",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "title": "Tes update buku 1",
            "writer": "Update writer 1",
            "cover_image": "Images/CoverBooks/1.png",
            "slug": "tes-update-buku-1-1",
            "status": "Active",
            "created_at": "2022-06-22T15:58:35.0435371+07:00"
        },
        {
            "id": 3,
            "title": "Tes input Service",
            "writer": "Saya",
            "cover_image": "Images/CoverBooks/1.png",
            "slug": "tes-input-service-s-int-1",
            "status": "Active",
            "created_at": "2022-06-22T15:58:35.0435371+07:00"
        },
        {
            "id": 4,
            "title": "test create",
            "writer": "sayaa",
            "cover_image": "1",
            "slug": "test-create-1",
            "status": "Active",
            "created_at": "2022-06-23T13:24:24.2651241+07:00"
        }
    ]
}
```

### Get All User's Books
Untuk mendapatkan data buku yg user uplaod

Endpoint `Method : GET`
```
'Domain'/api/v1/books/user/'user_id'
```

Parameter `Body : none`
```
none
```
Response
```
{
    "meta": {
        "message": "List of Books",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "title": "Tes update buku 1",
            "writer": "Update writer 1",
            "cover_image": "Images/CoverBooks/1.png",
            "slug": "tes-update-buku-1-1",
            "status": "Active",
            "created_at": "2022-06-22T15:58:35.0435371+07:00"
        },
        {
            "id": 3,
            "title": "Tes input Service",
            "writer": "Saya",
            "cover_image": "Images/CoverBooks/1.png",
            "slug": "tes-input-service-s-int-1",
            "status": "Active",
            "created_at": "2022-06-22T15:58:35.0435371+07:00"
        },
        {
            "id": 4,
            "title": "test create",
            "writer": "sayaa",
            "cover_image": "1",
            "slug": "test-create-1",
            "status": "Active",
            "created_at": "2022-06-23T13:24:24.2651241+07:00"
        }
    ]
}
```

### Get Book Details
Untuk mendapatkan data buku secara detail

Endpoint `Method : GET`
```
'Domain'/api/v1/books/'book_id'
```

Parameter `Body : none`
```
none
```
Response
```
{
    "meta": {
        "message": "Detail of Book",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "title": "Tes update buku 1",
        "writer": "Update writer 1",
        "pages": 550,
        "synopsis": "Update synopsis 1",
        "cover_image": "Images/CoverBooks/1.png",
        "file": "Images/Books/1.pdf",
        "status": "Active",
        "slug": "tes-update-buku-1-1",
        "created_at": "2022-06-22T15:58:35.0435371+07:00",
        "category": [
            "Karya Ilmiah",
            "Komik"
        ]
    }
}
```

### Create New Book
Untuk menambahkan buku baru

Endpoint `Method : POST`
```
'Domain'/api/v1/books
```

Parameter `Body : JSON`
```
{
    "title": "Tes update buku 2",
    "writer": "Update writer 2",
    "pages": 550,
    "synopsis": "Update synopsis 2",
    "category": ["1", "4"]
}
```
Response
```
{
    "meta": {
        "message": "Success to create new book",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 18,
        "title": "Tes update buku 2",
        "writer": "Update writer 2",
        "cover_image": "",
        "slug": "tes-update-buku-2-1",
        "status": "Pending",
        "created_at": "0001-01-01T00:00:00Z"
    }
}
```

### Upload Book's Cover
Untuk upload Cover Buku

Endpoint `Method : POST`
```
'Domain'/api/v1/books/upload-image/'book_id'
```

Parameter `Body : form-data`
```
input type file dengan nama `image_cover` dengan value type image
<input type="file" name="image_cover" .....>
```
Response
```
{
    "meta": {
        "message": "Successfully upload image",
        "code": 200,
        "status": "success"
    },
    "data": {
        "is_uploaded": true
    }
 }
```

### Upload Book's File
Untuk upload File Buku

Endpoint `Method : POST`
```
'Domain'/api/v1/books/file/'book_id'
```

Parameter `Body : form-data`
```
input type file dengan nama `file` dengan value type image
<input type="file" name="file" .....>
```
Response
```
{
    "meta": {
        "message": "Successfully upload image",
        "code": 200,
        "status": "success"
    },
    "data": {
        "is_uploaded": true
    }
 }
```

### Read Book's File
Untuk membaca File Buku

Endpoint `Method : GET`
```
'Domain'/api/v1/books/read/'book_id'
```

Parameter `Body : none`
```
none
```
Response
```
{
    "meta": {
        "message": "Detail of Book",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 18,
        "title": "Tes update buku 2",
        "pages": 550,
        "file": "bookFiles/1-Sesi 3 - Communication and Presentation Skills.pptx.pdf"
    }
}
```

### Approve or Reject Book
Untuk approval Buku yang diupload

Endpoint `Method : POST`
```
'Domain'/api/v1/books/update-status
```

Parameter `Body : none` `Status : Active / Reject`
```
{
    "id": 18,
    "status": "Reject"
}
```
Response
```
{
    "meta": {
        "message": "Success to create new book",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 18,
        "title": "Tes update buku 2",
        "writer": "Update writer 2",
        "cover_image": "images/cover/1-visual-code-portada (1).png",
        "slug": "tes-update-buku-2-1",
        "status": "Reject",
        "created_at": "2022-06-26T19:36:46.5260755+07:00"
    }
}
```

### Search Book by title
Untuk pencarian buku berdasarkan title

Endpoint `Method : POST`
```
'Domain'/api/v1/books/search
```

Parameter `Body : JSON` 
```
{
    "title": "buku"
}
```
Response
```
{
    "meta": {
        "message": "List of Books",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "title": "Tes update buku 1",
            "writer": "Update writer 1",
            "cover_image": "Images/CoverBooks/1.png",
            "slug": "tes-update-buku-1-1",
            "status": "Active",
            "created_at": "2022-06-22T15:58:35.0435371+07:00"
        },
        {
            "id": 13,
            "title": "Buku Sekolah",
            "writer": "anda",
            "cover_image": "1",
            "slug": "tes-buku-1-1",
            "status": "Active",
            "created_at": "2022-06-24T15:39:56.7011263+07:00"
        },
        {
            "id": 17,
            "title": "tes buku 1",
            "writer": "anda",
            "cover_image": "",
            "slug": "tes-buku-1-1",
            "status": "Active",
            "created_at": "2022-06-24T18:10:43.0087386+07:00"
        }
    ]
}
```

### Search Book by Category
Untuk pencarian buku berdasarkan kategory

Endpoint `Method : GET`
```
'Domain'/api/v1/books/category/'category_id'
```

Parameter `Body : none` 
```
none
```
Response
```
{
    "meta": {
        "message": "List of Books",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "title": "Tes update buku 1",
            "writer": "Update writer 1",
            "cover_image": "Images/CoverBooks/1.png",
            "slug": "tes-update-buku-1-1",
            "status": "Active",
            "created_at": "2022-06-22T15:58:35.0435371+07:00"
        },
        {
            "id": 17,
            "title": "tes buku 1",
            "writer": "anda",
            "cover_image": "",
            "slug": "tes-buku-1-1",
            "status": "Active",
            "created_at": "2022-06-24T18:10:43.0087386+07:00"
        }
    ]
}
```

### Update Book's Data
Untuk mengupdate data buku

Endpoint `Method : POST`
```
'Domain'/api/v1/books/update/'book_id'
```

Parameter `Body : JSON`
```
{
    "title": "Tes update buku 2",
    "writer": "Update writer 2",
    "pages": 550,
    "synopsis": "Update synopsis 2",
    "category": ["1", "4"]
}
```
Response
```
{
    "meta": {
        "message": "Success to update book",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 2,
        "title": "Tes update buku 2 update",
        "writer": "Update writer 2 update",
        "cover_image": "1",
        "slug": "tes-update-buku-2-update-1",
        "status": "Pending",
        "created_at": "0001-01-01T00:00:00Z"
    }
}
```

### Get Last Reader 
Untuk mendapatkan data pembaca buku terakhir

Endpoint `Method : GET`
```
'Domain'/api/v1/books/history/'book_id'
```

Parameter `Body : none`
```
none
```
Response
```
{
    "meta": {
        "message": "Detail of Book",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "name": "Rachmat",
            "gender": "Male",
            "file_avatar": "images/avatars/1-deSims - BuildWith Angga.png"
        }
    ]
}
```

Note : Panduan akan diupdate secara berkala sesuai tahap development
