package clientapp

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/KznRkjp/go-keeper.git/internal/config"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

func PostDataLP(user *models.ClientUser, data *models.LoginPassword) error {
	url := config.Client.ServerAddress + config.Client.URI.PostLP
	var err error
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
	mlogger.Info("resp: " + string(resp))
	// fmt.Println(string(resp))
	return nil
}

func PostDataBC(user *models.ClientUser, data *models.BankCard) error {
	url := config.Client.ServerAddress + config.Client.URI.PostBC
	var err error
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
	mlogger.Info("resp: " + string(resp))
	return nil
}

func PostDataTxt(user *models.ClientUser, data *models.TextMessage) error {
	url := config.Client.ServerAddress + config.Client.URI.PostTxt
	var err error
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
	mlogger.Info("resp: " + string(resp))
	return nil
}

func PostDataBM(user *models.ClientUser, data *models.BinaryMessage) error {
	url := config.Client.ServerAddress + config.Client.URI.PostBM
	var err error
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
	mlogger.Info("resp: " + string(resp))
	return nil
}

// fmt.Println(
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
	if resp.StatusCode != 201 {
		err = errors.New(url +
			"resp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return nil, err
	}
	return nil, nil
}
