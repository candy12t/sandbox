# api-server

GET /users
request
```
curl localhost:8080/users
```
response
```
[
  {
    "id": 1,
    "name": "Tom",
    "created_at": "2022-06-15 13:00:00",
    "updated_at": "2022-06-15 13:00:00",
  },
  {
    "id": 2,
    "name": "Mike",
    "created_at": "2022-06-15 13:00:00",
    "updated_at": "2022-06-15 13:00:00",
  }
]
```

GET /users/:id
request
```
curl localhost:8080/users/1
```
response
```
{
  "id": 1,
  "name": "Tom",
  "created_at": "2022-06-15 13:00:00",
  "updated_at": "2022-06-15 13:00:00",
}
```

POST /users/
request
```
curl -X POST localhost:8080/users \
  -d {"name": "Mike"}
```
response
```
{
  "message": "create user is success!!",
  "result": {
    "id": 1,
    "name: "Mike",
    "created_at": "2022-06-16 00:00:00",
    "updated_at": "2022-06-16 00:00:00"
  }
}
```

PUT /users/:id
request
```
curl -X PUT localhost:8080/users/1 \
  -d {"name": "Mike"}
```
response
```
{
  "message": "update user is success!!",
  "result": {
    "id": 1,
    "name: "Mike",
    "created_at": "2022-06-16 00:00:00",
    "updated_at": "2022-06-17 00:00:00"
  }
}
```

DELETE /users/:id
request
```
curl -X DELETE localhost:8080/users/1
```
response
```
{
  "message": "delete user is success!!"
}
```
