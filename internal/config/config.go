package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Client ClientConfig

// Стуркрута для разбора файла конфигурации
type Config struct {
	ServerAddress string `json:"server_address"`
	DatabaseDSN   string `json:"database_dsn"`
}

type ClientURI struct {
	RegisterUser string `json:"register_user"`
	LoginUser    string `json:"login_user"`
	GetData      string `json:"get_data"`
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
