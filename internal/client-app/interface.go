package clientapp

import (
	"fmt"

	"github.com/KznRkjp/go-keeper.git/internal/models"
)

var User models.ClientUser

func MainInterface() {
	var i int
	// var user models.ClientUser
	fmt.Println("To register input type 1, to login type 2")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("Enter e-mail")
		fmt.Scan(&User.User.Email)
		fmt.Println("Enter password")
		fmt.Scan(&User.User.Password)
		RegisterUser(&User)
		fmt.Println(User)
	case 2:
		fmt.Println("Enter e-mail")
		fmt.Scan(&User.User.Email)
		fmt.Println("Enter password")
		fmt.Scan(&User.User.Password)
		LoginUser(&User)
	default:
		panic("wrong input")
	}
	GetData(&User)
}
