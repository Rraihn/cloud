package helper

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
)

func ToDiscountResponse(disc domain.Discount) web.DiscountResponse {
	return web.DiscountResponse{
		Discount_id:      disc.Discount_id,
		Discount_request: disc.Discount_request,
		Discount_status:  disc.Discount_status,
		Discount_amount:  disc.Discount_amount,
		Invoice_id:       disc.Invoice_id,
		Created_at:       disc.Created_at,
		Updated_at:       disc.Updated_at,
	}
}

func ToDiscountResponses(disc []domain.Discount) []web.DiscountResponse {
	var discountResponses []web.DiscountResponse
	for _, ordeer := range disc {
		discountResponses = append(discountResponses, ToDiscountResponse(ordeer))
	}
	return discountResponses
}
