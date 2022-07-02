package booking_hotel

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

	router := app.NewRouter(discountController)

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
