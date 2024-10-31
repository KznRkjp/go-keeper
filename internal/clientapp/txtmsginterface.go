package clientapp

import (
	"fmt"

	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
	"github.com/KznRkjp/go-keeper.git/internal/prettyprint"
	"github.com/MasterDimmy/go-cls"
)

func TxtMessageInterface(message string) {
	cls.CLS()
	GetData(&User)
	if message != "" {
		fmt.Println(message)
	}
	// cls.CLS()
	fmt.Println("You logged in as: " + User.User.Email)
	mlogger.Info(User.JWT)
	// mlogger.Info(string(UserData))
	prettyprint.PrintBC(UserData.BankCards, &User)
	fmt.Println("Enter ID of record you want to edit, 0 to go back and type \"add\" to add a new record")
	var i string
	fmt.Scan(&i)
	switch i {
	case "0":
		InnerInterface()
	case "add":
		AddLoginPassword()
	default:
		LoginPasswordInterface("You've entered wrong ID")
	}
}

func AddBankCard1() {
	cls.CLS()
	var lp models.LoginPassword
	fmt.Println("You logged in as: " + User.User.Email)
	fmt.Println("Type name for the record you can identify it by later on (web-site, app etc.)")
	var name, login, password string
	var err error
	fmt.Scan(&name)
	lp.Name, err = encrypt.EncryptData(User.User.Password, name)
	if err != nil {
		mlogger.Info(err.Error())
		LoginPasswordInterface(err.Error())
	}
	fmt.Println("Type login")
	fmt.Scan(&login)
	lp.Login, err = encrypt.EncryptData(User.User.Password, login)
	if err != nil {
		mlogger.Info(err.Error())
		LoginPasswordInterface(err.Error())
	}
	fmt.Println("Type password")
	fmt.Scan(&password)
	lp.Password, err = encrypt.EncryptData(User.User.Password, password)
	if err != nil {
		mlogger.Info(err.Error())
		LoginPasswordInterface(err.Error())
	}

	UserData.LoginPass = append(UserData.LoginPass, lp)
	mlogger.Info("User data local append - ok")
	err = PostDataLP(&User, &lp)
	if err != nil {
		mlogger.Info(err.Error())
		LoginPasswordInterface(err.Error())
	}

	LoginPasswordInterface("")
}
