package main

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

//Schema for the local db, email will be the pk for each tuple
//first letter of each attribute HAS to be capitalized
type userDetail struct {
	Email     string `sql:"email,pk"`
	Password  string `sql:"password"`
	FirstName string `sql:"firstName"`
	LastName  string `sql:"lastName"`
}

//Schema that doesn't include password
type userDetailNoPassword struct {
	Email     string `sql:"email,pk" json:"email"`
	FirstName string `sql:"firstName" json:"firstName"`
	LastName  string `sql:"lastName" json:"lastName"`
}

/*
ConnectDB()
This function connects the program to a local db and create a schema based on the userDetail struct.
The schema will only be created once if the table already exists.

IMPORTANT:
1. change the user and password attribute when testing this program
2. must have a database called "users" OR change the Database attribute as well.
*/
func connectDB() (dbRef *pg.DB) {
	opts := &pg.Options{
		User:     "",     //TODO: Change
		Password: "", //TODO: Change
		Database: "users",
	}

	//establishing a connection with the local db
	var db = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to db\n")
		os.Exit(100)
	}
	log.Printf("Successuflly connected to the db \n")

	//creating a table using the userDetal schema
	err := createTable(db)
	if err != nil {
		log.Printf("Failed to create tavke %v\n", err)
		os.Exit(100)
	}
	return db
}

/*
closeDB()
This function ensures that the db connection is properly closed after the program intereacts
with the db

IMPORTANT:
1. change the user and password attribute when testing this program
2. must have a database called "users" OR change the Database attribute as well.
*/
func closeDB() {
	opts := &pg.Options{
		User:     "",     //TODO: Change
		Password: "", //TODO: Change
		Database: "users",
	}

	var db = pg.Connect(opts)
	closeERR := db.Close()
	if closeERR != nil {
		log.Printf("Error close the connection. \n")
		os.Exit(100)
	}
	log.Printf("Successfully closed the connection \n")
}

/*
createTable()
create a table using the userDetail schema
*/
func createTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	Err := db.CreateTable(&userDetail{}, opts)

	if Err != nil {
		log.Printf("Error while creating table %v", Err)
		return Err
	}
	log.Printf("Everything is fine")
	return nil
}

/*
insertUser()
insert a user into the local db when PostSignup is called
*/
func (u *userDetail) insertUser(db *pg.DB) error {
	Err := db.Insert(u)
	if Err != nil {
		log.Printf("Error inserting into the table. Error: %v", Err)
		return Err
	}

	log.Printf("User %s inserted\n", u.FirstName)
	return nil
}

/*
findUser()
return a userDetail struct when given a user email
*/
func (u *loginInfo) findUser(db *pg.DB) *userDetail {
	user := userDetail{
		Email: u.Email,
	}

	//select using the user's primary key while populating data into the
	//userDetail struct
	Err := db.Select(&user)
	if Err != nil {
		log.Printf("Error finding the user. Error: %v", Err)
		return nil
	}
	return &user
}

/*
findUser()
return a userDetail struct when given a user email
*/
func verifyUser(e string, p string, db *pg.DB) interface{} {
	user := userDetail{
		Email: e,
	}

	//select using the user's primary key while populating data into the
	//userDetail struct
	Err := db.Select(&user)
	if Err != nil {
		log.Printf("Error finding the user. Error: %v", Err)
		return false
	}

	//verify the entered password and the password stored in db
	if user.Password != p {
		log.Printf("Wrong password.")
		return false
	}

	log.Printf("Found the user")
	return true
}

/*
getAllUser()
return all users in the database but only their email,first name and last name
*/

func getAllUser(db *pg.DB) ([]userDetailNoPassword, error) {
	var users []userDetailNoPassword

	//returns all rows with only email, firstName and lastName
	_, err := db.Query(&users, `SELECT email,"firstName","lastName" FROM user_details`)
	if err != nil {
		log.Printf("Error returning all users. Error: %v", err)
		return nil, err
	}

	return users, nil
}

/*
updateName()
given the primary key in this case the email address, we will:
	1. retrive data
	2. update data
	3. update db
	4. return the updated user struct
*/
func (n *name) updateName(pk string, db *pg.DB) *userDetail {
	user := &userDetail{
		Email: pk,
	}
	//select using the user's primary key while populating data into the
	//userDetail struct
	Err := db.Select(user)
	if Err != nil {
		log.Printf("Error finding the user. Error: %v", Err)
		return nil
	}
	//altering the first and last name
	user.FirstName = n.FirstName
	user.LastName = n.LastName
	log.Printf("name has been changed")

	//updating the db
	Err = db.Update(user)
	if Err != nil {
		log.Printf("Error updating the information. Error: %v", Err)
		return nil
	}
	return user
}
