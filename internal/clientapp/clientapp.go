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
)

// curl -X POST http://localhost:4443/api/v1/register -H 'Content-Type: application/json' -d '{"email":"john@ne.doe","password":"my_password"}'
// k := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password)
var UserData models.DBSearchAll

func RegisterUser(user *models.ClientUser) error {
	url := config.Client.ServerAddress + config.Client.URI.RegisterUser

	json := []byte(fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	for _, c := range resp.Cookies() {
		if c.Name == "JWT" {
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
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	for _, c := range resp.Cookies() {
		if c.Name == "JWT" {
			user.JWT = c.Value
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

	err = json.Unmarshal(resp, &UserData)
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	// defer resp.Close()
	fmt.Println("###############")
	// prettyprint.PrintLP(UserData.LoginPass)
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

func HTTPwithCookiesPost(url string, user *models.ClientUser, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "JWT", Value: user.JWT})
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// var body []byte

	// reader, err := gzip.NewReader(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer reader.Close()
	// body, err = io.ReadAll(reader)
	// if err != nil {
	// 	panic(err)
	// }

	if resp.StatusCode != 201 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return nil, err
	}

	return nil, nil
}

func PostDataLP(user *models.ClientUser, data *models.LoginPassword) error {
	url := config.Client.ServerAddress + config.Client.URI.PostLP
	json, err := json.Marshal(data)
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	resp, err := HTTPwithCookiesPost(url, user, json)
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	fmt.Println(string(resp))
	return nil
}

// err = json.Unmarshal(resp, &UserData)
// if err != nil {
// 	mlogger.Info(err.Error())
// 	return err
// }
// // defer resp.Close()
// fmt.Println("###############")
// prettyprint)
