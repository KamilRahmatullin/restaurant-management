package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kamilrahmatullin/restaurant-management/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	invoiceGroup := incomingRoutes.Group("/invoices")
	{
		invoiceGroup.GET("/", controller.GetInvoices())
		invoiceGroup.GET("/:invoice_id", controller.GetInvoice())
		invoiceGroup.POST("/", controller.CreateInvoice())
		invoiceGroup.PATCH("/:invoice_id", controller.UpdateInvoice())
	}
}
