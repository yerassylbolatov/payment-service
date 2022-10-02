package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	paymentservice "github.com/yerassylbolatov/payment-service"
	"github.com/yerassylbolatov/payment-service/pkg/handler"
	"github.com/yerassylbolatov/payment-service/pkg/repository"
	"github.com/yerassylbolatov/payment-service/pkg/service"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error at initConfig occured %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env parameters: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DDName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslMode")})
	if err != nil {
		log.Fatalf("failed to initialize db %s", err.Error())
	}
	repos := repository.NewRepository(db)    // business logic layer
	services := service.NewService(repos)    // database layer
	handlers := handler.NewHandler(services) // handlers layer

	srv := new(paymentservice.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
