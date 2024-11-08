package clientapp

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/KznRkjp/go-keeper.git/internal/config"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

// Delete - функция удаления записи, получает на вход ID записи и тип, в зависимости от типа вызывает удаление по строке.
func Delete(dataType string, id string) error {
	mlogger.Info("delete.Delete id:" + id)
	var url string
	switch dataType {
	case "lp":
		url = config.Client.ServerAddress + config.Client.URI.DeleteLP + id
	case "bc":
		url = config.Client.ServerAddress + config.Client.URI.DeleteBC + id
	case "txt":
		url = config.Client.ServerAddress + config.Client.URI.DeleteTxt + id
	case "bm":
		url = config.Client.ServerAddress + config.Client.URI.DeleteBM + id
	default:
		mlogger.Info("Delete went wrong")

	}

	err := HTTPwithCookiesDelete(url, &User)
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}

	return nil
}

func HTTPwithCookiesDelete(url string, user *models.ClientUser) error {
	url = strings.TrimRight(url, "\n")
	mlogger.Info("url: " + url)
	req, err := http.NewRequest("DELETE", url, nil)
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
			" resp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return err
	}
	return nil
}
