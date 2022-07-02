package main

import (
	"booking-hotel/app"
	"booking-hotel/controller"
	"booking-hotel/helper"
	"booking-hotel/middleware"
	"booking-hotel/repository"
	"booking-hotel/service"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	discountRepostiory := repository.NewDiscountRepository()
	discountService := service.NesDiscountService(discountRepostiory, db, validate)
	discountController := controller.NewDiscountController(discountService)

	invoiceRepository := repository.NewInvoiceRepository()
	invoiceService := service.NewInvoiceService(invoiceRepository, db, validate)
	invoiceController := controller.NewInvoiceController(invoiceService)

	router := app.NewRouter(discountController, invoiceController)

	server := http.Server{
		Addr:    "localhost:3080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
