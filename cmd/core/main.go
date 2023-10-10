package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"igo/iternal/app/common/db"
	handlerCommon "igo/iternal/app/common/handler"
	server2 "igo/iternal/app/common/server"
	"igo/iternal/app/healthCheck/handler"
	"igo/iternal/app/healthCheck/repository"
	"igo/iternal/app/healthCheck/service"
	handler2 "igo/iternal/app/user/handler"
	repository2 "igo/iternal/app/user/repository"
	service2 "igo/iternal/app/user/service"
	"log"
	"net/http"
)

func main() {
	// 1. Load config variables
	loadConfigVariables()

	loadApp()

	// 2. Declare API methods and handlers

	serverUrl := viper.GetString("server.url")
	fmt.Printf("Server running, route: %s/helloworld\n", serverUrl)

	// 3. Start server
	if err := http.ListenAndServe(serverUrl, nil); err != nil {
		log.Fatal(err)
	}

}

func loadConfigVariables() {
	viper.AddConfigPath("config")
	viper.SetConfigName("test")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func loadApp() {
	database := loadDB()
	healthCheckRepository := repository.NewHealthCheckRepository(database)
	healthCheckService := service.NewHealthCheckService(healthCheckRepository)
	healthCheckHandler := handler.NewHealthCheckHandler(healthCheckService)

	userRepository := repository2.NewUserRepository(database)
	userService := service2.NewUserService(userRepository)
	userListHandler := handler2.NewUserListHandler(userService)

	initServer(healthCheckHandler, userListHandler)
}

func loadDB() *sqlx.DB {
	config := db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DbName:   viper.GetString("db.db_name"),
		Login:    viper.GetString("db.db_user"),
		Password: viper.GetString("db.db_pass"),
	}

	dataBase, err := db.NewPostgres(config)

	if err != nil {
		log.Fatal("Fail connect to DB")
	}

	return dataBase
}

func initServer(healthCheckHandler handler.HealthCheckHandler, userListHandler handler2.UserListHandler) {
	globalHandler := handlerCommon.NewHandler(healthCheckHandler, userListHandler)
	server := new(server2.Server)

	if err := server.Run(viper.GetString("server.host"), viper.GetString("server.port"), globalHandler.Init()); err != nil {
		log.Fatalf("Ошибка при старте сервиса: %s", err.Error())
	}
}
