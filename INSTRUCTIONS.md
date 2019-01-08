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
