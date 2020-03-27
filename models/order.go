package models

import (
	"time"
)

type Order struct {
	OrderId       uint32 `gorm:"primary_key"`
	Name          string
	LogisticsType uint8
	ShopName      string
	TotalAmount   uint16
	CreatedAt     time.Time
}

type OrderDetail struct {
	OdId        uint32 `gorm:"primary_key"`
	OrderId     uint32
	CommodityId uint8
	Price       uint16
	Quantity    uint32
}
