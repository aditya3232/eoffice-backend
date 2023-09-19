# API Spec E-office BIT

# 1 Account

## 1.1 Auth

### 1.1.1 Login

Request :
- Method : POST
- Endpoint : `/api/v1/auth/login`
- Header :
    - Content-Type : multipart/form-data
- Body (form-data) :
    - nip : integer, required,
    - password : string, required,

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "auth": {
            "expired": "string",
            "token": "string"
        },
        "user": {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "employee_id": "integer",
            "password": "string",
            "token": "string",
            "role_id": "integer",
            "last_login": "string"
        }
    }
}
```
### 1.1.2 Logout

Request :
- Method : POST
- Endpoint : `/api/v1/auth/logout`
- Header :
    - Content-Type : multipart/form-data,
    - Authorization : token

Response :

```json 
{
    "meta" : {
        "message": "string",
        "code": "integer"
    },
    "pagination" : {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data" : null
}
```
## 1.2 Profile
### 1.2.1 Get Profile

Request :
- Method : POST
- Endpoint : `/api/v1/profile`
- Header :
    - Content-Type : multipart/form-data
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "user": {
             "id": 1,
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "employee_id": "integer",
            "password": "string",
            "token": "string",
            "role_id": "integer",
            "last_login": "string"
        },
        "employees": {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "nama": "string",
            "nip": "integer",
            "tempat_lahir": "string",
            "tanggal_lahir": "string",
            "alamat": "string",
            "no_hp": "string",
            "email_personal": "string",
            "email_corporate": "string",
            "division_id": "integer",
            "position_id": "integer",
            "start_date": "string",
            "end_date": "string",
            "avatar": "string"
        }
    }
}
```
### 1.2.2 Get Permission

Request :
- Method : POST
- Endpoint : `/api/v1/profile/permission`
- Header :
    - Content-Type : multipart/form-data
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": [
        {
            "id": "integer",
            "parent_id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "nama": "string",
            "url": "string",
            "position": "integer"
        }
    ]
}
```

# 2 Master Data

## 2.1 User

### 2.1.1 Get All

Request :
- Method : GET
- Endpoint : `/api/v1/user`
- Header :
    - Authorization : token
- Params (Query Params) :
    - limit : integer, required,
    - page : string, required,
    - sort : string, required,
    - order : string, required,
    - id : intger,
    - roles_id : integer,
    - employees_id : integer,
    - last_login : string,
    - remarks : string,
    - updated_at : string,

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": [
        {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "employee_id": "integer",
            "password": "string",
            "token": "string",
            "role_id": "integer",
            "last_login": "string"
        }
    ]
}
```
### 2.1.2 Get One

Request :
- Method : GET
- Endpoint : `/api/v1/user/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "employee_id": "integer",
        "password": "string",
        "token": "string",
        "role_id": "integer",
        "last_login": "string"
    }
    
}
```
### 2.1.3 Create

Request :
- Method : POST
- Endpoint : `/api/v1/user`
- Header :
    - Authorization : token
- Body (form-data) :
    - employee_id : integer, required,
    - password : string, required,
    - roles_id : integer, required

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "employee_id": "integer",
        "password": "string",
        "token": "string",
        "role_id": "integer",
        "last_login": "string"
    }
    
}
```
### 2.1.4 Update

Request :
- Method : PUT
- Endpoint : `/api/v1/user/{:id}`
- Header :
    - Authorization : token
- Body (form-data) :
    - remarks : string,
    - password : string,
    - roles_id : integer

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "employee_id": "integer",
        "password": "string",
        "token": "string",
        "role_id": "integer",
        "last_login": "string"
    }
    
}
```

### 2.1.5 Delete

Request :
- Method : DEL
- Endpoint : `/api/v1/user/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": null
    
}
```
## 2.2 Employee

### 2.2.1 Get All

Request :
- Method : GET
- Endpoint : `/api/v1/employee`
- Header :
    - Authorization : token
