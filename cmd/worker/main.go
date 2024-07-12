package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/morf1lo/notification-system/internal/user/pb"
	"github.com/morf1lo/notification-system/internal/worker/service"
	"github.com/morf1lo/notification-system/pkg/mq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := initEnv(); err != nil {
		logrus.Fatalf("error initializing env: %s", err.Error())
	}

	userServiceConn, err := grpc.NewClient(
		viper.GetString("userService.target"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 32)),
	)
	if err != nil {
		logrus.Fatalf("error connecting to gRPC user service: %s", err.Error())
	}
	defer userServiceConn.Close()

	userService := pb.NewUserClient(userServiceConn)

	rabbitMQ, err := mq.New(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		logrus.Fatalf("error connecting to rabbitMQ: %s", err.Error())
	}
	defer rabbitMQ.Close()

	services := service.New(userService, rabbitMQ)

	go services.Feed.ProcessFeeds(context.Background())

	logrus.Print("Worker Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Worker Shutting Down")
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("worker")
	return viper.ReadInConfig()
}

func initEnv() error {
	return godotenv.Load("worker.env")
}
