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
		AddBm()
	default:
		for _, k := range UserData.BinaryMsgs {
			j, _ := strconv.ParseInt(i, 10, 64)
			if k.ID == j {
				fmt.Println(k)
				mlogger.Info("Editing record with ID: " + strconv.FormatInt(k.ID, 10))
				EditBmInterface(&k)
			}
		}

		BinMessageInterface("You've entered wrong ID")
	}
}

func EditBmInterface(bmRecord *models.BinaryMessage) {
	var lp []models.BinaryMessage
	lp = append(lp, *bmRecord)
	prettyprint.PrintBM(lp, &User)
	fmt.Println("Type \"d\" to delete, \"e\" to edit, 0 to go back")
	var i string
	fmt.Scan(&i)
	switch i {
	case "0":
		BinMessageInterface("")
	case "d":
		DeleteBm(bmRecord.ID)
	case "e":
		EditBm(bmRecord)
	default:
		BinMessageInterface("You've entered wrong command")

	}
}

func EditBm(bm *models.BinaryMessage) {
	var err error
	fmt.Println("You logged in as: " + User.User.Email)
	//##
	fmt.Println("Type name for the file record you can identify it by later on (funny video, will etc.)")
	name := cliReader()
	bm.Name, err = encrypt.EncryptData(User.User.Password, name)
	if err != nil {
		mlogger.Info(err.Error())
		BinMessageInterface(err.Error())
	}
	fmt.Println("Type filename")
	filename := cliReader()
	bm.FileName, err = encrypt.EncryptData(User.User.Password, filename)
	if err != nil {
		mlogger.Info(err.Error())
		BinMessageInterface(err.Error())
	}
	fmt.Println("Location? Have to think this over")
	location := cliReader()
	bm.Location, err = encrypt.EncryptData(User.User.Password, location)
	if err != nil {
		mlogger.Info(err.Error())
		BinMessageInterface(err.Error())
	}

	PutData(nil, nil, nil, bm)
	BinMessageInterface("Record edited")
}

func DeleteBm(id int64) {
	for i, k := range UserData.BinaryMsgs {
		if k.ID == id {
			UserData.BinaryMsgs = append(UserData.BinaryMsgs[:i], UserData.BinaryMsgs[i+1:]...)
		}
	}
	stringId := strconv.FormatInt(id, 10)
	Delete("bm", stringId)
}

func AddBm() {
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
		BinMessageInterface(err.Error())
	}
	fmt.Println("Location? Have to think this over")
	location := cliReader()
	lp.Location, err = encrypt.EncryptData(User.User.Password, location)
	if err != nil {
		mlogger.Info(err.Error())
		BinMessageInterface(err.Error())
	}
	UserData.BinaryMsgs = append(UserData.BinaryMsgs, lp)
	mlogger.Info("User data local append - ok")
	err = PostDataBM(&User, &lp)
	if err != nil {
		BinMessageInterface(err.Error())
	}

	BinMessageInterface("")

}
