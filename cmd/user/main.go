package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/morf1lo/notification-system/internal/user/config"
	"github.com/morf1lo/notification-system/internal/user/repository"
	"github.com/morf1lo/notification-system/internal/user/repository/postgres"
	"github.com/morf1lo/notification-system/internal/user/server"
	"github.com/morf1lo/notification-system/internal/user/service"
	"github.com/morf1lo/notification-system/internal/user/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := initEnv(); err != nil {
		logrus.Fatalf("error initializing env: %s", err.Error())
	}

	dbConfig := &config.DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}
	db, err := postgres.Connect(ctx, dbConfig)
	if err != nil {
		logrus.Fatalf("error opening database: %s", err.Error())
	}
	defer func ()  {
		if err := db.Close(ctx); err != nil {
			logrus.Fatalf("error closing database connection: %s", err.Error())
		}
	}()

	repo := repository.New(db)
	services := service.New(repo)

	for range 499 {
		repo.Postgres.Subscriber.Create(ctx, &model.Subscriber{
			Email: "0208timur0208@gmail.com",
		})
	}

	serverConfig := &config.GRPCServerConfig{
		Network: viper.GetString("app.network"),
		Addr: viper.GetString("app.addr"),
	}
	go func ()  {
		if err := server.Run(serverConfig, services); err != nil {
			logrus.Fatalf("error running gRPC server: %s", err.Error())
		}
	}()

	logrus.Print("User gRPC Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("User gRPC Server Shutting Down")
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("user")
	return viper.ReadInConfig()
}

func initEnv() error {
	return godotenv.Load("user.env")
}
