package types

type (
	OrderProducts struct {
		ProductId string `json:"product_id"`
		Quantity  int32  `json:"quantity"`
	}
	PlaceOrderBody struct {
		Products []OrderProducts `json:"products"`
	}
	PlaceOrderResponse struct {
		Product_Id string `json:"product_id"`
		Placed     bool   `json:"placed"`
	}
	ProcessOrderBody struct {
		OrderId string `json:"order_id"`
		Status  string `json:"status"`
	}
)
