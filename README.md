# Mini Project Go Buku API developed by using Go Gin, Gorm and PostgreSQL
This is a simple restful api project that deployed to heroku. You can access it with this url:  
```
https://go-buku.herokuapp.com
```
## Project Overview =
- CRUD API Books using Gin, Gorm and PostgreSQL
- This Project has 3 domains :
    a. books => for Books Data 
    b. user => for Get Data User and Update/Edit Data User
    c. auth => for Login and Register User
- Clean Design Architecture
- JWT Token
- Hash and Salt Password
- Pagination
- Unit Testing

## PostgreSQL

SQL implementation in this project :

### users
```sql
CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "password" TEXT NOT NULL,
);
```

### books
```sql
CREATE TABLE "books" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "title" VARCHAR(255) NOT NULL,
  "description" TEXT NOT NULL,
  "author" TEXT NOT NULL,
  "price" INT NOT NULL
);

INSERT INTO "books" (title, description, author, price, user_id)
VALUES
	('Rich Dad Poor Dad','Good book for investing','Robert Kiyosaki',1, 120000),
	('Pemrograman JavaScript untuk pemula sampai mahir','Buku untuk mempelajari bahasa pemrograman JavaScript','Eko Kurniawan Khannedy',1, 95000),
	('The Subtle Art Of Not Giving a Fuck','Buku tentang pengembangan diri','Mark Manson',1, 80000),
	('Atomic Habit','Self development book that very recommended','James Clear',1, 135000)
	('Grit','Buku tentang motivasi diri','Angela Duckworth',1, 220000)
;
```


# Documentation

## register
Registering a new user

<b>POST</b>
```
https://go-buku.herokuapp.com/api/auth/register
```

Request Body
```
{
    "name": "Putra Fajar F",
    "email": "putra@gmail.com",
    "password": "putraf"
}
```

Response success (Status: 201)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Putra Fajar F",
        "email": "putra@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsImV4cCI6MTY2MDAwNjc4NCwiaWF0IjoxNjU5OTg1MTg0LCJpc3MiOiJhZGFpc3N1ZSJ9.UBCOMytjBsyfsuZi0hG6A-pfTv5CHzpRII9NQvIVjkE"
    }
}
```

Response error (Status : 400) [Depends on what error]
```
{
    "status": false,
    "message": "Failed to process request",
    "errors": [
        "Key: 'RegisterRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
    ],
    "data": {}
}
```

## login
Authenticate by email and password

<b>POST</b>
```
https://go-buku.herokuapp.com/api/auth/login
```

Request body
```
{
    "email": "putra@gmail.com",
    "password": "putraf"
}
```

Response Success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Putra Fajar F",
        "email": "putra@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsImV4cCI6MTY2MDAwNjk1MCwiaWF0IjoxNjU5OTg1MzUwLCJpc3MiOiJhZGFpc3N1ZSJ9.9YTRaEQeP_8SWbjwLvJfbwGConuYLUHR0MKuLKvgr9A"
    }
}
```

Response error, wrong credential (Status: 401)
```
{
    "status": false,
    "message": "Failed to login",
    "errors": [
        "failed to login. check your credential"
    ],
    "data": {}
}
```


## Get Profile
Get current info from logged user

<b>GET</b>
```
https://go-buku.herokuapp.com/api/user/profile
```

Headers
```
Authorization: yourToken
```

