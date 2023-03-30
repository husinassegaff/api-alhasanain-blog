# api-alhasanain-blog

## Table of Contents
* [Authentication](#authentication)
    1. [Register](#register)
    2. [Login](#login)
    3. [Get All User](#get-all-user) 

  
## Authentication

### 1. Register

- URL: `/api/user/register`
- Method: `POST`
- Request body:
    - `name`: string
    - `email`: string
    - `password`: string
    - `role`: string (default: `user`)

**Example success response**
```json
{
  "data": {
    "id_user": "4ce80200-949e-43c2-a310-3d69275214f5",
    "name": "husin",
    "email": "husin123@gmail.com",
    "role": "user",
    "token": ""
  },
  "message": "success",
  "success": true
}
```

**Example error response**
- Password must be more than 6 characters
```json
{
  "error": "password must be more than 6 characters",
  "message": "failed",
  "success": false
}
```
- Role must be ` `, `user`  or `admin` (` ` -> empty string)
```json
{
  "error": "role not valid",
  "message": "failed",
  "success": false
}
```
- Email already registered
```json
{
  "error": "email already exists",
  "message": "failed",
  "success": false
}
```

### 2. Login

- URL: `/api/user/login`
- Method: `POST`
- Request body:
  - `email`: string
  - `password`: string

**Example success response**
```json
{
  "data": {
    "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
    "name": "husin",
    "email": "husin456@gmail.com",
    "role": "admin",
    "token": "joQ7TuPtbJYb7v619iulKiD9qrQ1atCt/D6d9EOQ0zY="
  },
  "message": "success",
  "success": true
}
```

**Example error response**
- user already logged in
```json
{
  "error": "user already login",
  "message": "failed",
  "success": false
}
```
- email or password is wrong
```json
{
  "error": "email or password is wrong",
  "message": "failed",
  "success": false
}
```

### 3. Get All User

- URL: `/api/user/get/all`
- Method: `GET`
- Request Header: 
  - `Authorization`: `Bearer <token>`

**Example success response**
```json
{
  "data": [
    {
      "id_user": "5a3257c7-7c94-4ada-aa2f-639f89372571",
      "name": "husinassegaff",
      "email": "husin@gmail.com",
      "role": "user",
      "token": ""
    },
    {
      "id_user": "4ce80200-949e-43c2-a310-3d69275214f5",
      "name": "husin",
      "email": "husin123@gmail.com",
      "role": "user",
      "token": ""
    },
    {
      "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
      "name": "husin",
      "email": "husin456@gmail.com",
      "role": "admin",
      "token": "hXgRmV1IZLyNnuXaJ3Ah/yOoJkWQVDMt2OvNhF3KybM="
    }
  ],
  "message": "success",
  "success": true
}
```

**Example error response**
- authorization not found
```json
{
    "message": "Authorization header required",
    "success": false
}
```
- token is invalid
```json
{
    "message": "Invalid token",
    "success": false
}
```
- token from role user
```json
{
    "message": "You are not authorized to access this resource",
    "success": false
}
```