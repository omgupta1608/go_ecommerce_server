package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/db"
	"github.com/omgupta1608/aftershoot_task/types"
	"github.com/omgupta1608/aftershoot_task/utils"
)

func AddNewProductHandler(c *gin.Context) {
	
	var body types.AddNewProductBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	product, err := db.Conn.CreateProduct(c, db.CreateProductParams{
		Name:    body.Name,
		Price:   body.Price,
		InStock: body.InStock,
	})

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Product Added!", map[string]any{

		"product_id":          product.ID.String(),
		"product_name":        product.Name,
		"product_price":       product.Price,
		"is_product_in_stock": product.InStock,
	})
	return
}

func GetProductsHandler (c *gin.Context) {
	
}
