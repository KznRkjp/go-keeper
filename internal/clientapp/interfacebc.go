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

// BankCardInterface - реализует интерфейс для работы с банковскими картами, основное меню
func BankCardInterface(message string) {
	cls.CLS()
	GetData(&User)
	if message != "" {
		fmt.Println(message)
	}
	fmt.Println("You logged in as: " + User.User.Email)
	mlogger.Info(User.JWT)
	// вывод всех сохраненных карт
	prettyprint.PrintBC(UserData.BankCards, &User)

	fmt.Println("Enter ID of bank card you want to edit/delete, 0 to go back and type \"add\" to add a new record")
	var i string
	fmt.Scan(&i)
	switch i {
	case "0":
		InnerInterface()
	case "add":
		AddBankCard()
	default:
		for _, k := range UserData.BankCards {
			j, _ := strconv.ParseInt(i, 10, 64)
			if k.ID == j {
				fmt.Println(k)
				mlogger.Info("Editing record with ID: " + strconv.FormatInt(k.ID, 10))
				EditBankCardInterface(&k)
			}
		}

		BankCardInterface("You've entered wrong ID")
	}
}

// EditBankCardInterface - реализует интерфейс для редактирования или удаления банковской карты
func EditBankCardInterface(bcRecord *models.BankCard) {
	var lp []models.BankCard
	lp = append(lp, *bcRecord)
	prettyprint.PrintBC(lp, &User)
	fmt.Println("Type \"d\" to delete, \"e\" to edit, 0 to go back")
	var i string
	fmt.Scan(&i)
	switch i {
	case "0":
		BankCardInterface("")
	case "d":
		DeleteBC(bcRecord.ID)
	case "e":
		EditBC(bcRecord)
	default:
		BankCardInterface("You've entered wrong command")

	}
}

// EditBC - реализует функцию редактирования банковской карты
func EditBC(bc *models.BankCard) {
	var err error
	fmt.Println("You logged in as: " + User.User.Email)
	//##
	fmt.Println("Type name for the card you can identify it by later on (Visa, MIR, Hornes&Hooves etc.)")
	cardName := cliReader()
	bc.CardName, err = encrypt.EncryptData(User.User.Password, cardName)
	if err != nil {
		mlogger.Info(err.Error())
		BankCardInterface(err.Error())
	}
	//##
	fmt.Println("Type cardholder name")
	cardholderName := cliReader()
	bc.CardHolderName, err = encrypt.EncryptData(User.User.Password, cardholderName)
	if err != nil {
		mlogger.Info(err.Error())
		BankCardInterface(err.Error())
	}
	fmt.Println("Type card number")
	cardNumber := cliReader()
	bc.CardNumber, err = encrypt.EncryptData(User.User.Password, cardNumber)
	if err != nil {
		mlogger.Info(err.Error())
		BankCardInterface(err.Error())
	}
	fmt.Println("Type expiration date")
	expirationDate := cliReader()
	bc.ExpirationDate, err = encrypt.EncryptData(User.User.Password, expirationDate)
	if err != nil {
		mlogger.Info(err.Error())
		BankCardInterface(err.Error())
	}
	PutData(nil, bc, nil, nil)
	BankCardInterface("Record edited")
}

// DeleteBC - реализует функцию удаления банковской карты
func DeleteBC(id int64) {
	for i, k := range UserData.BankCards {
		if k.ID == id {
			UserData.BankCards = append(UserData.BankCards[:i], UserData.BankCards[i+1:]...)
		}
	}
	stringId := strconv.FormatInt(id, 10)
	Delete("bc", stringId)
}

// AddBankCard - реализует функцию добавления банковской карты
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
