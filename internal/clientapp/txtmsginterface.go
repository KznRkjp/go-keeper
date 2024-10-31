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
	prettyprint.PrintTxt(UserData.TextMsgs, &User)
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

func AddTxtMsg() {
	cls.CLS()
	var lp models.TextMessage
	fmt.Println("You logged in as: " + User.User.Email)
	fmt.Println("Type name for the text record you can identify it by later on (web-site, app, list of people to kill,  etc.)")
	var err error
	name := cliReader()
	lp.Name, err = encrypt.EncryptData(User.User.Password, name)
	if err != nil {
		mlogger.Info(err.Error())
		TxtMessageInterface(err.Error())
	}
	fmt.Println("Type text you wish to save")
	text := cliReader()
	lp.Text, err = encrypt.EncryptData(User.User.Password, text)
	if err != nil {
		mlogger.Info(err.Error())
		TxtMessageInterface(err.Error())
	}
	UserData.TextMsgs = append(UserData.TextMsgs, lp)
	mlogger.Info("User data local append - ok")
	err = PostDataTxt(&User, &lp)
	if err != nil {
		TxtMessageInterface(err.Error())
	}

	TxtMessageInterface("")

}
