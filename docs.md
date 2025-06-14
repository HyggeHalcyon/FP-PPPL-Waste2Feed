# User

## POST `/api/user/register`
### With Email
Request:
```json
{
    "email": "test@test.com",
    "password": "password"
}
```
Response:
```json
{
    "status": true,
    "message": "success register user",
    "data": {
        "id": "b5ae5963-ea66-4389-ae67-511d880c4cac",
        "username": "",
        "name": "",
        "address": "",
        "email": "test@test.com",
        "phone_number": "",
        "role": "user",
        "verified": false
    }
}
```
### With Phone Number
Request:
```json
{
    "phone_number": "081215126272",
    "password": "password"
}
```
Response:
```json
{
    "status": true,
    "message": "success register user",
    "data": {
        "id": "97c6bd6f-44ed-4f66-8ba6-485b33ed6b16",
        "username": "",
        "name": "",
        "address": "",
        "email": "",
        "phone_number": "081215126272",
        "role": "user",
        "verified": false
    }
}
```

## POST `/api/user/login`
### With Email
Request:
```json
{
    "email": "test@test.com",
    "password": "password"
}
```
Response:
```json
{
    "status": true,
    "message": "success login",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMDExZDdhOGYtMmNiNi00NTI5LTgwYTgtOTdhYTRlY2FkNTE4Iiwicm9sZSI6InVzZXIiLCJpc3MiOiJXYXN0ZTJGZWVkIiwiZXhwIjoxNzQ5NzU2OTcxLCJpYXQiOjE3NDk3Mzg5NzF9.S8ydvisfwm-I9CmHPTjhqC5-_a1vIRyh-X56knONu7g",
        "role": "user"
    }
}
```
### With Phone Number
Request:
```json
{
    "phone_number": "081215126272",
    "password": "password"
}
```
Response:
```json
{
    "status": true,
    "message": "success login",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTdjNmJkNmYtNDRlZC00ZjY2LThiYTYtNDg1YjMzZWQ2YjE2Iiwicm9sZSI6InVzZXIiLCJpc3MiOiJXYXN0ZTJGZWVkIiwiZXhwIjoxNzQ5NzU2OTMzLCJpYXQiOjE3NDk3Mzg5MzN9.-f27FGuITAtoPbNPu6VCwAudi9hZIqlA1gNYqFaXfpw",
        "role": "user"
    }
}
```

## PATCH `/api/user/update`
Request:
```json
{
    "username": "test@test.com",    # optional
    "name": "name",                 # optional
    "old_password": "password",     # optional
    "new_password": "password",     # optional
    "address": "bekasi",            # optional
    "gender": "Wanita",             # optional, Wanita or Pria
    "role": "farmer"                # optional, user or farmer
}
```
Response:
```json
{
    "status": true,
    "message": "success update user"
}
```

## GET `/api/user/me`
Response:
```json
{
    "status": true,
    "message": "success get user",
    "data": {
        "id": "97c6bd6f-44ed-4f66-8ba6-485b33ed6b16",
        "username": "test@test.com",
        "name": "password",
        "address": "bekasi",
        "email": "",
        "phone_number": "081215126272",
        "points": 0,
        "gender": "",
        "role": "farmer",
        "verified": false
    }
}
```

## POST `/api/user/redeem`
Request:
```json
{
    "bank": "BCA",
    "account": "norekening123123",
    "nominal": 10
}
```
Reponse:
```json
{
    "status": true,
    "message": "success redeem points"
}
```

