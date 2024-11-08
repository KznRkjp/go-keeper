package clientapp

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KznRkjp/go-keeper.git/internal/config"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

// curl -X POST http://localhost:4443/api/v1/register -H 'Content-Type: application/json' -d '{"email":"john@ne.doe","password":"my_password"}'
// k := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password)

// Данные пользователя
var UserData models.DBSearchAll

// Регистрация пользователя
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

// Авторизация пользователя
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

// Получение данных пользователя
func GetData(user *models.ClientUser) error {
	url := config.Client.ServerAddress + config.Client.URI.GetData
	resp, err := HTTPwithCookiesGet(url, user)
	if err != nil {
		mlogger.Info(err.Error())
		time.Sleep(1 * time.Second)
		return err
	}

	err = json.Unmarshal(resp, &UserData)
	if err != nil {
		mlogger.Info(err.Error())
		return err
	}
	return nil
}

// Отправка данных пользователя c примесью cookie
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

func cliReader() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