- Params (Query Params) :
    - limit : integer, required,
    - page : string, required,
    - sort : string, required,
    - order : string, required,
    - id : intger,
    - created_at : string,
    - updated_at : string,
    - deleted_at : string,
    - remarks" : string,
    - nip : integer,
    - nama : string,
    - tempat_lahir : string,
    - alamat : string,
    - no_hp : string,
    - email_personal : string,
    - email_corporate : string,
    - division_id : integer,
    - position_id : integer,
    - start_date : string,
    - end_date : string,
    - avatar : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": [
        {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "nama": "string",
            "nip": "integer",
            "tempat_lahir": "string",
            "tanggal_lahir": "string",
            "alamat": "string",
            "no_hp": "integer",
            "email_personal": "string",
            "email_corporate": "string",
            "division_id": "integer",
            "position_id": "integer",
            "start_date": "string",
            "end_date": "string",
            "avatar": "string"
        }
    ]
}
```
### 2.2.2 Get One

Request :
- Method : GET
- Endpoint : `/api/v1/employee/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string",
        "nip": "integer",
        "tempat_lahir": "string",
        "tanggal_lahir": "string",
        "alamat": "string",
        "no_hp": "integer",
        "email_personal": "string",
        "email_corporate": "string",
        "division_id": "integer",
        "position_id": "integer",
        "start_date": "string",
        "end_date": "string",
        "avatar": "string"
    }
    
}
```
### 2.2.3 Create

Request :
- Method : POST
- Endpoint : `/api/v1/employee`
- Header :
    - Authorization : token
- Body (form-data) :
    - nama : string, required,
    - nip : string, required,
    - tempat_lahir : string, required,
    - tanggal_lahir : string, required,
    - alamat : string, required,
    - no_hp : integer, required,
    - email_personal : string, required,
    - email_corporate : string, required,
    - division_id : integer, required,
    - position_id : integer, required,
    - start_date : string, required,
    - end_date : string, required,
    - avatar : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string",
        "nip": "integer",
        "tempat_lahir": "string",
        "tanggal_lahir": "string",
        "alamat": "string",
        "no_hp": "integer",
        "email_personal": "string",
        "email_corporate": "string",
        "division_id": "integer",
        "position_id": "integer",
        "start_date": "string",
        "end_date": "string",
        "avatar": "string"
    }
    
}
```
### 2.2.4 Update

Request :
- Method : PUT
- Endpoint : `/api/v1/employee/{:id}`
- Header :
    - Authorization : token
- Body (form-data) :
    - nama : string,
    - nip : string,
    - tempat_lahir : string,
    - tanggal_lahir : string,
    - alamat : string,
    - no_hp : integer,
    - email_personal : string,
    - email_corporate : string,
    - division_id : integer,
    - position_id : integer,
    - start_date : string,
    - end_date : string,
    - avatar : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string",
        "nip": "integer",
        "tempat_lahir": "string",
        "tanggal_lahir": "string",
        "alamat": "string",
        "no_hp": "integer",
        "email_personal": "string",
        "email_corporate": "string",
        "division_id": "integer",
        "position_id": "integer",
        "start_date": "string",
        "end_date": "string",
        "avatar": "string"
    }
    
}
```

### 2.2.5 Delete

Request :
- Method : DEL
- Endpoint : `/api/v1/employee/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": null
    
}
```
## 2.3 Role

### 2.3.1 Get All

Request :
- Method : GET
- Endpoint : `/api/v1/role`
- Header :
    - Authorization : token
- Params (Query Params) :
    - limit : integer, required,
    - page : string, required,
    - sort : string, required,
    - order : string, required,
    - id : intger,
    - created_at : string,
    - updated_at : string,
    - deleted_at : string,
    - remarks : string,
    - nama : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": [
        {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "nama": "string"
        }
    ]
}
```
### 2.3.2 Get One

Request :
- Method : GET
- Endpoint : `/api/v1/role/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string"
    }
    
}
```
### 2.3.3 Create

