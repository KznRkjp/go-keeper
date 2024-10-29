package clientapp

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

//curl -X POST http://localhost:4443/api/v1/register -H 'Content-Type: application/json' -d '{"email":"john@ne.doe","password":"my_password"}'

func RegisterUser(user *models.ClientUser) error {
	url := "http://localhost:4443/api/v1/register"
	// k := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password)
	json := []byte(fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password))
	fmt.Println(json)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	for _, c := range resp.Cookies() {
		if c.Name == "JWT" {
			fmt.Println(c.Value)
			user.JWT = c.Value
			mlogger.Info("Got cookie for user: " + user.JWT)
			return nil
		}
	}
	err = errors.New("no cookie")
	mlogger.Info(err.Error())
	return err
}
