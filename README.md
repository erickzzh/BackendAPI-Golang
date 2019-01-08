# Set up 

```bash
go get -u github.com/go-pg/pg
go get -u github.com/gorilla/mux
go get -u github.com/dgrijalva/jwt-go
```

Create a database called "users"
Change the "user" and "password" attribute on line 31,32 and 64,65 in ```db.go```

This is is an examle of a request body in the JSON format
```json
{
"email": "anoffffftherone@gmail.com",
"password": "sadf",
"firstName": "erick",
"lastName": "zhang"
}
```
For the database, we are using PostgreSQL.
https://github.com/go-pg/pg

## API Specs

### `POST /signup`
Endpoint to create an user row in postgres db. The payload should have the following fields:

```json
{
  "email": "test@axiomzen.co",
  "password": "axiomzen",
  "firstName": "Alex",
  "lastName": "Zimmerman"
}
```

where `email` is an unique key in the database.

The response body should return a JWT on success that can be used for other endpoints:

```json
{
  "token": "some_jwt_token" 
}
```

### `POST /login`
Endpoint to log an user in. The payload should have the following fields:

```json
{
  "email": "test@axiomzen.co",
  "password": "axiomzen"
}
```

The response body should return a JWT on success that can be used for other endpoints:

```json
{
  "token": "some_jwt_token"
}
```

### `GET /users`
Endpoint to retrieve a json of all users. This endpoint requires a valid `x-authentication-token` header to be passed in with the request.

The response body should look like:
```json
{
  "users": [
    {
      "email": "test@axiomzen.co",
      "firstName": "Alex",
      "lastName": "Zimmerman"
    }
  ]
}
```

### `PUT /users`
Endpoint to update the current user `firstName` or `lastName` only. This endpoint requires a valid `x-authentication-token` header to be passed in and it should only update the user of the JWT being passed in. The payload can have the following fields:

```json
{
  "firstName": "NewFirstName",
  "lastName": "NewLastName"
}
```

The response can body can be empty.
