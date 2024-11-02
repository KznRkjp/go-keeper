package clientapp

import (
	"fmt"
	"strconv"

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
	fmt.Println("Enter ID of Text record you want to edit, 0 to go back and type \"add\" to add a new record")
	var i string
	fmt.Scan(&i)
	switch i {
	case "0":
		InnerInterface()
	case "add":
		AddTxtMsg()
	default:
		for _, k := range UserData.TextMsgs {
			j, _ := strconv.ParseInt(i, 10, 64)
			if k.ID == j {
				fmt.Println(k)
				mlogger.Info("Editing record with ID: " + strconv.FormatInt(k.ID, 10))
				EditTextMsgInterface(&k)
			}
		}

		TxtMessageInterface("You've entered wrong ID")
	}
}

func EditTextMsgInterface(txtRecord *models.TextMessage) {
	var lp []models.TextMessage
	lp = append(lp, *txtRecord)
	prettyprint.PrintTxt(lp, &User)
	fmt.Println("Type \"d\" to delete, \"e\" to edit, 0 to go back")
	var i string
	fmt.Scan(&i)
	switch i {
	case "0":
		TxtMessageInterface("")
	case "d":
		DeleteTxt(txtRecord.ID)
	case "e":
		EditTxt(txtRecord)
	default:
		TxtMessageInterface("You's entered wrong command")

	}
}

func EditTxt(txt *models.TextMessage) {
	var err error
	fmt.Println("You logged in as: " + User.User.Email)
	//##
	fmt.Println("Type name for the text record you can identify it by later on (web-site, app, list of people to kill,  etc.)")
	name := cliReader()
	txt.Name, err = encrypt.EncryptData(User.User.Password, name)
	if err != nil {
		mlogger.Info(err.Error())
		TxtMessageInterface(err.Error())
	}
	fmt.Println("Type text you wish to save")
	text := cliReader()
	txt.Text, err = encrypt.EncryptData(User.User.Password, text)
	if err != nil {
		mlogger.Info(err.Error())
		TxtMessageInterface(err.Error())
	}
	PutData(nil, nil, txt, nil)
	TxtMessageInterface("Record edited")
}

func DeleteTxt(id int64) {
	for i, k := range UserData.TextMsgs {
		if k.ID == id {
			UserData.TextMsgs = append(UserData.TextMsgs[:i], UserData.TextMsgs[i+1:]...)
		}
	}
	stringId := strconv.FormatInt(id, 10)
	Delete("txt", stringId)
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
