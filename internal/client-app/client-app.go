package clientapp

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/KznRkjp/go-keeper.git/internal/config"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
	"github.com/KznRkjp/go-keeper.git/internal/prettyprint"
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
	resp, err := HTTPwithCookiesGet(url, user)
	if err != nil {
		return err
	}
	var data models.DBSearchAll
	err = json.Unmarshal(resp, &data)
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	// defer resp.Close()
	fmt.Println("###############")
	fmt.Println(data)
	prettyprint.PrintLP(data.LoginPass)
	fmt.Println("###############")
	return nil
}

func HTTPwithCookiesGet(url string, user *models.ClientUser) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "JWT", Value: user.JWT})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var body []byte

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	body, err = io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return nil, err
	}

	return body, err
}