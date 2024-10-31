package clientapp

import (
	"fmt"

	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
	"github.com/KznRkjp/go-keeper.git/internal/prettyprint"
	"github.com/MasterDimmy/go-cls"
)

func BankCardInterface(message string) {
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
		AddBankCard()
	default:
		BankCardInterface("You've entered wrong ID")
	}
}

func AddBankCard() {
	cls.CLS()
	var lp models.BankCard
	fmt.Println("You logged in as: " + User.User.Email)
	fmt.Println("Type name for the card you can identify it by later on (Visa, MIR, Hornes&Hooves etc.)")
	var err error
	cardName := cliReader()
	lp.CardName, err = encrypt.EncryptData(User.User.Password, cardName)
	if err != nil {
		mlogger.Info(err.Error())
		BankCardInterface(err.Error())
	}
	fmt.Println("Type cardholder name")
	cardholderName := cliReader()
	lp.CardHolderName, err = encrypt.EncryptData(User.User.Password, cardholderName)
	if err != nil {
		mlogger.Info(err.Error())
		BankCardInterface(err.Error())
	}
	fmt.Println("Type card number")
	cardNumber := cliReader()
	lp.CardNumber, err = encrypt.EncryptData(User.User.Password, cardNumber)
	if err != nil {
		mlogger.Info(err.Error())
		BankCardInterface(err.Error())
	}
	fmt.Println("Type expiration date")
	expirationDate := cliReader()
	lp.ExpirationDate, err = encrypt.EncryptData(User.User.Password, expirationDate)
	if err != nil {
		mlogger.Info(err.Error())
		BankCardInterface(err.Error())
	}

	UserData.BankCards = append(UserData.BankCards, lp)
	mlogger.Info("User data local append - ok")
	err = PostDataBC(&User, &lp)
	if err != nil {
		mlogger.Info(err.Error())
		BankCardInterface(err.Error())
	}

	BankCardInterface("")
}
