package app

import (
	"booking-hotel/controller"
	"booking-hotel/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(discountController controller.DiscountController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/discount", discountController.FindAllDiscount)
	router.GET("/api/discount/:discountId", discountController.FindDiscountById)
	router.POST("/api/discount", discountController.CreateDiscount)
	router.PUT("/api/discount/:discountId", discountController.UpdateDiscount)
	router.DELETE("/api/discount/:discountId", discountController.DeleteDiscount)

	router.PanicHandler = exception.ErrorHandler

	return router
}