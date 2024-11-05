package app

import (
	"encoding/json"
	"net/http"

	"github.com/KznRkjp/go-keeper.git/internal/database"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

// PutDataLP - хэндлер для /api/v1/data/bc метод PUT - редактирования записи login_password
func PutDataLP(res http.ResponseWriter, req *http.Request) {
	userId := checkCookie(req)
	if userId == 0 {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}
	var data models.LoginPassword
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&data)
	if err != nil {
		http.Error(res, "can't decode body", http.StatusBadRequest)
		return
	}
	err = database.PutDataLP(&data, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)

}

// PutDataBC - хэндлер для /api/v1/data/bc метод PUT - редактирования записи bank_card
func PutDataBC(res http.ResponseWriter, req *http.Request) {
	userId := checkCookie(req)
	if userId == 0 {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}
	var data models.BankCard
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&data)
	if err != nil {
		http.Error(res, "can't decode body", http.StatusBadRequest)
		return
	}
	err = database.PutDataBC(&data, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)

}

// PutDataTxt - хендлер для /api/v1/data/txt метод PUT - редактирования записи text_message
func PutDataTxt(res http.ResponseWriter, req *http.Request) {
	userId := checkCookie(req)
	if userId == 0 {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}
	var data models.TextMessage
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&data)
	if err != nil {
		http.Error(res, "can't decode body", http.StatusBadRequest)
		return
	}
	err = database.PutDataTxt(&data, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)

}

// PutDataBM - хендлер для /api/v1/data/bm метод PUT - редактирования записи binary_message
func PutDataBM(res http.ResponseWriter, req *http.Request) {
	userId := checkCookie(req)
	if userId == 0 {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}
	var data models.BinaryMessage
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&data)
	if err != nil {
		http.Error(res, "can't decode body", http.StatusBadRequest)
		return
	}
	err = database.PutDataBM(&data, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)

}
