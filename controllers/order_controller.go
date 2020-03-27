package controllers

import (
	"SimpleShopping/services"
	"SimpleShopping/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"sync"
)

type OrderController struct {
	OrderService services.OrderService
}

var orderController *OrderController
var orderControllerOnce sync.Once

func GetOrderController() *OrderController {
	orderControllerOnce.Do(func() {
		orderController = &OrderController{
			OrderService: services.NewOrderService(),
		}
	})
	return orderController
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order validators.Order
	ctx.ShouldBind(&order)
	if err := validators.GlobalValidator.Check(order); err != nil {
		ctx.JSON(http.StatusBadRequest, strings.Split(err.Error(), "|"))
		return
	}
	for _, orderDetail := range order.OrderDetails {
		if err := validators.GlobalValidator.Check(orderDetail); err != nil {
			ctx.JSON(http.StatusBadRequest, strings.Split(err.Error(), "|"))
			return
		}
	}

	result := c.OrderService.CreateOrder(&order)
	if result == "" {
		ctx.JSON(http.StatusOK, "success")
	} else {
		ctx.JSON(http.StatusBadRequest, []string{result})
	}
}

func (c *OrderController) GetOrders(ctx *gin.Context) {
	orderId := ctx.Query("orderId")
	orders := c.OrderService.GetOrders(orderId)
	ctx.JSON(http.StatusOK, orders)
}
