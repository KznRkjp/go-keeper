package clientapp

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
)

//curl -X POST http://localhost:4443/api/v1/register -H 'Content-Type: application/json' -d '{"email":"john@ne.doe","password":"my_password"}'

func RegisterUser(user *models.ClientUser) {
	url := "http://localhost:4443/api/v1/register"
	// k := fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password)
	// fmt.Println(k)
	json := []byte(fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.User.Email, user.User.Password))
	fmt.Println(json)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		mlogger.Info(err.Error())
	}
	// req.Header.Set("Content-Type", "application/json")
	bytesResp, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		log.Println(err1)
		return
	}

	// Выводим содержимое тела ответа
	fmt.Println(string(bytesResp))
	// bytesRespheader, err2 := io.ReadAll(resp.Header.Get("Cookie"))
	// if err2 != nil {
	// 	log.Println(err2)
	// 	return
	// }
	fmt.Println(resp.Request.Cookie("JWT"))
}
