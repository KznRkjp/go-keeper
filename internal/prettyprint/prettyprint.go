package prettyprint

import (
	"os"

	"github.com/KznRkjp/go-keeper.git/internal/models"
	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintLP(lp []models.LoginPassword) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Login", "Password"})
	// t.AppendRows([]table.Row{
	//     {1, "Arya", "Stark", 3000},
	//     {20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	// })
	// t.AppendSeparator()
	for _, l := range lp {
		t.AppendRow([]interface{}{l.ID, l.Name, l.Login, l.Password})
	}
	// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.SetStyle(table.StyleLight)
	t.Render()
}
