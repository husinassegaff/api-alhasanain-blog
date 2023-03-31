# api-alhasanain-blog

## Table of Contents
* [Authentication](#authentication)
    1. [Register](#register)
    2. [Login](#login)
    3. [Get All User](#get-all-user)
    4. [Get User By Id](#get-user-by-id)
    5. [Logout](#logout)

* [Post](#post)
    1. [Create Post](#create-post)
    2. [Update Post](#update-post)
    3. [Delete Post](#delete-post)
    4. [Change Status Post](#change-status-post)
    5. [Get All Post](#get-all-post)
    6. [Get Post By Id](#get-post-by-id) 

  
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

### 4. Get User By Id

- URL: `/api/user/get/:id`
- Method: `GET`

**Example success response**
```json
{
  "data": {
    "id_user": "5a3257c7-7c94-4ada-aa2f-639f89372571",
    "name": "husinassegaff",
    "email": "husin@gmail.com",
    "role": "user",
    "token": "AnPLBiUwjxXKSoCOmVS5uzvIbCU5fZj5oYunXwRzja4="
  },
  "message": "success",
  "success": true
}
```
**Example error response**
- ID not found
```json
{
  "error": "user not found",
  "message": "failed",
  "success": false
}
```

### 5. Logout

- URL: `/api/user/logout`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `email`: string

**Example success response**
```json
{
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
```json
{
  "message": "You are not authorized to access this resource",
  "success": false
}
```

## Post

### 1. Create Post

- URL: `/api/post/create`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `id_user`: string
  - `title`: string
  - `content`: string
  - `summary`: string (optional)
  - `status`: string (`draft` || `publish` || `delete`)

**Example success response**
```json
{
  "data": {
    "id_post": "9ed9546f-a34b-4fc5-835a-949557ccfc32",
    "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
    "title": "Lorem Ipsum",
    "meta_title": null,
    "slug": "lorem-ipsum",
    "content": "lorem ipsum dolor sit amet",
    "summary": null,
    "status": "draft",
    "published_at": null,
    "created_at": "2023-03-31T00:00:00Z",
    "updated_at": null
  },
  "message": "success"
}
```
**Example error response**
- title, content, status is required
```json
{
  "message": "title, content, status must be filled",
  "success": false
}
```
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

### 2. Update Post

- URL: `/api/post/update`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `id_user`: string
  - `title`: string (optional)
  - `content`: string (optional)
  - `summary`: string (optional)

**Example success response**
```json
{
  "message": "success",
  "success": true
}
```
**Example error response**
- ID post not found
```json
{
  "message": "Post not found",
  "success": false
}
```

### 3. Delete Post
- URL: `/api/post/delete/:id`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`

**Example success response**
```json
{
  "message": "success",
  "success": true
}
```

### 4. Change Status Post
- URL: `/api/post/delete/:id`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `id_post`: string
  - `status`: string

**Example success response**
```json
{
  "message": "success",
  "success": true
}
```

### 5. Get All Post
- URL: `/api/post/get/all`
- Method: `GET`

**Example success response**
```json
{
  "data": [
    {
      "id_post": "44f88b7e-3151-40bd-8b29-6cb2e330bb7a",
      "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
      "title": "Lorem Ipsum",
      "meta_title": "Lorem-Ipsum",
      "slug": "lorem-ipsum",
      "content": "lorem ipsum dolor sit amet",
      "summary": null,
      "status": "publish",
      "published_at": "2023-03-31T00:00:00Z",
      "created_at": "2023-03-31T00:00:00Z",
      "updated_at": null
    },
    {
      "id_post": "101fd898-a148-4599-9326-db03579bfc30",
      "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
      "title": "Lorem Ipsum",
      "meta_title": "Lorem-Ipsum",
      "slug": "lorem-ipsum",
      "content": "lorem ipsum dolor sit amet",
      "summary": null,
      "status": "draft",
      "published_at": null,
      "created_at": "2023-03-31T00:00:00Z",
      "updated_at": null
    },
    {
      "id_post": "43ac8f03-136a-4243-8bdb-ea51437f8995",
      "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
      "title": "Lorem Ipsum",
      "meta_title": "Lorem-Ipsum",
      "slug": "lorem-ipsum",
      "content": "lorem ipsum dolor sit amet",
      "summary": null,
      "status": "delete",
      "published_at": null,
      "created_at": "2023-03-31T00:00:00Z",
      "updated_at": null
    }
  ],
  "message": "success",
  "success": true
}
```

### 6. Get Post By Id

- URL: `/api/post/get/:id`
- Method: `GET`

**Example success response**
```json
{
  "data": {
    "id_post": "698d3ab1-249b-41e3-b289-d992f0d70e79",
    "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
    "title": "asd Loremasd ASDASD asdasd",
    "meta_title": "asd-Loremasd-ASDASD-asdasd",
    "slug": "asd-loremasd-asdasd-asdasd",
    "content": "lorem ipsum dolor sit amet",
    "summary": null,
    "status": "draft",
    "published_at": null,
    "created_at": "2023-03-31T00:00:00Z",
    "updated_at": "2023-03-31T00:00:00Z"
  },
  "message": "success",
  "success": true
}
```