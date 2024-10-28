package app

import (
	"encoding/json"
	"net/http"

	"github.com/KznRkjp/go-keeper.git/internal/database"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

// PostData - handler для /api/v1/data/lp - добаляает данные login password в базу.
func PostDataLP(res http.ResponseWriter, req *http.Request) {
	mlogger.Info("Posting data - LP")
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
	// data.UserID = userId
	err = database.PostDataLP(&data, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

// PostData - handler для /api/v1/data/bc - добаляает данные bank card в базу.
func PostDataBC(res http.ResponseWriter, req *http.Request) {
	mlogger.Info("Posting data - Bank Card")
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
	err = database.PostDataBC(&data, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

// PostData - handler для /api/v1/data/txt - добаляает данные text в базу.
func PostDataTxt(res http.ResponseWriter, req *http.Request) {
	mlogger.Info("Posting data - Text")
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
	err = database.PostDataTxt(&data, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

func PostDataBM(res http.ResponseWriter, req *http.Request) {
	mlogger.Info("Posting data - Binary Message")
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
	err = database.PostDataBM(&data, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)

}
