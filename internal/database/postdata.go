// для облегчения чтения разбил на куски
package database

import (
	"context"

	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

// PostDataLP - добавление данных типа login password
func PostDataLP(data *models.LoginPassword, userId *int, ctx context.Context) error {
	insertDynStmt := `insert into "logopass"("name", "login", "password","go_k_user_id") values($1, $2, $3, $4)`
	_, err := db.ExecContext(ctx, insertDynStmt, data.Name, data.Login, data.Password, userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}

// PostDataLP - добавление данных типа bank card
func PostDataBC(data *models.BankCard, userId *int, ctx context.Context) error {
	insertDynStmt := `insert into "bank_card"("card_name","card_holder_name", "card_number", "expiration_date", "go_k_user_id") values($1, $2, $3, $4, $5)`
	_, err := db.ExecContext(ctx, insertDynStmt, data.CardName, data.CardHolderName, data.CardNumber, data.ExpirationDate, userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}

// PostDataLP - добавление данных типа text message
func PostDataTxt(data *models.TextMessage, userId *int, ctx context.Context) error {
	insertDynStmt := `insert into "text_data("name", "text", "go_k_user_id") values($1, $2, $3)`
	_, err := db.ExecContext(ctx, insertDynStmt, data.Name, data.Text, userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}

func PostDataBM(data *models.BinaryMessage, userId *int, ctx context.Context) error {
	insertDynStmt := `insert into "binary_data"("name", "file_name", "location","go_k_user_id") values($1, $2, $3)`
	_, err := db.ExecContext(ctx, insertDynStmt, data.Name, data.FileName, data.Location, userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}
