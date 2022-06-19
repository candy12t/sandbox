# api-server

## GET /users

request

```shell
curl localhost:8080/users
```

response

```json
[
  {
    "id": 1,
    "name": "Tom",
    "created_at": "2022-06-15 13:00:00",
    "updated_at": "2022-06-15 13:00:00"
  },
  {
    "id": 2,
    "name": "Mike",
    "created_at": "2022-06-15 13:00:00",
    "updated_at": "2022-06-15 13:00:00"
  }
]
```

## GET /users/:id

request

```shell
curl localhost:8080/users/1
```

response

```json
{
  "id": 1,
  "name": "Tom",
  "created_at": "2022-06-15 13:00:00",
  "updated_at": "2022-06-15 13:00:00"
}
```

## POST /users/

request

```shell
curl -X POST localhost:8080/users \
  -d {"name": "Mike"}
```

response

```json
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

## PUT /users/:id

request

```shell
curl -X PUT localhost:8080/users/1 \
  -d {"name": "Mike"}
```

response

```json
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

## DELETE /users/:id

request

```shell
curl -X DELETE localhost:8080/users/1
```

response

```json
{
  "message": "delete user is success!!"
}
```

## GET /tags

## GET /tags/:id

## POST /tags

## PUT /tags/:id

## DELETE /tags/:id

## GET /users/:id/tags
