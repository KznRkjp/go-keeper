package clientapp

import (
	"fmt"
	"os"

	"github.com/KznRkjp/go-keeper.git/internal/models"
	"github.com/MasterDimmy/go-cls"
	"gitlab.com/david_mbuvi/go_asterisks"
)

var User models.ClientUser

func MainInterface() {
	cls.CLS()
	var i int
	// var err error
	fmt.Println("To register enter 1, to login enter 2, to exit press Ctrl+C")
	fmt.Print(">")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("Enter e-mail")
		fmt.Print(">")
		User.User.Email = cliReader()
		fmt.Println("Enter password")
		fmt.Print(">")
		pass, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		if err != nil {
			fmt.Println(err.Error())
		}
		User.User.Password = string(pass)
		RegisterUser(&User)
		InnerInterface()
	case 2:
		fmt.Println("Enter e-mail")
		fmt.Print(">")
		User.User.Email = cliReader()
		fmt.Println("Enter password")
		fmt.Print(">")
		pass, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		if err != nil {
			fmt.Println(err.Error())
		}
		User.User.Password = string(pass)
		LoginUser(&User) // need error check
		InnerInterface()
	default:
		panic("wrong input")
	}

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
	fmt.Print(">")

	var i int
	fmt.Scan(&i)
	switch i {
	case 1:
		// fmt.Println("1")
		LoginPasswordInterface("")
		// prettyprint.PrintLP(UserData.LoginPass)
	case 2:
		BankCardInterface("")
	case 3:
		TxtMessageInterface("")
	case 4:
		BinMessageInterface("")
	case 5:
		cls.CLS()
		MainInterface()
	default:
		InnerInterface()
	}
}
