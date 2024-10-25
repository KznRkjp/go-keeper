package app

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/KznRkjp/go-keeper.git/internal/database"
	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
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

func PostLoginUser(res http.ResponseWriter, req *http.Request) {
	mlogger.Info("User login")
	var user models.User
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&user)
	if err != nil {
		http.Error(res, "can't decode body", http.StatusBadRequest)
		return
	}
	mlogger.Info("Decoded user: " + user.Email + " with password " + user.Password)

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

// writes cookie directly to response
func createCookie(res *http.ResponseWriter, user *models.User) {
	mlogger.Info("Creating cookie for user" + user.Email + " with ID" + strconv.FormatInt(user.ID, 10))
	expires := time.Now().Add(encrypt.TokenExp)
	token, err := encrypt.BuildJWTString(int(user.ID))
	if err != nil {
		mlogger.Info(err.Error())
		return
	}
	mlogger.Info("JWT: " + token + " Expires:" + expires.String())
	cookie := http.Cookie{
		Name:    "JWT",
		Value:   token,
		Expires: expires,
	}
	http.SetCookie(*res, &cookie)
}
