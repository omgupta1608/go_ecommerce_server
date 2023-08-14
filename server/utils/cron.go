package utils

import (
	"context"
	"fmt"

	"github.com/omgupta1608/aftershoot_task/db"
	"github.com/robfig/cron/v3"
)

func retryFailedOrders() {
	ctx := context.Background()

	failedOrders, err := db.Conn.GetFailedOrderProducts(ctx)
	if err != nil {
		fmt.Println(fmt.Errorf("Error in 'failedOrderRetry' cron : %s", err.Error()))
		return
	}

	for _, fOrder := range failedOrders {
		product, err := db.Conn.GetProductById(ctx, fOrder.ProductID)
		if err != nil {
			// error
		}

		if product.InStock >= fOrder.Quantity {
			_, err := db.Conn.UpdateOrderProductStatus(ctx, db.UpdateOrderProductStatusParams{
				Placed: true,
				ID:     fOrder.ID,
			})

			if err != nil {
				// error
			}

			_, err = db.Conn.UpdateProductInStockUnits(ctx, db.UpdateProductInStockUnitsParams{
				InStock: product.InStock - fOrder.Quantity,
				ID:      product.ID,
			})

			if err != nil {
				// error
			}
		}
	}
}

func ScheduleJobs() {
	scheduler := cron.New()

	// run retryFailedOrders cron job every day at 6 am
	scheduler.AddFunc("0 6 * * *", retryFailedOrders)

	scheduler.Start()
}