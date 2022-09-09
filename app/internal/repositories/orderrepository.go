package repositories

import "github.com/onpavlov/mysql_to_clickhouse_transfer/internal/model"

type OrderRepository interface {
	Get(limit int) ([]model.Order, error)
	FindById(id int) (model.Order, error)
}
