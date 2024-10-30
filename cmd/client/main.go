package main

import (
	"fmt"

	"github.com/KznRkjp/go-keeper.git/internal/buildinfo"
	"github.com/KznRkjp/go-keeper.git/internal/clientapp"
	"github.com/KznRkjp/go-keeper.git/internal/config"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/MasterDimmy/go-cls"
	"go.uber.org/zap"
)

// var User models.ClientUser

func main() {
	cls.CLS()
	mlogger.Debug = true
	//создаем экземпляр логгера
	mlogger.Logger = zap.Must(zap.NewProduction())
	defer mlogger.Logger.Sync()

	// **** пока так
	config.Client.ServerAddress = "http://localhost:4443"
	config.Client.URI.RegisterUser = "/api/v1/register"
	config.Client.URI.LoginUser = "/api/v1/login"
	config.Client.URI.GetData = "/api/v1/data"
	config.Client.URI.PostLP = "/api/v1/data/lp"
	config.Client.URI.PostBC = "/api/v1/data/bc"
	config.Client.URI.PostTxt = "/api/v1/data/txt"
	config.Client.URI.PostBM = "/api/v1/data/bm"
	config.Client.URI.DeleteLP = "/api/v1/data/lp/"   //{id}
	config.Client.URI.DeleteBC = "/api/v1/data/bc/"   //{id}
	config.Client.URI.DeleteTxt = "/api/v1/data/txt/" //{id}
	config.Client.URI.DeleteBM = "/api/v1/data/bm/"   //{id}
	config.Client.URI.PutLP = "/api/v1/data/lp"
	config.Client.URI.PutBC = "/api/v1/data/bc"
	config.Client.URI.PutTxt = "/api/v1/data/txt"
	config.Client.URI.PutBM = "/api/v1/data/bm"
	//***

	fmt.Println("go-keeper-client")
	buildinfo.PrintBuildVersionDate()

	//interface
	clientapp.MainInterface()
	clientapp.InnerInterface()
	// clientapp.GetData(&clientapp.User)

	// secretString, err := encrypt.EncryptData(clientapp.User.User.Password, clientapp.User.User.Password)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(secretString)
	// fmt.Println(encrypt.DecryptData(clientapp.User.User.Password, secretString))
}
