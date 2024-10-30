package clientapp

import (
	"encoding/json"

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