Request :
- Method : POST
- Endpoint : `/api/v1/role`
- Header :
    - Authorization : token
- Body (form-data) :
    - nama : string, required,

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string"
    }
    
}
```
### 2.3.4 Update

Request :
- Method : PUT
- Endpoint : `/api/v1/role/{:id}`
- Header :
    - Authorization : token
- Body (form-data) :
    - nama : string,
    - remarks : string,

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string"
    }
    
}
```

### 2.3.5 Delete

Request :
- Method : DEL
- Endpoint : `/api/v1/role/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": null
    
}
```

## 2.4 Permission

### 2.4.1 Get All

Request :
- Method : GET
- Endpoint : `/api/v1/permission`
- Header :
    - Authorization : token
- Params (Query Params) :
    - limit : integer, required,
    - page : string, required,
    - sort : string, required,
    - order : string, required,
    - id : integer,
    - created_at : string,
    - updated_at : string,
    - deleted_at : string,
    - remarks : string,
    - parent_id : integer,

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": [
        {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "parent_id": "integer",
            "nama": "string",
            "url": "string",
            "position": "integer"
        }
    ]
}
```
### 2.4.2 Get One

Request :
- Method : GET
- Endpoint : `/api/v1/permission/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "parent_id": "integer",
        "nama": "string",
        "url": "string",
        "position": "integer"
    }
    
}
```
### 2.4.3 Create

Request :
- Method : POST
- Endpoint : `/api/v1/permission`
- Header :
    - Authorization : token
- Body (form-data) :
    - parent_id : integer, required,
    - nama : string, required,
    - url : string, required,
    - position : integer, required

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "parent_id": "integer",
        "nama": "string",
        "url": "string",
        "position": "integer"
    }
    
}
```
### 2.4.4 Update

Request :
- Method : PUT
- Endpoint : `/api/v1/permission/{:id}`
- Header :
    - Authorization : token
- Body (form-data) :
    - parent_id : integer,
    - remarks : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "parent_id": "integer",
        "nama": "string",
        "url": "string",
        "position": "integer"
    }
    
}
```

### 2.4.5 Delete

Request :
- Method : DEL
- Endpoint : `/api/v1/permission/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": null
    
}
```

## 2.5 Role Permission

### 2.5.1 Get All

Request :
- Method : GET
- Endpoint : `/api/v1/role-permission`
- Header :
    - Authorization : token
- Params (Query Params) :
    - limit : integer, required,
    - page : string, required,
    - sort : string, required,
    - order : string, required,
    - id : integer,
    - created_at : string,
    - updated_at : string,
    - deleted_at : string,
    - remarks : string,
    - role_id : integer,
    - permission_id : integer

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": [
        {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "role_id": "integer",
            "permission_id": "integer"
        }
    ]
}
```
### 2.5.2 Get One

Request :
- Method : GET
- Endpoint : `/api/v1/role-permission/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "role_id": "integer",
        "permission_id": "integer"
    }
    
}
```
### 2.5.3 Create

Request :
- Method : POST
- Endpoint : `/api/v1/role-permission`
- Header :
    - Authorization : token
- Body (form-data) :
    - role_id : integer, required,
    - permission_id : integer, required,

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "role_id": "integer",
        "permission_id": "integer"
    }
    
}
```
### 2.5.4 Update

Request :
- Method : PUT
- Endpoint : `/api/v1/role-permission/{:id}`
- Header :
    - Authorization : token
- Body (form-data) :
    - role_id : integer,
    - permission_id : integer,
    - remarks : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "role_id": "integer",
        "permission_id": "integer"
    }
    
}
```

### 2.5.5 Delete

Request :
- Method : DEL
- Endpoint : `/api/v1/role-permission/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": null
    
}
```

## 2.6 Division

### 2.6.1 Get All

Request :
- Method : GET
- Endpoint : `/api/v1/division`
- Header :
    - Authorization : token
