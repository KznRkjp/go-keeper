package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/KznRkjp/go-keeper.git/internal/buildinfo"
	"github.com/KznRkjp/go-keeper.git/internal/flags"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/router"
	"go.uber.org/zap"
)

var HTTPS bool

func main() {
	//Печатем билд и дату
	buildinfo.PrintBuildVersionDate()

	//создаем экземпляр логгера
	mlogger.Logger = zap.Must(zap.NewProduction())
	defer mlogger.Logger.Sync()

	//получаем переменные для запуска
	flags.ParseFlags()

	//TODO: create database

	//TODO: https
	HTTPS := false
	//server
	dd := router.Main()
	server := &http.Server{
		Handler: dd,
		Addr:    flags.FlagRunAddr,
	}

	if HTTPS {
		go func() {
			mlogger.Logger.Info("Server started", zap.String("address:", flags.FlagRunAddr))
			err := server.ListenAndServeTLS("server.crt", "server.key")
			if err != nil {
				log.Println(err)
			}
		}()

	} else {
		go func() {
			mlogger.Logger.Info("Server started", zap.String("address:", flags.FlagRunAddr))
			if err := server.ListenAndServe(); err != nil {
				// записываем в лог ошибку, если сервер не запустился
				mlogger.ServerStartLog(err.Error())
			}
		}()
	}

	//Gracefull shutdown
	// через этот канал сообщим основному потоку, что соединения закрыты
	idleConnsClosed := make(chan struct{})
	// канал для перенаправления прерываний
	// поскольку нужно отловить всего одно прерывание,
	// ёмкости 1 для канала будет достаточно
	sigint := make(chan os.Signal, 1)
	// регистрируем перенаправление прерываний
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	go func() {
		// читаем из канала прерываний
		// поскольку нужно прочитать только одно прерывание,
		// можно обойтись без цикла
		<-sigint
		// получили сигнал os.Interrupt, запускаем процедуру graceful shutdown
		if err := server.Shutdown(context.Background()); err != nil {
			// ошибки закрытия Listener
			log.Printf("HTTP server Shutdown: %v", err)
		}
		// сообщаем основному потоку,
		// что все сетевые соединения обработаны и закрыты
		close(idleConnsClosed)
	}()
	// ждём завершения процедуры graceful shutdown
	<-idleConnsClosed
	// получили оповещение о завершении
	// здесь можно освобождать ресурсы перед выходом,
	// например закрыть соединение с базой данных,
	// закрыть открытые файлы
	fmt.Println("Server Shutdown gracefully")

}
