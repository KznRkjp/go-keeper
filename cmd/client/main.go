package main

import (
	"fmt"

	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"

	"github.com/KznRkjp/go-keeper.git/internal/buildinfo"
	clientapp "github.com/KznRkjp/go-keeper.git/internal/client-app"
	"github.com/KznRkjp/go-keeper.git/internal/config"
	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
	"go.uber.org/zap"
)

const (
	password = "x35k9f"
	msg      = `0ba7cd8c624345451df4710b81d1a349ce401e61bc7eb704ca` +
		`a84a8cde9f9959699f75d0d1075d676f1fe2eb475cf81f62ef` +
		`f701fee6a433cfd289d231440cf549e40b6c13d8843197a95f` +
		`8639911b7ed39a3aec4dfa9d286095c705e1a825b10a9104c6` +
		`be55d1079e6c6167118ac91318fe`
)

var User models.ClientUser

func main() {
	mlogger.Debug = true
	//создаем экземпляр логгера
	mlogger.Logger = zap.Must(zap.NewProduction())
	defer mlogger.Logger.Sync()

	// **** пока так
	config.Client.ServerAddress = "http://localhost:4443"
	config.Client.URI.RegisterUser = "/api/v1/register"
	config.Client.URI.LoginUser = "/api/v1/login"
	config.Client.URI.GetData = "/api/v1/data"
	//***
	buildinfo.PrintBuildVersionDate()
	var i int
	// var user models.ClientUser
	fmt.Println("To register input type 1, to login type 2")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("Enter e-mail")
		fmt.Scan(&User.User.Email)
		fmt.Println("Enter password")
		fmt.Scan(&User.User.Password)
		clientapp.RegisterUser(&User)
		fmt.Println(User)
	case 2:
		fmt.Println("Enter e-mail")
		fmt.Scan(&User.User.Email)
		fmt.Println("Enter password")
		fmt.Scan(&User.User.Password)
		clientapp.LoginUser(&User)
	default:
		panic("wrong input")
	}
	clientapp.GetData(&User)
	fmt.Println(User)
	// fmt.Println(i)
	key := sha256.Sum256([]byte(password))
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		panic(err)
	}
	// создаём вектор инициализации
	nonce := key[len(key)-aesgcm.NonceSize():]
	encrypted, err := hex.DecodeString(msg)
	if err != nil {
		panic(err)
	}
	// расшифровываем
	decrypted, err := aesgcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decrypted))
	secretString, err := encrypt.EncryptData(password, "secret message")
	if err != nil {
		panic(err)
	}
	fmt.Println(secretString)
	fmt.Println(encrypt.DecryptData(password, secretString))
}
