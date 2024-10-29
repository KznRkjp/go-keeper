package app

import (
	"net/http"

	"github.com/KznRkjp/go-keeper.git/internal/database"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/go-chi/chi"
)

// DeleteDataLP - удаление записи logopass po ID записи и ID пользователя
func DeleteDataLP(res http.ResponseWriter, req *http.Request) {
	recordId := chi.URLParam(req, "id")
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

func DeleteDataBC(res http.ResponseWriter, req *http.Request) {
	recordId := chi.URLParam(req, "id")
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

func DeleteDataTxt(res http.ResponseWriter, req *http.Request) {
	recordId := chi.URLParam(req, "id")
	mlogger.Info("Deleting data - Text_message with id: " + recordId)
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

func DeleteDataBM(res http.ResponseWriter, req *http.Request) {
	recordId := chi.URLParam(req, "id")
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
