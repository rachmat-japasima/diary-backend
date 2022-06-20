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
Endpoint
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
Endpoint
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
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNX0.XDQKQijiPi7lN_r6PvKNwfcfXZpf0eSMNqSB9kWv8V0"
    }
}
```

### Email Check
Endpoint ini berguna sebagai validasi pada saat registrasi user untuk mencegah adanya email yang double 

Endpoint
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

Endpoint
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


Note : Panduan akan diupdate secara berkala sesuai tahap development
