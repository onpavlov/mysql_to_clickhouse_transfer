package server

import (
	"log"
	"net/http"

	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/config"
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/repositories"

	"github.com/gorilla/mux"
)

type API struct {
	router       *mux.Router
	repositories repositories.Repositories
}

func NewAPI(c *config.Config) *API {
	return &API{
		router:       mux.NewRouter(),
		repositories: c.Repositories,
	}
}

func (s *API) Start() {
	s.registerHandlers()

	log.Fatal(http.ListenAndServe(":8080", s.router))
}

func (s *API) registerHandlers() {
	s.router.HandleFunc("/", s.HomeHandler)
	s.router.HandleFunc("/orders", s.OrdersHandler)
}
