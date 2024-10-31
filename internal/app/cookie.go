package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

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

// вспомогательная функция проверки куки и полчения из него ID пользователя
// надо конечно во wrapper, но что получилось - то получилось
func checkCookie(req *http.Request) int {
	mlogger.Info("Checking cookie")

	cookie, err := req.Cookie("JWT")
	// fmt.Println(cookie)
	if err != nil {
		mlogger.Info(err.Error())
		return 0
	}
	userId, err := encrypt.GetUserID(cookie.Value)
	if err != nil {
		mlogger.Info(err.Error())
		return 0
	}
	return userId
}
