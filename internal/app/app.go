package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/KznRkjp/go-keeper.git/internal/database"
	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

func PostRegisterUser(res http.ResponseWriter, req *http.Request) {
	newUser := models.User{
		Email:    "john@doe.com",
		Password: "verystrongpassword",
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
