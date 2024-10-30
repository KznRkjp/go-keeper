package clientapp

import (
	"fmt"

	"github.com/KznRkjp/go-keeper.git/internal/models"
	"github.com/KznRkjp/go-keeper.git/internal/prettyprint"
	"github.com/MasterDimmy/go-cls"
)

var User models.ClientUser

func MainInterface() {
	// cls.CLS()
	var i int
	// var user models.ClientUser
	fmt.Println("To register enter 1, to login enter 2, to exit press Ctrl+C")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("Enter e-mail")
		fmt.Scan(&User.User.Email)
		fmt.Println("Enter password")
		fmt.Scan(&User.User.Password)
		RegisterUser(&User)
		InnerInterface()
	case 2:
		fmt.Println("Enter e-mail")
		fmt.Scan(&User.User.Email)
		fmt.Println("Enter password")
		fmt.Scan(&User.User.Password)
		LoginUser(&User)
		InnerInterface()
	default:
		panic("wrong input")
	}
	// GetData(&User)
}

func InnerInterface() {
	cls.CLS()
	fmt.Println("You logged in as: " + User.User.Email)
	fmt.Println("To print login/password sheet enter 1")
	fmt.Println("To print bank cards sheet enter 2")
	fmt.Println("To print text messages sheet enter 3")
	fmt.Println("To print files sheet enter 4")
	fmt.Println("To go back to login screen enter 5")
	fmt.Println("To exit press Ctrl+C")
	var i int
	fmt.Scan(&i)
	switch i {
	case 1:
		// fmt.Println("1")
		LoginPasswordInterface("")
		// prettyprint.PrintLP(UserData.LoginPass)
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		cls.CLS()
		MainInterface()
	default:
		InnerInterface()
	}
}

func LoginPasswordInterface(message string) {
	cls.CLS()
	if message != "" {
		fmt.Println(message)
	}
	// cls.CLS()
	fmt.Println("You logged in as: " + User.User.Email)
	prettyprint.PrintLP(UserData.LoginPass)
	fmt.Println("Enter ID of record you want to edit, 0 to go back")
	var i int
	fmt.Scan(&i)
	switch i {
	case 0:
		InnerInterface()
	case 2:
		fmt.Println("2")
	default:
		LoginPasswordInterface("You've entered wrong ID")
	}
}

// func InnerInterface() {
// 	cls.CLS()
// 	fmt.Println("You logged in as: " + User.User.Email)
// 	fmt.Println("To print login/password sheet enter 1")
// 	fmt.Println("To print bank cards sheet enter 2")
// 	fmt.Println("To print text messages sheet enter 3")
// 	fmt.Println("To print files sheet enter 4")
// 	fmt.Println("To go back to login screen enter 5")
// 	fmt.Println("To exit press Ctrl+C")
// 	var i int
// 	fmt.Scan(&i)
// 	switch i {
// 	case 1:
// 		// fmt)
