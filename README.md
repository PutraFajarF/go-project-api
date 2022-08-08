# golang heroku deploy
This is a simple restful api project that deployed to heroku. You can access it with this url:  
```
https://go-buku.herokuapp.com
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
        "id": 2,
        "name": "Putra Fajar F",
        "email": "putra@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsImV4cCI6MTY1MTgyMDAwMCwiaWF0IjoxNjIwMjg0MDAwLCJpc3MiOiJhZG1pbiJ9.HtnuWlBaevEO3fHAI4McH5W8axvw_3Og47RUI3m9IyI"
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
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsImV4cCI6MTY1MTgyMTQ0NiwiaWF0IjoxNjIwMjg1NDQ2LCJpc3MiOiJhZG1pbiJ9.2m-r1qrCXhNkAxzK-sb4hL0Pzat3zwOmzktES_uzwts"
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
    "name": "Putra Fajar F",
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
        "name": "Putra Fajar F",
        "email": "putra@gmail.com"
    }
}
```


<b>=============================================</b>
## All products (based on user who logged in)
Only shows products by user who logged in

<b>GET</b>
```
https://go-buku.herokuapp.com/api/product
```


Headers
```
Authorization: yourToken
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": [
        {
            "id": 1,
            "title": "Rich Dad Poor Dad",
            "description": "Book for investing by robert kiyosaki",
            "user": {
                "id": 1,
                "name": "Putra Fajar F",
                "email": "putra@gmail.com"
            }
        },
        {
            "id": 2,
            "title": "JavaScript Pemula Sampai Mahir",
            "description": "Buku untuk belajar javascript by eko kurniawan khannedy",
            "user": {
                "id": 1,
                "name": "Putra Fajar F",
                "email": "putra@gmail.com"
            }
        }
    ]
}
```

## Create product
Create a product with owner is the user who logged in

<b>POST</b>
```
https://go-buku.herokuapp.com/api/product
```

Headers
```
Authorization: yourToken
```

Request body
```
{
    "title": "Rich Dad Poor Dad",
    "description": "Book for investing by robert kiyosaki"
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
        "title": "Rich Dad Poor Dad",
        "description": "Book for investing by robert kiyosaki",
        "user": {
            "id": 1,
            "name": "Putra Fajar F",
            "email": "putra@gmail.com"
        }
    }
}
```

## Find one product by id
Find product by id

<b>GET</b>
```
https://go-buku.herokuapp.com/api/product/{id}
```

Headers
```
Authorization: yourToken
```


Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "title": "Rich Dad Poor Dad",
        "description": "Book for investing by robert kiyosaki",
        "user": {
            "id": 1,
            "name": "Putra Fajar F",
            "email": "putra@gmail.com"
        }
    }
}
```

## Update product
<b>You can only update your own product If you are trying to update someone else product, it will return error. </b>  

<b>PUT</b>
```
https://go-buku.herokuapp.com/api/product/{id}
```

Request body
```
{
    "title": "Rich Dad Poor Dad Version 2",
    "description": "Book for investing by robert kiyosaki"
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "title": "Rich Dad Poor Dad Version 2",
        "description": "Book for investing by robert kiyosaki",
        "user": {
            "id": 1,
            "name": "Putra Fajar F",
            "email": "putra@gmail.com"
        }
    }
}
```


## Delete product
You can only delete your own product

<b>DELETE</b>
```
https://go-buku.herokuapp.com/api/product/{id}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {}
}
```