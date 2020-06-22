package main

import (
	"flag"
	"fmt"
	"thunes/objects/models"
	"thunes/settings"
	"thunes/tools"
)

func main() {
	settings.Init()
	models.Init()

	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")
	flag.Parse()

	if len(*username) == 0 {
		fmt.Println("Invalid username")
	}
	if len(*password) == 0 {
		fmt.Println("Invalid password")
	}

	createUser(*username, *password)
	fmt.Println("Success")
}

func createUser(username, password string) {
	passwordHash := tools.PasswordHash(password)
	if user, err := models.DefaultUserManager.Create(username, passwordHash); err != nil {
		panic(err)
	} else {
		fmt.Printf("User: %s\nID: %d\n", user.Username, user.UID)
	}
}
