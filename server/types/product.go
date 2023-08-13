package types

type (
	AddNewProductBody struct {
		Name    string  `json:"name"`
		Price   float64 `json:"price"`
		InStock int32    `json:"in_stock"`
	}
	RateProductBody struct {
		ProductId string `json:"product_id"`
		Rating    int    `json:"rating"`
	}
)
