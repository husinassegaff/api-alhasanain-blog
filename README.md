# api-alhasanain-blog

## Table of Contents
* [Authentication](#authentication)
    1. [Register](#register)

  
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