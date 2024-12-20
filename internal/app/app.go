package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KznRkjp/go-keeper.git/internal/database"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

// PostRegisterUser gets user credentials, saves to database and returns JWT token.
// curl -X POST http://localhost:4443/api/v1/register
// -H 'Content-Type: application/json'
// -d '{"email":"john@ne.doe","password":"my_password"}'
func PostRegisterUser(res http.ResponseWriter, req *http.Request) {
	var newUser models.User
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&newUser)
	if err != nil {
		http.Error(res, "can't decode body", http.StatusBadRequest)
		return
	}
	user, err := database.RegisterUser(&newUser, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user == nil {
		mlogger.Info("User already exists")
		res.WriteHeader(http.StatusConflict)
		return
	}
	createCookie(&res, user)
	res.WriteHeader(http.StatusCreated)
	mlogger.Info("Added user: " + user.Email + " with ID " + strconv.FormatInt(user.ID, 10))
}

// PostLoginUser gets user credentials, checks them against database and returns JWT token.
// curl -X POST http://localhost:4443/api/v1/login
// -H 'Content-Type: application/json'
// -d '{"email":*************,"password":"my_password}'
func PostLoginUser(res http.ResponseWriter, req *http.Request) {
	mlogger.Info("User login")
	var user models.User
	fmt.Println(user)
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&user)
	if err != nil {
		http.Error(res, "can't decode body", http.StatusBadRequest)
		return
	}
	mlogger.Info("Decoded user (from body): " + user.Email + " with password " + user.Password)

	newUser, err := database.LoginUser(&user, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if newUser == nil {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}
	createCookie(&res, newUser)
	res.WriteHeader(http.StatusOK)
}

// GetData - handler для /api/v1/data - получает ВСЕ данные из базы.
// основной хэндл для получания данных клиентом
func GetData(res http.ResponseWriter, req *http.Request) {
	mlogger.Info("Getting data")
	userId := checkCookie(req)
	if userId == 0 {
		mlogger.Info("User not found")
		res.WriteHeader(http.StatusNotFound)
		return
	}
	data, err := database.GetData(&userId, req.Context())
	if err != nil {
		mlogger.Info(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	// if data == nil {
	// 	mlogger.Info("Data not found")
	// 	res.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}
