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

func PutLP(dataType string, lp *models.LoginPassword) error {
	mlogger.Info("Modifying LP record")
	var url string
	switch dataType {
	case "lp":
		url = config.Client.ServerAddress + config.Client.URI.PutLP
	case "bc":
		url = config.Client.ServerAddress + config.Client.URI.PutBC
	case "txt":
		url = config.Client.ServerAddress + config.Client.URI.PutTxt
	case "bm":
		url = config.Client.ServerAddress + config.Client.URI.PutBM
	default:
		mlogger.Info("Something went wrong")

	}

	json, err := json.Marshal(lp)
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}

	err = HTTPwithCookiesPut(url, &User, json)
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}

	return nil
}

func HTTPwithCookiesPut(url string, user *models.ClientUser, data []byte) error {

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "JWT", Value: user.JWT})
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 {
		err = errors.New(url +
			"resp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return err
	}
	return nil
}
