package main

import (
	"nashimenshie_api/internal/handler"
	"nashimenshie_api/internal/repository"
	"nashimenshie_api/internal/server"
	"nashimenshie_api/internal/service"
	"nashimenshie_api/pkg/client/moysklad"
	"nashimenshie_api/pkg/logging"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	logger := logging.NewLogger()

	if err := initConfig(); err != nil {
		logger.Fatal("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatal("error loading env variables: %s", err.Error())
	}

	moyskladClient := moysklad.NewMoyskladClient(os.Getenv("MOYSKLAD_AUTH_TOKEN"))
	repository := repository.NewRepository(moyskladClient, logger)
	service := service.NewService(repository, logger)
	handler := handler.NewHandler(service, logger)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("http.port"), handler.Init()); err != nil {
		logger.Fatal("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("main")
	return viper.ReadInConfig()
}
