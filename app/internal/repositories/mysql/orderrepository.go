package mysql

import (
	"github.com/onpavlov/mysql_to_clickhouse_transfer/internal/model"

	"github.com/bxcodec/faker/v4"
	"github.com/jmoiron/sqlx"
)

type Faker struct {
	TestString string `faker:"oneof: test1, test2, test3"`
}

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	repo := &OrderRepository{
		db: db,
	}

	return repo
}

func (r *OrderRepository) Get(limit int) ([]model.Order, error) {
	var orders []model.Order
	dbOrders := []model.DbOrder{}

	r.db.Select(&dbOrders, "SELECT ID, USER_ID, DATE_INSERT, DATE_STATUS, STORE_ID, STATUS_ID, PRICE, PRICE_DELIVERY FROM b_sale_order LIMIT ?", limit)
	// fmt.Printf("%+v", orders)
	// os.Exit(1)
	for _, dbOrder := range dbOrders {
		fakeOrder := r.getFakeOrder()
		fakeOrder.OrderId = dbOrder.OrderId
		fakeOrder.CreatedAt = dbOrder.CreatedAt
		// fakeOrder.Status.Name = dbOrder.StatusName
		// fakeOrder.Status.UpdatedAt = dbOrder.StatusUpdatedAt
		fakeOrder.UserId = dbOrder.UserId
		// fakeOrder.Store = dbOrder.Store
		// fakeOrder.Amount.Delivery = dbOrder.PriceDelivery
		// fakeOrder.Amount.Total = dbOrder.PriceTotal
		orders = append(orders, fakeOrder)
	}

	return orders, nil
}

func (r *OrderRepository) FindById(id int) (model.Order, error) {
	var dbOrder model.DbOrder

	r.db.Get(&dbOrder, "SELECT ID, USER_ID, DATE_INSERT, DATE_STATUS, STORE_ID, STATUS_ID, PRICE, PRICE_DELIVERY FROM b_sale_order WHERE ID=?", id)

	fakeOrder := r.getFakeOrder()
	fakeOrder.OrderId = dbOrder.OrderId
	fakeOrder.CreatedAt = dbOrder.CreatedAt
	// fakeOrder.Status.Name = dbOrder.StatusName
	// fakeOrder.Status.UpdatedAt = dbOrder.StatusUpdatedAt
	fakeOrder.UserId = dbOrder.UserId
	// fakeOrder.Store = dbOrder.Store
	// fakeOrder.Amount.Delivery = dbOrder.PriceDelivery
	// fakeOrder.Amount.Total = dbOrder.PriceTotal

	return fakeOrder, nil
}

func (r *OrderRepository) getFakeOrder() model.Order {
	order := model.Order{}

	err := faker.FakeData(&order)
	if err != nil {
		panic(err.Error())
	}

	return order
}
