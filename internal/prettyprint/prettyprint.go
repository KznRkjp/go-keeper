package prettyprint

import (
	"os"

	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/models"
	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintLP(lp []models.LoginPassword, user *models.ClientUser) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Login", "Password"})
	// t.AppendRows([]table.Row{
	//     {1, "Arya", "Stark", 3000},
	//     {20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	// })
	// t.AppendSeparator()
	for _, l := range lp {
		name, err := encrypt.DecryptData(user.User.Password, l.Name)
		if err != nil {
			name = string(l.Name)
		}
		// l.Name = name
		login, err := encrypt.DecryptData(user.User.Password, l.Login)
		if err != nil {
			login = string(l.Login)
		}
		password, err := encrypt.DecryptData(user.User.Password, l.Password)
		if err != nil {
			password = string(l.Password)
		}
		// l.Password = password
		// t.AppendRow([])
		t.AppendRow([]interface{}{l.ID, name, login, password})
	}
	// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.SetStyle(table.StyleLight)
	t.Render()
}

func PrintBC(lp []models.BankCard, user *models.ClientUser) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Card name", "Cardholder name", "Card number", "Expiration date"})
	for _, l := range lp {
		cardName, err := encrypt.DecryptData(user.User.Password, l.CardName)
		if err != nil {
			cardName = string(l.CardName)
		}

		cardholderName, err := encrypt.DecryptData(user.User.Password, l.CardHolderName)
		if err != nil {
			cardholderName = string(l.CardHolderName)
		}
		cardNumber, err := encrypt.DecryptData(user.User.Password, l.CardNumber)
		if err != nil {
			cardNumber = string(l.CardNumber)
		}
		expirationDate, err := encrypt.DecryptData(user.User.Password, l.ExpirationDate)
		if err != nil {
			expirationDate = string(l.ExpirationDate)
		}
		t.AppendRow([]interface{}{l.ID, cardName, cardholderName, cardNumber, expirationDate})
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}

func PrintTxt(lp []models.Text, user *models.ClientUser) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Text"})
	for _, l := range lp {
		name, err := encrypt.DecryptData(user.User.Password, l.Name)
		if err != nil {
			name = string(l.Name)
		}
		t.AppendRow([]interface{}{l.ID, name, string(l.Text)})
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}
