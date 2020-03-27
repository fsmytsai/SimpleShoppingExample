package route

import (
	"SimpleShopping/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	apiRoute := app.Group("/api")
	{
		apiRoute.POST("/createOrder", controllers.GetOrderController().CreateOrder)
		apiRoute.GET("/getOrders", controllers.GetOrderController().GetOrders)
	}
}
