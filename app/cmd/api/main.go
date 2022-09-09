package main

import (
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/config"
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/server"
)

func main() {
	c := config.NewConfig()
	s := server.NewAPI(c)
	s.Start()

	defer c.CloseDb()
	c.CloseAmqp()
}
