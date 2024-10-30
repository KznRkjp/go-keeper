package main

import (
	"fmt"

	"github.com/KznRkjp/go-keeper.git/internal/buildinfo"
	clientapp "github.com/KznRkjp/go-keeper.git/internal/client-app"
	"github.com/KznRkjp/go-keeper.git/internal/config"
	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"go.uber.org/zap"
)

// var User models.ClientUser

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

	clientapp.MainInterface()

	secretString, err := encrypt.EncryptData(clientapp.User.User.Password, clientapp.User.User.Password)
	if err != nil {
		panic(err)
	}
	fmt.Println(secretString)
	fmt.Println(encrypt.DecryptData(clientapp.User.User.Password, secretString))
}
