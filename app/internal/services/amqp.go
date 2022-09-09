package services

type Amqp interface {
	Produce(body []byte) error
	Consume() (interface{}, error)
	DeclareQueue(queue string) error
}
