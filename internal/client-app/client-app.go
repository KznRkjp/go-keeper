package clientapp

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

//curl -X POST http://localhost:4443/api/v1/register -H 'Content-Type: application/json' -d '{"email":"john@ne.doe","password":"my_password"}'

func RegisterUser(user *models.ClientUser) {
	url := "http://localhost:4443/api/v1/register"
	// k := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password)
	json := []byte(fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password))
	fmt.Println(json)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		mlogger.Info(err.Error())
	}
	// defer resp.Body.Close()
	// bytesResp, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(string(bytesResp))
	fmt.Println("resp.Cookies")
	fmt.Println(resp.Cookies())
	for _, c := range resp.Cookies() {
		fmt.Println(c.Value)
	}
	fmt.Println("resp.Request.Cookies")
	fmt.Println(resp.Request.Cookies())
	fmt.Println("Getting cookie")
	fmt.Println(resp.Request.Header)
	cookie, err := resp.Request.Cookie("JWT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Got cookie")

	fmt.Println(cookie)
	user.JWT = cookie.Value
}