Response success (status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Putra Fajar F",
        "email": "putra@gmail.com"
    }
}
```

## Update profile
Update user data who logged in

<b>PUT</b>
```
https://go-buku.herokuapp.com/api/user/profile
```

Headers
```
Authorization: yourToken
```

Request Body
```
{
    "name": "Putra Fajar Febrianto",
    "email": "putra@gmail.com"
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Putra Fajar Febrianto",
        "email": "putra@gmail.com"
    }
}
```


<b>=============================================</b>
## All books (based on user who logged in)
Only shows books by user who logged in

<b>GET with Default Endpoint</b>
```
https://go-buku.herokuapp.com/api/books
```


Headers
```
Authorization: yourToken
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "limit": 5,
        "page": 0,
        "total_rows": 5,
        "from_row": 1,
        "to_row": 0,
        "rows": [
            {
                "id": 1,
                "title": "Rich Dad Poor Dad",
                "description": "Good book for investing",
                "author": "Robert Kiyosaki",
                "price": "120000",
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            },
            {
                "id": 2,
                "title": "Pemrograman JavaScript untuk pemula sampai mahir",
                "description": "Buku untuk mempelajari bahasa pemrograman JavaScript",
                "author": "Eko Kurniawan Khannedy",
                "price": "95000",
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            },
            {
                "id": 3,
                "title": "The Subtle Art Of Not Giving a Fuck",
                "description": "Buku tentang pengembangan diri",
                "author": "Mark Manson",
                "price": "80000",
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            },
            {
                "id": 4,
                "title": "Atomic Habit",
                "description": "Self development book that very recommended",
                "author": "James Clear",
                "price": "135000",
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            },
            {
                "id": 5,
                "title": "Grit",
                "description": "Buku tentang motivasi diri",
                "author": "Angela Duckworth",
                "price": "220000",
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            }
        ]
    }
}
```

<b>GET with Pagination End Point</b>
```
https://go-buku.herokuapp.com/api/books?page=0&limit=6
```


Headers
```
Authorization: yourToken
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "limit": 6,
        "page": 0,
        "total_rows": 15,
        "from_row": 1,
        "to_row": 0,
        "rows": [
            {
                "id": 1,
                "title": "Rich Dad Poor Dad New Version",
                "description": "Good book for investing",
                "author": "Robert Kiyokasi",
                "price": 145000,
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            },
            {
                "id": 2,
                "title": "Pemrograman JavaScript untuk pemula sampai mahir",
                "description": "Buku untuk mempelajari bahasa pemrograman JavaScript",
                "author": "Eko Kurniawan Khannedy",
                "price": 95000,
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            },
            {
                "id": 3,
                "title": "The Subtle Art Of Not Giving a Fuck",
                "description": "Buku tentang pengembangan diri",
                "author": "Mark Manson",
                "price": 80000,
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            },
            {
                "id": 4,
                "title": "Atomic Habit",
                "description": "Self development book that very recommended",
                "author": "James Clear",
                "price": 135000,
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            },
            {
                "id": 5,
                "title": "Grit",
                "description": "Buku tentang motivasi diri",
                "author": "Angela Duckworth",
                "price": 220000,
                "user": {
                    "id": 1,
                    "name": "Putra Fajar Febrianto",
                    "email": "putra@gmail.com"
                }
            },
            {
                "id": 6,
                "title": "Eat That Frog",
                "description": "Buku tentang self development",
                "author": "Brian Tracy",
                "price": 125000,
                "user": {
                    "id": 2,
                    "name": "Bambang Hartono",
                    "email": "bambang@gmail.com"
                }
            }
        ]
    }
}
```

## Create books
Create a books with owner is the user who logged in

<b>POST</b>
```
https://go-buku.herokuapp.com/api/books
```

Headers
```
Authorization: yourToken
```

Request body
```
{
    "title":"Eat That Frog",
    "description":"Buku tentang self development",
    "author":"Brian Tracy",
    "price": 125000
}
```

Response success (Status: 201)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 6,
        "title": "Eat That Frog",
        "description": "Buku tentang self development",
        "author": "Brian Tracy",
        "price": 125000,
        "user": {
            "id": 2,
            "name": "Bambang Hartono",
            "email": "bambang@gmail.com"
        }
    }
}
```

## Find one books by id
Find books by id

<b>GET</b>
```
https://go-buku.herokuapp.com/api/books/{id}
```

Headers
```
Authorization: yourToken
```


Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "title": "Rich Dad Poor Dad",
        "description": "Good book for investing",
        "author": "Robert Kiyosaki",
        "price": 120000,
        "user": {
            "id": 1,
            "name": "Putra Fajar Febrianto",
            "email": "putra@gmail.com"
        }
    }
}
```

## Update books
<b>You can only update your own books If you are trying to update someone else books, it will return error. </b>  

<b>PUT</b>
```
https://go-buku.herokuapp.com/api/books/{id}
```

Request body
```
{
    "id":1,
    "title":"Rich Dad Poor Dad New Version",
    "description":"Good book for investing",
    "author":"Robert Kiyokasi",
    "price":145000
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "title": "Rich Dad Poor Dad New Version",
        "description": "Good book for investing",
        "author": "Robert Kiyokasi",
        "price": 145000,
        "user": {
            "id": 1,
            "name": "Putra Fajar Febrianto",
            "email": "putra@gmail.com"
        }
    }
}
```


## Delete books
You can only delete your own books

<b>DELETE</b>
```
https://go-buku.herokuapp.com/api/books/{id}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "Deleted",
    "errors": null,
    "data": {}
}
```