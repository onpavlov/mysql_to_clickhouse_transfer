package config

import (
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/repositories"
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/repositories/mysql"
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/services"
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/services/rabbitmq"

	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Repositories repositories.Repositories
	db           *sqlx.DB
	amqp         *amqp.Channel
	Amqp         services.Amqp
}

func NewConfig() *Config {
	c := &Config{}
	var err error

	c.db, err = c.connectDb()
	c.failOnError(err)

	rabbitConn, err := amqp.Dial("amqp://guest:guest@docker.for.mac.localhost:5672/")
	c.failOnError(err)

	c.amqp, err = rabbitConn.Channel()
	c.failOnError(err)

	c.Amqp = rabbitmq.NewRabbitMQ(*c.amqp)

	c.Repositories = repositories.Repositories{
		Order: mysql.NewOrderRepository(c.db),
	}

	return c
}

func (c *Config) connectDb() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "nameless:secret@tcp(docker.for.mac.localhost:3306)/clickhouse_test")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (c *Config) CloseDb() {
	c.db.Close()
}

func (c *Config) CloseAmqp() {
	c.amqp.Close()
}

func (c *Config) failOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
