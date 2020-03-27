package services

import (
	"SimpleShopping/datamodels"
	"SimpleShopping/models"
	"SimpleShopping/validators"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type OrderService interface {
	CreateOrder(order *validators.Order) string
	GetOrders(orderId string) *[]datamodels.Order
}

func NewOrderService() OrderService {
	return &orderService{
		db: models.DB.Mysql,
	}
}

type orderService struct {
	db *gorm.DB
}

func (s *orderService) CreateOrder(order *validators.Order) string {
	var totalAmount uint32 = 0
	for index := range order.OrderDetails {
		orderDetail := &order.OrderDetails[index]
		totalAmount += viper.GetUint32(fmt.Sprintf("commodityprice%d", orderDetail.CommodityId)) * orderDetail.Quantity
	}

	totalAmount += viper.GetUint32("freight")
	newOrder := models.Order{Name: order.Name, LogisticsType: order.LogisticsType, ShopName: order.ShopName, TotalAmount: cast.ToUint16(totalAmount)}
	if err := s.db.Create(&newOrder).Error; err != nil {
		return err.Error()
	}

	for index := range order.OrderDetails {
		orderDetail := &order.OrderDetails[index]
		price := cast.ToUint16(viper.GetInt(fmt.Sprintf("commodityprice%d", orderDetail.CommodityId)))
		newOrderDetail := models.OrderDetail{OrderId: newOrder.OrderId, CommodityId: orderDetail.CommodityId, Price: price, Quantity: orderDetail.Quantity}
		if err := s.db.Create(&newOrderDetail).Error; err != nil {
			return err.Error()
		}
	}
	return ""
}

func (s *orderService) GetOrders(orderId string) *[]datamodels.Order {
	var orders []datamodels.Order
	tempSql := s.db.Preload("OrderDetails", func(db *gorm.DB) *gorm.DB {
		return db.Select("order_id, commodity_id, price, quantity")
	}).Order("order_id desc").Limit(30)

	if orderId != "" {
		tempSql = tempSql.Where("order_id < ?", orderId)
	}

	tempSql.Find(&orders)

	return &orders
}
