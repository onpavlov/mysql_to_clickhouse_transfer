package main

import (
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/config"
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/controllers"
)

func main() {
	c := config.NewConfig()
	ctrl := controllers.NewExporter(c)

	err := ctrl.ExportOrders()
	if err != nil {
		panic(err.Error())
	}

	c.CloseAmqp()
	c.CloseDb()
}
