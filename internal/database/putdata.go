package database

import (
	"context"
	"strconv"

	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

// PutDataLP - обновляет запись logopass в базе данных.
func PutDataLP(data *models.LoginPassword, userId *int, ctx context.Context) error {
	mlogger.Info("Updating logopass record with ID: " + strconv.FormatInt(data.ID, 10) + " for user ID: " + strconv.Itoa(*userId))
	updateDynStmt := `UPDATE logopass 
	SET 
	("name", 
	"login", 
	"password") 
	= ($1, $2, $3) 
	WHERE id = $4
	and go_k_user_id = $5`
	_, err := db.ExecContext(ctx, updateDynStmt, data.Name, data.Login, data.Password, data.ID, *userId)
	// return err
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil

}

// PutDataBC - обновляет запись bankcard в базе данных.
func PutDataBC(data *models.BankCard, userId *int, ctx context.Context) error {
	mlogger.Info("Updating bankcard record with ID: " + strconv.FormatInt(data.ID, 10) + " for user ID: " + strconv.Itoa(*userId))
	updateDynStmt := `UPDATE bank_card 
	SET 
	("card_holder_name", 
	"card_number", 
	"expiration_date") 
	= ($1, $2, $3) 
	WHERE id = $4
	and go_k_user_id = $5`
	_, err := db.ExecContext(ctx, updateDynStmt, data.CardHolderName, data.CardNumber, data.ExpirationDate, data.ID, *userId)
	// return err
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}

// PutDataTxt - обновляет запись text_data в базе данных.
func PutDataTxt(data *models.TextMessage, userId *int, ctx context.Context) error {
	mlogger.Info("Updating TextMessage record with ID: " + strconv.FormatInt(data.ID, 10) + " for user ID: " + strconv.Itoa(*userId))
	updateDynStmt := `UPDATE text_data 
	SET 
	("name", 
	"text") 
	= ($1, $2)
	WHERE id = $3
	and go_k_user_id = $4`
	_, err := db.ExecContext(ctx, updateDynStmt, data.Name, data.Text, data.ID, *userId)
	// return err
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}

func PutDataBM(data *models.BinaryMessage, userId *int, ctx context.Context) error {
	mlogger.Info("Updating BinaryMessage record with ID: " + strconv.FormatInt(data.ID, 10) + " for user ID: " + strconv.Itoa(*userId))
	updateDynStmt := `UPDATE binary_data
	SET 
	("name", 
	"file_name",
	"location") 
	= ($1, $2, $3)
	WHERE id = $4
	and go_k_user_id = $5`
	_, err := db.ExecContext(ctx, updateDynStmt, data.Name, data.FileName, data.Location, data.ID, *userId)
	// return err
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}
