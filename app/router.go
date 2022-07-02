package app

import (
	"booking-hotel/controller"
	"booking-hotel/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(discountController controller.DiscountController, invoiceController controller.InvoiceController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/discount", discountController.FindAllDiscount)
	router.GET("/api/discount/:discountId", discountController.FindDiscountById)
	router.POST("/api/discount", discountController.CreateDiscount)
	router.PUT("/api/discount/:discountId", discountController.UpdateDiscount)
	router.DELETE("/api/discount/:discountId", discountController.DeleteDiscount)

	router.GET("/api/invoice", invoiceController.FindAllInvoice)
	router.GET("/api/invoice/:invoiceId", invoiceController.FindInvoiceById)
	router.POST("/api/invoice", invoiceController.CreateInvoice)
	router.PUT("/api/invoice/:invoiceId", invoiceController.UpdateInvoice)
	router.DELETE("/api/invoice/:invoiceId", invoiceController.DeleteInvoice)

	router.PanicHandler = exception.ErrorHandler

	return router
}
