package handlers

import (
	"errors"
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

	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	// this is where each product in that products array will be handled by a separate go routing

	order, err := db.Conn.CreateOrder(c, db.CreateOrderParams{
		UserID: user.UserId,
		Status: utils.ORDER_INITIATED,
	})

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	var resp []types.PlaceOrderResponse
	total := 0
	actual_total := 0
	for _, order_product := range body.Products {
		placed := true
		product_id, err := uuid.Parse(order_product.ProductId)
		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}
		product, err := db.Conn.GetProductById(c, product_id)
		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}
		if product.InStock < order_product.Quantity {
			placed = false
		}

		db.Conn.CreateOrderProduct(c, db.CreateOrderProductParams{
			OrderID:   order.ID,
			ProductID: product_id,
			Quantity:  order_product.Quantity,
			Placed:    placed,
		})

		if placed {
			_, err := db.Conn.UpdateProductInStockUnits(c, db.UpdateProductInStockUnitsParams{
				InStock: product.InStock - order_product.Quantity,
				ID:      product.ID,
			})

			if err != nil {
				// error
			}

			total += int(product.Price * float64(order_product.Quantity))
		}
		actual_total += int(product.Price * float64(order_product.Quantity))

		resp = append(resp, types.PlaceOrderResponse{
			Product_Id: order_product.ProductId,
			Placed:     placed,
		})
	}

	msg := "Order Placed. Thanks"
	if total < actual_total {
		msg = "Order Placed partially. We'll process the missing products as soon as they are available. Thanks"
	}
	utils.SendResponse(c, msg, map[string]any{
		"products":    resp,
		"order_total": total,
	})
	return
}

func ProcessOrderHandler(c *gin.Context) {
	var body types.ProcessOrderBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	order_id, err := uuid.Parse(body.OrderId)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	order, err := db.Conn.GetOrderById(c, order_id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	if order.Status == "COMPLETED" {
		utils.SendError(c, http.StatusBadRequest, errors.New("Order is already completed"))
		return
	}

	_, err = db.Conn.UpdateOrderStatus(c, db.UpdateOrderStatusParams{
		ID:     order_id,
		Status: body.Status,
	})

	utils.SendResponse(c, "Order Updated", map[string]any{})
	return
}

func GetOrderDetails(c *gin.Context) {
	order_id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		// error
	}

	order, err := db.Conn.GetOrderDetails(c, order_id)
	if err != nil {
		// error
	}

	if order == nil || len(order) == 0 {
		// error - no such order
	}

	var products []types.OrderProduct
	user_name := order[0].UserName
	user_email := order[0].Email
	order_total := 0

	for _, p := range order {
		products = append(products, types.OrderProduct{
			Id:       p.ProductID.String(),
			Name:     p.Name,
			Quantity: p.Quantity,
			Price:    p.Price,
			Placed:   p.Placed,
		})

		order_total += (int(p.Price) * int(p.Quantity))

	}
	utils.SendResponse(c, "Order Details", map[string]any{

		"user_name":   user_name,
		"user_email":  user_email,
		"products":    products,
		"order_total": order_total,
	})
	return
}
