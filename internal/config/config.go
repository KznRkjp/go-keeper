package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Client ClientConfig
var Server Config

// Стуркрута для разбора файла конфигурации
type Config struct {
	ServerAddress string `json:"server_address"`
	DatabaseDSN   string `json:"database_dsn"`
	Debug         string `json:"debug"`
}

type ClientURI struct {
	RegisterUser string `json:"register_user"`
	LoginUser    string `json:"login_user"`
	GetData      string `json:"get_data"`
	PostLP       string `json:"post_lp"`
	PostBC       string `json:"post_bc"`
	PostTxt      string `json:"post_txt"`
	PostBM       string `json:"post_bm"`
	DeleteLP     string `json:"delete_lp"`
	DeleteBC     string `json:"delete_bc"`
	DeleteTxt    string `json:"delete_txt"`
	DeleteBM     string `json:"delete_bm"`
	PutLP        string `json:"put_lp"`
	PutBC        string `json:"put_bc"`
	PutTxt       string `json:"put_txt"`
	PutBM        string `json:"put_bm"`
}
type ClientConfig struct {
	ServerAddress string    `json:"server_address"`
	URI           ClientURI `json:"uri"`
}

// func main() {
// 	OpenConfigFile("config.json")
// }

// Разбор файла конфигурации
func OpenConfigFile(filename string) (Config, error) {
	var Conf Config
	f, err := os.Open(filename)
	if err != nil {
		return Conf, err
	}

	fileByte, _ := io.ReadAll(f)
	json.Unmarshal(fileByte, &Conf)
	fmt.Println(Conf)
	return Conf, nil
}

// func ClientConfigData()ClientURI{
// 	// var client ClientURI
// 	// client.ServerAddress
// 	// client.ServerAddress = "http://localhost:4443"
// 	// client.URI.RegisterUser = "/api/v1/register"
// 	// Client.URI.LoginUser = "/api/v1/login"
// 	// Client.URI.GetData = "/api/v1/data"
// 	// Client.URI.PostLP = "/api/v1/data/lp"
// 	// Client.URI.PostBC = "/api/v1/data/bc"
// 	// Client.URI.PostTxt = "/api/v1/data/txt"
// 	// Client.URI.PostBM = "/api/v1/data/bm"
// 	// Client.URI.DeleteLP = "/api/v1/data/lp/"   //{id}
// 	// Client.URI.DeleteBC = "/api/v1/data/bc/"   //{id}
// 	// Client.URI.DeleteTxt = "/api/v1/data/txt/" //{id}
// 	// Client.URI.DeleteBM = "/api/v1/data/bm/"   //{id}
// 	// Client.URI.PutLP = "/api/v1/data/lp"
// 	// Client.URI.PutBC = "/api/v1/data/bc"
// 	// Client.URI.PutTxt = "/api/v1/data/txt"
// 	// Client.URI.PutBM = "/api/v1/data/bm"
// }
