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

* [Category](#category)
    1. [Create Category](#create-category)
    2. [Update Category](#update-category)
    3. [Delete Category](#delete-category)
    4. [Get All Category](#get-all-category)
    5. [Get Category By Id](#get-category-by-id)

* [Meta](#meta)
  1. [Create Meta](#create-meta)
  2. [Update Meta](#update-meta)
  3. [Delete Meta](#delete-meta)
  4. [Get All Meta](#get-all-meta)
  5. [Get Meta By Id](#get-meta-by-id)

* [Tag](#tag)
  1. [Create Tag](#create-tag)
  2. [Delete Tag](#delete-tag)
  3. [Get All Tag](#get-all-tag)

* [Comment](#comment)
  1. [Create Comment](#create-comment)
  2. [Delete Comment](#delete-comment)
  3. [Get All Comment](#get-all-comment)

  
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

## Category

### 1. Create Category

- URL: `/api/category/create`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `title`: string
  - `content`: string

**Example success response**
```json
{
  "data": {
    "id_category": "c1b61edc-e281-4fa2-8fc5-bf9f747c7c53",
    "title": "Bahasa Arab",
    "meta_title": "Bahasa-Arab",
    "slug": "bahasa-arab",
    "content": "Topik yang berkaitan dengan bahasa arab",
    "created_at": "2023-03-31T18:35:30.4788098+07:00",
    "updated_at": null
  },
  "message": "Category created",
  "success": true
}
```
**Example error response**
```json
{
  "message": "Failed to create category",
  "success": false
}
```

### 2. Update Category

- URL: `/api/category/update`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `id_category`: string
  - `title`: string (optional)
  - `content`: string (optional)

**Example success response**
```json
{
  "message": "Category updated",
  "success": true
}
```
**Example error response**
- ID post not found
```json
{
  "message": "Failed to update category",
  "success": false
}
```

### 3. Delete Category
- URL: `/api/category/delete/:id`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`

**Example success response**
```json
{
  "message": "Category deleted",
  "success": true
}
```

### 4. Get All Category
- URL: `/api/category/get/all`
- Method: `GET`

**Example success response**
```json
{
  "data": [
    {
      "id_category": "a2d0e26b-c0e4-4f21-b72f-6730877e8960",
      "title": "Bahasa Arab",
      "meta_title": "Bahasa-Arab",
      "slug": "bahasa-arab",
      "content": "Topik yang berkaitan dengan bahasa arab",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": null
    },
    {
      "id_category": "a22f1530-bf3e-4d68-b6f7-3e60774ee1e6",
      "title": "Fiqih",
      "meta_title": "Fiqih",
      "slug": "fiqih",
      "content": "Topik yang berkaitan dengan fiqih",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": null
    }
  ],
  "message": "success",
  "success": true
}
```

### 5. Get Category By Id

- URL: `/api/category/get/:id`
- Method: `GET`

**Example success response**
```json
{
  "data": {
    "id_category": "a22f1530-bf3e-4d68-b6f7-3e60774ee1e6",
    "title": "Fiqih",
    "meta_title": "Fiqih",
    "slug": "fiqih",
    "content": "Topik yang berkaitan dengan fiqih",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": null
  },
  "message": "success",
  "success": true
}
```


## Meta

### 1. Create Meta

- URL: `/api/meta/create`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `key`: string
  - `content`: string

**Example success response**
```json
{
  "data": {
    "id_meta": "ef1e9957-22a8-4fe2-961a-d61bf5a9b9c0",
    "key": "image",
    "content": "https://avatars.githubusercontent.com/u/63222585?s=96&v=4",
    "created_at": "2023-03-31T19:44:25.8421922+07:00",
    "updated_at": null
  },
  "message": "Meta created",
  "success": true
}
```

### 2. Update Meta

- URL: `/api/meta/update`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `id_meta`: string
  - `key`: string (optional)
  - `content`: string (optional)

**Example success response**
```json
{
  "message": "Meta updated",
  "success": true
}
```
**Example error response**
- ID post not found
```json
{
  "message": "Failed to update meta",
  "success": false
}
```

### 3. Delete Meta
- URL: `/api/meta/delete/:id`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`

**Example success response**
```json
{
  "message": "Meta deleted",
  "success": true
}
```

### 4. Get All Meta
- URL: `/api/meta/get/all`
- Method: `GET`
- Request Header:
  - Authorization: `Bearer <token>`

**Example success response**
```json
{
  "data": [
    {
      "id_meta": "af62d501-ce16-4228-896f-f8754faff78e",
      "key": "image",
      "content": "https://avatars.githubusercontent.com/u/63222585?s=96&v=4",
      "created_at": "2023-03-31T00:00:00Z",
      "updated_at": null
    },
    {
      "id_meta": "57dec8c5-ae8a-4afc-8eea-86873a678134",
      "key": "image",
      "content": "https://avatars.githubusercontent.com/u/63222585?s=96&v=4",
      "created_at": "2023-03-31T00:00:00Z",
      "updated_at": null
    },
    {
      "id_meta": "80b9d857-c92d-407b-988d-2978de38aba6",
      "key": "image",
      "content": "https://avatars.githubusercontent.com/u/63222585?s=96&v=4",
      "created_at": "2023-03-31T00:00:00Z",
      "updated_at": null
    }
  ],
  "message": "Meta fetched",
  "success": true
}
```

### 5. Get Meta By Id

- URL: `/api/meta/get/:id`
- Method: `GET`

**Example success response**
```json
{
  "data": {
    "id_meta": "af62d501-ce16-4228-896f-f8754faff78e",
    "key": "image",
    "content": "https://avatars.githubusercontent.com/u/63222585?s=96&v=4",
    "created_at": "2023-03-31T00:00:00Z",
    "updated_at": null
  },
  "message": "Meta fetched",
  "success": true
}
```

## Tag

### 1. Create Tag

- URL: `/api/tag/create`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `id_post`: string
  - `title`: string

**Example success response**
```json
{
  "data": {
    "id_tag": "2d839637-738d-4f64-bdc3-7261c2119f74",
    "id_post": "101fd898-a148-4599-9326-db03579bfc30",
    "title": "fiqih",
    "meta_title": "fiqih",
    "created_at": "2023-03-31T21:30:57.7749509+07:00",
    "updated_at": null
  },
  "message": "Successfully create tag",
  "success": true
}
```

### 2. Delete Tag
- URL: `/api/tag/delete/:id`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `id_post`: string

**Example success response**
```json
{
  "message": "Successfully delete tag",
  "success": true
}
```

### 3. Get All Tag
- URL: `/api/tag/get/all`
- Method: `GET`

**Example success response**
```json
{
  "data": [
    {
      "id_tag": "dca50762-61db-4594-b194-37e7ad16e7d0",
      "id_post": "101fd898-a148-4599-9326-db03579bfc30",
      "title": "fiqih",
      "meta_title": "fiqih",
      "created_at": "2023-03-31T00:00:00Z",
      "updated_at": null
    },
    {
      "id_tag": "3ff8fb5f-cbb0-4e40-a87b-e4251ccb41a9",
      "id_post": "101fd898-a148-4599-9326-db03579bfc30",
      "title": "fiqih",
      "meta_title": "fiqih",
      "created_at": "2023-03-31T00:00:00Z",
      "updated_at": null
    },
    {
      "id_tag": "9d55e549-da68-4c78-9db4-1d3f382119f8",
      "id_post": "101fd898-a148-4599-9326-db03579bfc30",
      "title": "fiqih",
      "meta_title": "fiqih",
      "created_at": "2023-03-31T00:00:00Z",
      "updated_at": null
    }
  ],
  "message": "Successfully get all tag",
  "success": true
}
```

## Comment

### 1. Create Comment

- URL: `/api/comment/create`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `id_user`: string
  - `id_post`: string
  - `content`: string

**Example success response**
```json
{
  "data": {
    "id_comment": "ea7e1e8a-35fb-43aa-9b81-25e36dc62e90",
    "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
    "id_post": "101fd898-a148-4599-9326-db03579bfc30",
    "content": "lorem ipsum",
    "created_at": "2023-03-31T21:12:37.3626484+07:00"
  },
  "message": "Comment created",
  "success": true
}
```

### 2. Delete Comment
- URL: `/api/comment/delete/:id`
- Method: `POST`
- Request Header:
  - Authorization: `Bearer <token>`
- Request body:
  - `id_user`: string

**Example success response**
```json
{
  "message": "Comment deleted",
  "success": true
}
```

### 3. Get All Comment
- URL: `/api/comment/get/all`
- Method: `GET`

**Example success response**
```json
{
  "data": [
    {
      "id_comment": "17cf09da-85b3-4077-ada4-fa7823d1e221",
      "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
      "id_post": "101fd898-a148-4599-9326-db03579bfc30",
      "content": "lorem ipsum",
      "created_at": "2023-03-31T00:00:00Z"
    },
    {
      "id_comment": "aee54e2a-b8db-4ba7-a341-0def781fbb93",
      "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
      "id_post": "101fd898-a148-4599-9326-db03579bfc30",
      "content": "lorem ipsum",
      "created_at": "2023-03-31T00:00:00Z"
    },
    {
      "id_comment": "b505d9fe-92eb-4d17-bdb2-0ab7093c6b9a",
      "id_user": "78fc0eba-fbdc-4eb6-be19-4c6ce5b4624f",
      "id_post": "101fd898-a148-4599-9326-db03579bfc30",
      "content": "lorem ipsum",
      "created_at": "2023-03-31T00:00:00Z"
    }
  ],
  "message": "All comments",
  "success": true
}
```