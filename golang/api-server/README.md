# api-server

GET /users
request
```
curl localhost:8080/users
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

