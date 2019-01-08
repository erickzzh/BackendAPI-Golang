package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type loginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

/*
------PostSignup(w http.ResponseWriter, r *http.Request)------

Input:
{
  "email": "test@axiomzen.co",
  "password": "axiomzen",
  "firstName": "Alex",
  "lastName": "Zimmerman"
}

Output:
{
  "token": "some_jwt_token"
}

Notes:
1. need to return error if the user has already registerd
2. encode the whole payload
3. insert into the db
*/
func PostSignup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user userDetail

	//Decode the request body part and match it against the userDetail struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("bad payload \n")
		respondWithJson(w, http.StatusOK, err)

	}

	//establish connection with the local db and return a pointer to the db
	//Input:
	db := connectDB()
	newUser := &userDetail{
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	//Generate a JWT for the JSON object and inset this new user into our local db
	jwtToken := jwtSigning(newUser)
	err = newUser.insertUser(db)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, err)
		return
	}

	//Generate a response message that will return to the client server in JSON format
	//Output:
	var raw map[string]interface{}
	responseMsg := fmt.Sprintf(` {
		"token": "%s" 
	  }`, jwtToken)
	json.Unmarshal([]byte(responseMsg), &raw)
	respondWithJson(w, http.StatusOK, raw)
	closeDB()
	return
}

/*
------PostLogin(w http.ResponseWriter, r *http.Request)------

Input:
{
  "email": "test@axiomzen.co",
  "password": "axiomzen"
}

Output:
{
  "token": "some_jwt_token"
}

Notes:
1. need to capture the payload and extrat "usrname" and "password"
2. return the whole tuple after matching something in the db and pouplate a userDetail struct with it
*/
func PostLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	loginInformation := loginInfo{}

	//convet the payload into Byte for later proceeing use
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		respondWithJson(w, http.StatusBadRequest, err)
		return
	}

	db := connectDB()
	json.Unmarshal(bodyByte, &loginInformation)

	//locate the user within the db and return an error if nothing matches
	//When a match is found return the whole tuple and populate the userDetail struct
	userDetailInfo := loginInformation.findUser(db)
	if userDetailInfo == nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"result": "Incorrect username or password"})
		return
	}

	verifyResult := verifyUser(userDetailInfo.Email, userDetailInfo.Password, db)
	if verifyResult == false {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"result": "Incorrect Username or Password"})
		return
	}

	//genertae a JWT
	jwtToken := jwtSigning(userDetailInfo)

	//Generate a response message that will return to the client server in JSON format
	//Output:
	var raw map[string]interface{}
	responseMsg := fmt.Sprintf(` {
		"token": "%s"
	  }`, jwtToken)
	json.Unmarshal([]byte(responseMsg), &raw)
	respondWithJson(w, http.StatusOK, raw)
	closeDB()
	return
}

/*
------GetUsers(w http.ResponseWriter, r *http.Request)------

Input:
x-authentication-token: {value}

Output:
{
  "users": [
    {
      "email": "test@axiomzen.co",
      "firstName": "Alex",
      "lastName": "Zimmerman"
	},
	{
      "email": "test@axiomzen.co",
      "firstName": "Alex",
      "lastName": "Zimmerman"
    }
  ]
}

Notes:
1. extract the header
2. decode the JWT and find its corresponding email and password in the local db
3. validate the email and password in the db
4. return all users
*/

func GetUsers(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("x-authentication-token")

	//decode the JWT and convet it back to the claim type
	//https://jwt.io/introduction/
	claim := jwtDecoding(jwtToken)
	if claim == nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"result": "Incorrect JWT"})
		return
	}

	db := connectDB()
	//first it verifies the JWT token
	verifyResult := verifyUser(claim.Email, claim.Password, db)
	if verifyResult == false {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"result": "Incorrect JWT"})
		return
	}

	//getAllUser returns x amount of users and they are pacakaged into a struct called "users"
	allUsers, err := getAllUser(db)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, map[string]error{"result": err})
		return
	}
	finalResult := map[string]interface{}{"users": allUsers}
	respondWithJson(w, http.StatusOK, finalResult)

	closeDB()
	return

}

/*
------PutUsers(w http.ResponseWriter, r *http.Request)------

Input:
x-authentication-token: {value},
{
  "firstName": "NewFirstName",
  "lastName": "NewLastName"
}

Output:
{
  "users": [
    {
      "email": "test@axiomzen.co",
      "firstName": "Alex",
      "lastName": "Zimmerman"
    }
  ]
}

Notes:
1. extract the JWT in header
2. extractr the email and password in payload
3. validate the email and password in the db
4. update the information
5. update the db
*/
func PutUsers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newName := name{}
	jwtToken := r.Header.Get("x-authentication-token")

	//convet the payload into Byte for later proceeing use
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		respondWithJson(w, http.StatusBadRequest, err)
		return
	}

	//decode the JWT and convet it back to the claim type
	//https://jwt.io/introduction/
	claim := jwtDecoding(jwtToken)
	if claim == nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"result": "Incorrect JWT"})
		return
	}

	//check the username and password against the db
	db := connectDB()
	Err := verifyUser(claim.Email, claim.Password, db)
	if Err == nil {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"result": "Incorrect JWT"})
		return
	}

	//unmarshal to the name struct
	json.Unmarshal(bodyByte, &newName)

	//updating the local db
	newName.updateName(claim.Email, db)
	log.Printf("Name changing done")
	closeDB()
	return

}

//This is a helper function that converts responses into JSON format and send back to the client server
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
