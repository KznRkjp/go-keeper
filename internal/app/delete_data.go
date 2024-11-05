package app

import (
	"net/http"
	"strings"

	"github.com/KznRkjp/go-keeper.git/internal/database"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
)

// DeleteDataLP - хэндлер для /api/v1/data/lp/{id} - удаление записи logopass пo ID записи и ID пользователя
func DeleteDataLP(res http.ResponseWriter, req *http.Request) {
	recordId := strings.Trim(req.RequestURI, "/")
	recordId = strings.Split(recordId, "/")[len(strings.Split(recordId, "/"))-1]
	mlogger.Info("Deleting data - LP with id: " + recordId)
	userId := checkCookie(req)
	if userId == 0 {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}

	err := database.DeleteDataLP(&recordId, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

// DeleteDataBC - хэндлер для /api/v1/data/bc/{id} - удаление записи bankcard пo ID записи и ID пользователя
func DeleteDataBC(res http.ResponseWriter, req *http.Request) {
	recordId := strings.Trim(req.RequestURI, "/")
	recordId = strings.Split(recordId, "/")[len(strings.Split(recordId, "/"))-1]
	mlogger.Info("Deleting data - BC with id: " + recordId)
	userId := checkCookie(req)
	if userId == 0 {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}

	err := database.DeleteDataBC(&recordId, &userId, req.Context())
	if err != nil {
		mlogger.Info("failed to delete the payment card data: " + err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

// DeleteDataTxt - хэндлер для /api/v1/data/txt/{id} - удаление записи text пo ID записи и ID пользователя
func DeleteDataTxt(res http.ResponseWriter, req *http.Request) {
	recordId := strings.Trim(req.RequestURI, "/")
	recordId = strings.Split(recordId, "/")[len(strings.Split(recordId, "/"))-1]
	mlogger.Info("Deleting data - TXT with id: " + recordId)
	userId := checkCookie(req)
	if userId == 0 {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}

	err := database.DeleteDataTxt(&recordId, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

// DeleteDataBM - хэндлер для /api/v1/data/bm/{id} - удаление записи binary_message пo ID записи и ID пользователя
func DeleteDataBM(res http.ResponseWriter, req *http.Request) {
	recordId := strings.Trim(req.RequestURI, "/")
	recordId = strings.Split(recordId, "/")[len(strings.Split(recordId, "/"))-1]
	mlogger.Info("Deleting data - Binary_Message with id: " + recordId)
	userId := checkCookie(req)
	if userId == 0 {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}

	err := database.DeleteDataBM(&recordId, &userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
