package clientapp

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/KznRkjp/go-keeper.git/internal/config"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

//curl -X POST http://localhost:4443/api/v1/register -H 'Content-Type: application/json' -d '{"email":"john@ne.doe","password":"my_password"}'

func RegisterUser(user *models.ClientUser) error {
	url := config.Client.ServerAddress + config.Client.URI.RegisterUser
	// k := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password)
	json := []byte(fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password))
	// fmt.Println(json)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	for _, c := range resp.Cookies() {
		if c.Name == "JWT" {
			// fmt.Println(c.Value)
			user.JWT = c.Value
			mlogger.Info("Got cookie for user: " + user.JWT)
			return nil
		}
	}
	err = errors.New("no cookie")
	mlogger.Info(err.Error())
	return err
}

func LoginUser(user *models.ClientUser) error {
	url := config.Client.ServerAddress + config.Client.URI.LoginUser
	json := []byte(fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password))
	// fmt.Println(json)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	for _, c := range resp.Cookies() {
		if c.Name == "JWT" {
			fmt.Println(c.Value)
			user.JWT = c.Value
			fmt.Println(user.JWT)
			mlogger.Info("Got cookie for user: " + string(user.JWT))
			return nil
		}
	}
	err = errors.New("no cookie")
	mlogger.Info(err.Error())
	return err
}

func GetData(user *models.ClientUser) error {
	url := config.Client.ServerAddress + config.Client.URI.GetData
	cookie := &http.Cookie{
		Name:  "JWT",
		Value: user.JWT,
	}
	resp, err := http.Get(url)
	if err != nil {
		mlogger.Info(err.Error())
	}
	fmt.Println(resp)
	return nil
}
