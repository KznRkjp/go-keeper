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

func PutData(lp *models.LoginPassword, bc *models.BankCard, txt *models.BankCard, bm *models.BinaryMessage) error {
	var url string
	var err error
	var jsonBody []byte
	if lp != nil {
		mlogger.Info("Modifying LP record")
		url = putURL("lp")
		jsonBody, err = json.Marshal(lp)
		if err != nil {
			mlogger.Info(err.Error())
			return err
		}
	}
	if bc != nil {
		mlogger.Info("Modifying BC record")
		url = putURL("bc")
		jsonBody, err = json.Marshal(bc)
		if err != nil {
			mlogger.Info(err.Error())
			return err
		}
	}
	if txt != nil {
		mlogger.Info("Modifying TXT record")
		url = putURL("txt")
		jsonBody, err = json.Marshal(txt)
		if err != nil {
			mlogger.Info(err.Error())
			return err
		}
	}
	if bm != nil {
		mlogger.Info("Modifying BM record")
		url = putURL("bm")
		jsonBody, err = json.Marshal(bm)
		if err != nil {
			mlogger.Info(err.Error())
			return err
		}
	}
	err = HTTPwithCookiesPut(url, &User, jsonBody)
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}

	return nil
}

// func PutLP(dataType string, lp *models.LoginPassword) error {
// 	mlogger.Info("Modifying LP record")
// 	url := putURL(dataType)

// 	json, err := json.Marshal(lp)
// 	if err != nil {
// 		mlogger.Info(err.Error())
// 		return err
// 	}

// 	err = HTTPwithCookiesPut(url, &User, json)
// 	if err != nil {
// 		mlogger.Info(err.Error())
// 		return err
// 	}

// 	return nil
// }

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

func putURL(dataType string) string {
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
		url = ""

	}
	return url
}
