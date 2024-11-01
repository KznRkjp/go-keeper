package database

import (
	"context"
	"strconv"

	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
)

func DeleteDataLP(recordId *string, userId *int, ctx context.Context) error {
	mlogger.Info("recordId:" + *recordId + "userId:" + strconv.Itoa(*userId))
	insertDynStmt := `
	DELETE FROM "logopass"
	WHERE "id" = $1
	AND "go_k_user_id" = $2`
	_, err := db.ExecContext(ctx, insertDynStmt, recordId, userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}

func DeleteDataBC(recordId *string, userId *int, ctx context.Context) error {
	insertDynStmt := `
	DELETE FROM "bank_card
	WHERE "id = $1
	AND "go_k_user_id = $2`
	_, err := db.ExecContext(ctx, insertDynStmt, recordId, userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}

func DeleteDataTxt(recordId *string, userId *int, ctx context.Context) error {
	insertDynStmt := `
	DELETE FROM "text_data"
	WHERE "id = $1
	AND "go_k_user_id = $2`
	_, err := db.ExecContext(ctx, insertDynStmt, recordId, userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}

func DeleteDataBM(recordId *string, userId *int, ctx context.Context) error {
	insertDynStmt := `
	DELETE FROM "binary_data"
	WHERE "id = $1
	AND "go_k_user_id = $2`
	_, err := db.ExecContext(ctx, insertDynStmt, recordId, userId)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return err
	}
	return nil
}
