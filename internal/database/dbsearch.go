package database

import (
	"context"

	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

// dbSearchLoginPassword - вспомогательная функция GetData для получения данных типа login password
func dbSearchLoginPassword(userId *int, ctx context.Context) ([]models.LoginPassword, error) {
	var data []models.LoginPassword
	rows, err := db.QueryContext(ctx, "select id, name, login, password, created_at from logopass where go_k_user_id = $1", userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var d models.LoginPassword
		err = rows.Scan(&d.ID, &d.Name, &d.Login, &d.Password, &d.CreatedAt)
		if err != nil {
			mlogger.Logger.Error(err.Error())
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}

// dbSearchBankCard - вспомогательная функция GetData для получения данных типа bank card
func dbSearchBankCard(userId *int, ctx context.Context) ([]models.BankCard, error) {
	var data []models.BankCard
	rows, err := db.QueryContext(ctx, "select id, card_holder_name, card_number, expiration_date, created_at from bank_card where go_k_user_id = $1", userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var d models.BankCard
		err = rows.Scan(&d.ID, &d.CardHolderName, &d.CardNumber, &d.ExpirationDate, &d.CreatedAt)
		if err != nil {
			mlogger.Logger.Error(err.Error())
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}

// dbSearchTextMessage - вспомогательная функция GetData для получения данных типа text message
func dbSearchTextMessage(userId *int, ctx context.Context) ([]models.TextMessage, error) {
	var data []models.TextMessage
	rows, err := db.QueryContext(ctx, "select id, name, text, created_at from text_data where go_k_user_id = $1", userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var d models.TextMessage
		err = rows.Scan(&d.ID, &d.Name, &d.Text, &d.CreatedAt)
		if err != nil {
			mlogger.Logger.Error(err.Error())
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}

// dbSearchBinaryMessages - вспомогательная функция GetData для получения данных типа binary message
func dbSearchBinaryMessages(userId *int, ctx context.Context) ([]models.BinaryMessage, error) {
	var data []models.BinaryMessage
	rows, err := db.QueryContext(ctx, "select id, name, file_name, location, created_at from binary_data where go_k_user_id = $1", userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var d models.BinaryMessage
		err = rows.Scan(&d.ID, &d.Name, &d.FileName, &d.Location, &d.CreatedAt)
		if err != nil {
			mlogger.Logger.Error(err.Error())
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}
