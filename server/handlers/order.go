package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/omgupta1608/aftershoot_task/db"
	"github.com/omgupta1608/aftershoot_task/types"
	"github.com/omgupta1608/aftershoot_task/utils"
)

func PlaceOrderHandler(c *gin.Context) {
	var body types.PlaceOrderBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// get user details from context
	user, found := c.Keys["user"].(db.User)
	if !found {
		// error
	}

	// this is where each product in that products array will be handled by a separate go routing

	order, err := db.Conn.CreateOrder(c, db.CreateOrderParams{
		UserID: user.ID,
		Status: utils.ORDER_INITIATED,
	})

	if err != nil {
		// error
	}

	var resp []types.PlaceOrderResponse
	for _, order_product := range body.Products {
		product_id, err := uuid.Parse(order_product.ProductId)
		if err != nil {
			//error
			break
		}
		product, err := db.Conn.GetProductById(c, product_id)
		if err != nil {
			//error
			break
		}
		if product.InStock < order_product.Quantity {
			// error - insufficient
			// this is where retries come in
			resp = append(resp, types.PlaceOrderResponse{
				Product_Id: order_product.ProductId,
				Placed:     false,
			})
			continue
		}

		db.Conn.CreateOrderProduct(c, db.CreateOrderProductParams{
			OrderID:   order.ID,
			ProductID: product_id,
			Quantity:  order_product.Quantity,
		})
		db.Conn.UpdateProductInStockUnits(c, product.InStock-order_product.Quantity)
		resp = append(resp, types.PlaceOrderResponse{
			Product_Id: order_product.ProductId,
			Placed:     true,
		})
	}

	utils.SendResponse(c, "Order Placed", map[string]any{
		"products": resp,
	})
	return
}

func ProcessOrderHandler (c *gin.Context) {

}