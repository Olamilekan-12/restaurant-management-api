package routes

import (
	contrllers "restaurant-management-go/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", contrllers.GetInvoices)
	incomingRoutes.GET("/invoices/:id", contrllers.GetInvoice())
	incomingRoutes.POST("/invoices", contrllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:id", contrllers.UpdateInvoice())
}