- Params (Query Params) :
    - limit : integer, required,
    - page : string, required,
    - sort : string, required,
    - order : string, required,
    - id : integer,
    - created_at : string,
    - updated_at : string,
    - deleted_at : string,
    - remarks : string,
    - parent_id : integer,
    - nama : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": [
        {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "parent_id": "integer",
            "nama": "string"
        }
    ]
}
```
### 2.6.2 Get One

Request :
- Method : GET
- Endpoint : `/api/v1/division/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "parent_id": "integer",
        "nama": "string"
    }
    
}
```
### 2.6.3 Create

Request :
- Method : POST
- Endpoint : `/api/v1/division`
- Header :
    - Authorization : token
- Body (form-data) :
    - parent_id : integer, required,
    - nama : string, required,

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "parent_id": "integer",
        "nama": "string"
    }
    
}
```
### 2.6.4 Update

Request :
- Method : PUT
- Endpoint : `/api/v1/division/{:id}`
- Header :
    - Authorization : token
- Body (form-data) :
    - parent_id : integer,
    - nama : string,
    - remarks : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "parent_id": "integer",
        "nama": "string"
    }
    
}
```

### 2.6.5 Delete

Request :
- Method : DEL
- Endpoint : `/api/v1/division/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": null
    
}
```

## 2.7 Position

### 2.7.1 Get All

Request :
- Method : GET
- Endpoint : `/api/v1/position`
- Header :
    - Authorization : token
- Params (Query Params) :
    - limit : integer, required,
    - page : string, required,
    - sort : string, required,
    - order : string, required,
    - id : integer,
    - created_at : string,
    - updated_at : string,
    - deleted_at : string,
    - remarks : string,
    - nama : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": [
        {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "nama": "string"
        }
    ]
}
```
### 2.7.2 Get One

Request :
- Method : GET
- Endpoint : `/api/v1/position/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string"
    }
    
}
```
### 2.7.3 Create

Request :
- Method : POST
- Endpoint : `/api/v1/position`
- Header :
    - Authorization : token
- Body (form-data) :
    - nama : string, required

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string"
    }
    
}
```
### 2.7.4 Update

Request :
- Method : PUT
- Endpoint : `/api/v1/position/{:id}`
- Header :
    - Authorization : token
- Body (form-data) :
    - nama : string,
    - remarks : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string"
    }
    
}
```

### 2.7.5 Delete

Request :
- Method : DEL
- Endpoint : `/api/v1/position/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": null
    
}
```

## 2.8 Location

### 2.8.1 Get All

Request :
- Method : GET
- Endpoint : `/api/v1/location`
- Header :
    - Authorization : token
- Params (Query Params) :
    - limit : integer, required,
    - page : string, required,
    - sort : string, required,
    - order : string, required,
    - id : integer,
    - created_at : string,
    - updated_at : string,
    - deleted_at : string,
    - remarks : string,
    - nama : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": [
        {
            "id": "integer",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
            "remarks": "string",
            "nama": "string"
        }
    ]
}
```
### 2.8.2 Get One

Request :
- Method : GET
- Endpoint : `/api/v1/location/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string"
    }
    
}
```
### 2.8.3 Create

Request :
- Method : POST
- Endpoint : `/api/v1/location`
- Header :
    - Authorization : token
- Body (form-data) :
    - nama : string, required

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string"
    }
    
}
```
### 2.8.4 Update

Request :
- Method : PUT
- Endpoint : `/api/v1/location/{:id}`
- Header :
    - Authorization : token
- Body (form-data) :
    - nama : string,
    - remarks : string

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": {
        "id": "integer",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": "string",
        "remarks": "string",
        "nama": "string"
    }
    
}
```

### 2.8.5 Delete

Request :
- Method : DEL
- Endpoint : `/api/v1/location/{:id}`
- Header :
    - Authorization : token

Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer"
    },
    "pagination": {
        "page": "integer",
        "limit": "integer",
        "total": "integer",
        "total_filtered": "integer",
    },
    "data": null
    
}
```







