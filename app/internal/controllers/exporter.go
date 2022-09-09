package controllers

import (
	"encoding/json"

	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/config"
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/model"
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/repositories"
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/services"
)

type exporter struct {
	repositories repositories.Repositories
	amqp         services.Amqp
	// clickhouse   services.Clickhouse
}

func NewExporter(c *config.Config) exporter {
	return exporter{
		repositories: c.Repositories,
		amqp:         c.Amqp,
	}
}

func (e *exporter) ExportOrders() error {
	var orders []model.Order
	// order, err := e.repositories.Order.FindById(1)
	orders, err := e.repositories.Order.Get(1000000)
	if err != nil {
		return err
	}

	// orders = append(orders, order)

	err = e.amqp.DeclareQueue("1_default_orders")
	if err != nil {
		return err
	}

	for _, order := range orders {
		body, err := json.Marshal(order)
		if err != nil {
			return err
		}

		e.amqp.Produce([]byte(body))
	}

	return nil
}
