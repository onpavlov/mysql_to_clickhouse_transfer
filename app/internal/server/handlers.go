package server

import (
	"net/http"
	"text/template"

	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/model"
)

func (s *API) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Everithing works fine</h1>"))
}

func (s *API) OrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := s.repositories.Order.Get(50)
	if err != nil {
		panic(err.Error())
	}

	tmpl, err := template.ParseFiles("/app/internal/templates/orders.html")
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, model.OrdersData{
		PageTitle: "Orders",
		Orders:    orders,
	})
}
