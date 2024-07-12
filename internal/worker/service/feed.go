package service

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/morf1lo/notification-system/internal/user/pb"
	"github.com/morf1lo/notification-system/internal/worker/model"
	"github.com/morf1lo/notification-system/pkg/mq"
	"github.com/sirupsen/logrus"
)

type FeedService struct {
	userService pb.UserClient
	rabbitMQ *mq.MQConn
}

func NewFeedService(userService pb.UserClient, rabbitMQ *mq.MQConn) *FeedService {
	return &FeedService{
		userService: userService,
		rabbitMQ: rabbitMQ,
	}
}

func (s *FeedService) ProcessFeeds(ctx context.Context) {
	msgs, err := s.rabbitMQ.Consume(articleEmailNotificationMQ)
	if err != nil {
		logrus.Fatalf("error starting email consumer: %s", err.Error())
	}

	forever := make(chan bool, 1)

	go func ()  {
		for msg := range msgs {
			var message model.Article
			if err := json.Unmarshal(msg.Body, &message); err != nil {
				logrus.Errorf("failed unmarshal message body: %s", err.Error())
				msg.Nack(false, true)
				continue
			}

			subs, err := s.userService.GetAllSubscribers(ctx, &pb.Empty{})
			if err != nil {
				logrus.Errorf("failed get users list: %s", err.Error())
				msg.Nack(false, true)
				continue
			}

			var wg sync.WaitGroup

			for _, sub := range subs.GetSubs() {
				wg.Add(1)
				go func(sub *pb.Subscriber)  {
					defer wg.Done()
					if err := sendMail(sub.GetEmail(), &message); err != nil {
						logrus.Errorf("error sending email to %s: %s", sub.GetEmail(), err.Error())
					}
				}(sub)
			}

			wg.Wait()

			msg.Ack(false)

			logrus.Print(message)
			logrus.Print("Messages have been delivered successfully")
		}
	}()

	<-forever
}

func sendMail(to string, article *model.Article) error {
	// TODO: send mail functional
	return nil
}
