package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// creating the same struct in the database
type User struct {
	Name string
	Age  int
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost)/godb")
	if err != nil {
		panic(err)
	}

	user1 := User{"maria", 20}
	user2 := User{"maria",90}

	deleteuser(user2)
	updateuser(user1)
	createuser(user1)

	//get all users
	users, err := readuser()
	if err != nil {panic(err)}
	fmt.Println("--------------------------------")
	fmt.Println(users)
	fmt.Println("--------------------------------")
}

func createuser(user User) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO gotable VALUES ('%s', '%d')", user.Name, user.Age))
	if err != nil {
		return err
	}
	fmt.Println("created",user)
	return nil
}

func readuser() ([]User, error) {
	res, err := db.Query("SELECT * FROM gotable")
	if err != nil {
		return nil, err
	}

	users := []User{}

	for res.Next() {
		var user User
		if err := res.Scan(&user.Name, &user.Age); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func updateuser(user User) error {
	_, err := db.Exec(fmt.Sprintf("UPDATE gotable SET age = 90  WHERE name = '%s' ", user.Name))
	if err != nil {
		return err
	}
	fmt.Println("updated",user)
	return nil
}

func deleteuser(user User) error {
	_ , err := db.Exec(fmt.Sprintf("DELETE FROM gotable WHERE age = '%d' and name = '%s' ", user.Age, user.Name))
	if err != nil {
		return err
	}
	fmt.Println("deleted", user)
	return nil
}
