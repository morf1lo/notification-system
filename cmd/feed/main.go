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
	"github.com/morf1lo/notification-system/internal/feed/repository"
	"github.com/morf1lo/notification-system/internal/feed/repository/postgres"
	"github.com/morf1lo/notification-system/internal/feed/server"
	"github.com/morf1lo/notification-system/internal/feed/service"
	"github.com/morf1lo/notification-system/pkg/mq"
	"github.com/redis/go-redis/v9"
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
			logrus.Fatalf("error occured on database connection close: %s", err.Error())
		}
	}()

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	defer func ()  {
		if err := rdb.Close(); err != nil {
			logrus.Fatalf("error occured on redis connection close: %s", err.Error())
		}
	}()

	rabbitMQ, err := mq.New(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		logrus.Fatalf("error connecting to rabbitMQ: %s", err.Error())
	}
	defer func ()  {
		if err := rabbitMQ.Close(); err != nil {
			logrus.Fatalf("error occured on rabbitMQ connection close: %s", err.Error())
		}
	}()

	repos := repository.New(db, rdb)
	services := service.New(repos, rabbitMQ)
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
