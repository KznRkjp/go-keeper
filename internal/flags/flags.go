// Разбираем входные данные.
// проритет от низшего в высшему строка/env/конфиг файл (но это как то старнно если честно)
package flags

import (
	"flag"
	"os"

	"github.com/KznRkjp/go-keeper.git/internal/config"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
)

// FlagConfigPath содержит путь к файлу конфигурации
var FlagConfigPath string

// FlagRunAddr содержит адрес и порт для запуска сервера
var FlagRunAddr string

// FlagDBString содержит данные для подключения к БД
var FlagDBString string

// FlagBuildVersion содержит номер билда
var FlagBuildVersion string

func ParseFlags() {
	mlogger.Info("Strting parsing flags")
	flag.StringVar(&FlagConfigPath, "c", "", "path to config file")

	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&FlagRunAddr, "a", "localhost:4443", "address and port to run server")

	// регистрируем переменную FlagDBString - для подлкючения к базе данных
	flag.StringVar(&FlagDBString, "d", "", "String for DB connection")

	// регистрируем переменную FlagDBString - для подлкючения к базе данных
	flag.StringVar(&FlagBuildVersion, "b", "0.0.0-a.1", "Build version")

	flag.Parse()

	if envConfigPath := os.Getenv("CONFIG"); envConfigPath != "" {
		FlagConfigPath = envConfigPath
	}

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}

	if envDBString := os.Getenv("DATABASE_DSN"); envDBString != "" {
		FlagDBString = envDBString
	}
	if FlagConfigPath != "" {
		configuration, err := config.OpenConfigFile(FlagConfigPath)
		if err == nil {
			if FlagRunAddr == "" {
				FlagRunAddr = configuration.ServerAddress
			}

			if FlagDBString == "" {
				FlagDBString = configuration.DatabaseDSN
			}

		}
	}
	mlogger.Info("Flags parsed")

}
