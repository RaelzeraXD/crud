package main

import (
	"database/sql"
	"fmt"
	"log"

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
		log.Fatal("Error opening database:", err)
	}

	deleteuser("a", 40)
	updateage("a", 40)
	createuser("a", 10)

	//get all users
	users, err := readuser()
	if err != nil {
		log.Fatal("Error getting users:", err)
	}
	fmt.Println("--------------------------------")
	fmt.Println(users)
	fmt.Println("--------------------------------")
}

func createuser(name string, age int) error {
	_, err := db.Exec("INSERT INTO gotable (name, age) VALUES (?, ?)", name, age)
	if err != nil {
		log.Println("Error creating user:", err)
	}
	log.Printf("User created  {%s  %d}", name, age)
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

func updateage(name string, age int) error {
	_, err := db.Exec("UPDATE gotable SET age = ?  WHERE name = ? ", age, name)
	if err != nil {
		log.Println("Error updating user:", err)
	}
	log.Printf("User updated  {%s  %d}", name, age)
	return nil
}

func deleteuser(name string, age int) error {
	_, err := db.Exec("Delete from gotable WHERE age = ? AND name = ?", age, name)
	if err != nil {
		log.Println("Error deleting user:", err)
	}
	log.Printf("User deleted  {%s  %d}", name, age)
	return nil
}