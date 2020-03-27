package datamodels

import "time"

type Order struct {
	OrderId       uint32 `gorm:"primary_key"`
	Name          string
	LogisticsType uint8
	ShopName      string
	TotalAmount   uint16
	CreatedAt     time.Time
	OrderDetails  []OrderDetail `gorm:"ForeignKey:OrderId"`
}

type OrderDetail struct {
	OrderId     uint32
	CommodityId uint8
	Price       uint16
	Quantity    uint32
}
