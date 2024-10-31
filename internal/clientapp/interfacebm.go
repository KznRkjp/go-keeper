package clientapp

import (
	"fmt"

	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
	"github.com/KznRkjp/go-keeper.git/internal/prettyprint"
	"github.com/MasterDimmy/go-cls"
)

func BinMessageInterface(message string) {
	cls.CLS()
	GetData(&User)
	if message != "" {
		fmt.Println(message)
	}
	// cls.CLS()
	fmt.Println("You logged in as: " + User.User.Email)
	mlogger.Info(User.JWT)
	// mlogger.Info(string(UserData))
	prettyprint.PrintBM(UserData.BinaryMsgs, &User)
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

func AddBMsg() {
	cls.CLS()
	var lp models.BinaryMessage
	fmt.Println("You logged in as: " + User.User.Email)
	fmt.Println("Type name for the file record you can identify it by later on (funny video, will etc.)")
	var err error
	name := cliReader()
	lp.Name, err = encrypt.EncryptData(User.User.Password, name)
	if err != nil {
		mlogger.Info(err.Error())
		TxtMessageInterface(err.Error())
	}
	fmt.Println("Type filename")
	filename := cliReader()
	lp.FileName, err = encrypt.EncryptData(User.User.Password, filename)
	if err != nil {
		mlogger.Info(err.Error())
		TxtMessageInterface(err.Error())
	}
	fmt.Println("Location? Have to think this over")
	location := cliReader()
	lp.Location, err = encrypt.EncryptData(User.User.Password, location)
	if err != nil {
		mlogger.Info(err.Error())
		TxtMessageInterface(err.Error())
	}
	UserData.BinaryMsgs = append(UserData.BinaryMsgs, lp)
	mlogger.Info("User data local append - ok")
	err = PostDataBM(&User, &lp)
	if err != nil {
		TxtMessageInterface(err.Error())
	}

	TxtMessageInterface("")

}
