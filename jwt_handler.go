package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

/*
need to create a new customized claim as the standar claim does not include
some of the attributes
*/
type MyCustomClaims struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.StandardClaims
}

/*
jwtSigning()
takes in a userDetail struct and return a hashed string that will mask the
information
*/
func jwtSigning(u *userDetail) string {

	mySigningKey := []byte("Erick signed it")

	// Create the Claims
	claims := MyCustomClaims{
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	//I decided to use the SHA256 hashing method as it is relatively secure
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return ""
	}
	return ss
}

/*
jwtDecoding()
takes in a hashed string and tries to un-hash the string in order to extract
information. Will return the claim struct once the data have been populated
*/
func jwtDecoding(tokenString string) *MyCustomClaims {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Erick signed it"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims
	}
	fmt.Println(err)
	return nil

}
