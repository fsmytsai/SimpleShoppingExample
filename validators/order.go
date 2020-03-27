package validators

type Order struct {
	Name          string `validate:"required" comment:"姓名"`
	LogisticsType uint8  `validate:"min=1,max=2" comment:"物流類型"`
	ShopName      string `validate:"required" comment:"店名"`
	OrderDetails  []OrderDetail `validate:"min=1" comment:"訂單明細"`
}

type OrderDetail struct {
	CommodityId uint8 `validate:"min=1,max=3" comment:"商品編號"`
	Quantity    uint32 `validate:"min=1,max=20" comment:"數量"`
}
