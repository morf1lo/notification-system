package mq

import amqp "github.com/rabbitmq/amqp091-go"

type MQConn struct {
	conn *amqp.Connection
}

func New(url string) (*MQConn, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return &MQConn{conn: conn}, nil
}

func (mq *MQConn) Close() error {
	return mq.conn.Close()
}

func (mq *MQConn) Channel() (*amqp.Channel, error) {
	return mq.conn.Channel()
}

func (mq *MQConn) Publish(queue string, body []byte) error {
	ch, err := mq.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body: body,
		},
	)
}

func (mq *MQConn) Consume(queue string) (<-chan amqp.Delivery, error) {
	ch, err := mq.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}
