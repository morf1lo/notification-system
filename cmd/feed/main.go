package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/morf1lo/notification-system/internal/feed/config"
	"github.com/morf1lo/notification-system/internal/feed/handler"
	"github.com/morf1lo/notification-system/internal/feed/server"
	"github.com/morf1lo/notification-system/internal/feed/service"
	"github.com/morf1lo/notification-system/pkg/mq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := initEnv(); err != nil {
		logrus.Fatalf("error initializing env: %s", err.Error())
	}

	rabbitMQ, err := mq.New(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		logrus.Fatalf("error connecting to rabbitMQ: %s", err.Error())
	}
	defer rabbitMQ.Close()

	services := service.New(rabbitMQ)
	handlers := handler.New(services)

	srv := server.New()
	serverConfig := &config.ServerConfig{
		Port: viper.GetString("app.port"),
		Handler: handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	go func ()  {
		if err := srv.Run(serverConfig); err != nil {
			logrus.Fatalf("error running server: %s", err.Error())
		}
	}()

	logrus.Print("Feed Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("File Server Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error shutting down server: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("feed")
	return viper.ReadInConfig()
}

func initEnv() error {
	return godotenv.Load("feed.env")
}
