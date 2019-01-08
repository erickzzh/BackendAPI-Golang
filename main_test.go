package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_userSignUp(t *testing.T) {
	jsonData := map[string]string{"email": "erickzzh@gmail.com", "password": "yesss", "firstName": "Wrong", "lastName": "False"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("http://localhost:3000/signup", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.FailNow()

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func Test_userLogin(t *testing.T) {
	jsonData := map[string]string{"email": "erickzzh@gmail.com", "password": "yesss"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("http://localhost:3000/login", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.FailNow()

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func Test_userLogin_false(t *testing.T) {
	jsonData := map[string]string{"email": "false@gmail.com", "password": "false"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("http://localhost:3000/login", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.FailNow()

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func Test_getUser(t *testing.T) {
	client := &http.Client{}
	JWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVyaWNrenpoQGdtYWlsLmNvbSIsInBhc3N3b3JkIjoieWVzc3MiLCJmaXJzdE5hbWUiOiJXcm9uZyIsImxhc3ROYW1lIjoiRmFsc2UifQ.HDivTr3mTALn5rNLswDLFvpZOVM3e2oygZ7bTccXDM4"
	req, err := http.NewRequest("GET", "http://localhost:3000/users", nil)
	req.Header.Add("x-authentication-token", JWT)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.FailNow()

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func Test_getUser_false(t *testing.T) {
	client := &http.Client{}
	JWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVyaWNrenpoQGdtYWlsLmNvbSIsInBhc3N3b3JkIjoieWVzc3MiLCJmaXJzdE5hbWUiOiJOaWMiLCJsYXN0TmFtZSI6IlJhYm65In0.VKu94Ir-97Knxk4pgiAmbHfK6S33NWV5ReoT2QEzGI0"
	req, err := http.NewRequest("GET", "http://localhost:3000/users", nil)
	req.Header.Add("x-authentication-token", JWT)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.FailNow()

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func Test_userUpdateName(t *testing.T) {
	client := &http.Client{}
	JWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVyaWNrenpoQGdtYWlsLmNvbSIsInBhc3N3b3JkIjoieWVzc3MiLCJmaXJzdE5hbWUiOiJXcm9uZyIsImxhc3ROYW1lIjoiRmFsc2UifQ.HDivTr3mTALn5rNLswDLFvpZOVM3e2oygZ7bTccXDM4"

	jsonData := map[string]string{"firstName": "Erick", "lastName": "Zhang"}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("PUT", "http://localhost:3000/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("x-authentication-token", JWT)

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.FailNow()

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func Test_userUpdateName_false(t *testing.T) {
	client := &http.Client{}
	JWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVyaWNrenpoQGdtYWlsLmNvbSIsInBhc3N3b3JkIjoieWVzc3MiLCJmaXJzdE5hbWUiOiJOaWMiLCJsYXN0TmFtZSI6IlJhYm55In0.VKu94Ir-97Knxk4pgiAmbHfK6S33NWV5ReoT2QEzGI0"

	jsonData := map[string]string{"firstName": "Erick", "lastName": "Zhang"}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("PUT", "http://localhost:3000/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("x-authentication-token", JWT)

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.FailNow()

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func Test_theIdeaTest(t *testing.T) {
	Test_userSignUp(t)
	Test_userLogin(t)
	Test_getUser(t)
	Test_userUpdateName(t)
}
