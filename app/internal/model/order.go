package model

import "database/sql"

type DbOrder struct {
	OrderId         int           `db:"ID"`
	CreatedAt       string        `db:"DATE_INSERT"`
	StatusName      string        `db:"STATUS_ID"`
	StatusUpdatedAt string        `db:"DATE_STATUS"`
	UserId          int           `db:"USER_ID"`
	PriceTotal      float32       `db:"PRICE"`
	PriceDelivery   float32       `db:"PRICE_DELIVERY"`
	Store           sql.NullInt32 `db:"STORE_ID"`
}

type Order struct {
	OrderId        int
	MindboxOrderId int32
	CreatedAt      string `faker:"timestamp"`
	// Status         Status
	UserId int
	// Amount         Amount
	Platform string `faker:"oneof: ios, android, web"`
	// Promotion      Promotion
	// Bonuses        Bonuses
	DeliveryType string
	PayedType    string
	// Store        sql.NullInt32 `faker:"-"`
	City string
	// Quantity       Quantity
}

type Status struct {
	Name      string `faker:"oneof: N, R, F"`
	UpdatedAt string `faker:"timestamp"`
}

type Amount struct {
	Total    float32
	Delivery float32
	Order    float32
	Discount float32
}

type Promotion struct {
	PromoCode string
	Type      string
	SberPrime bool
}

type Bonuses struct {
	Withdraw float32
	Earned   float32
	Spasibo  float32
}

type Quantity struct {
	Quantity int32
}

type OrdersData struct {
	PageTitle string
	Orders    []Order
}
