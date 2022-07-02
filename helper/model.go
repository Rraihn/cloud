package helper

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"time"
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
	for _, discounts := range disc {
		discountResponses = append(discountResponses, ToDiscountResponse(discounts))
	}
	return discountResponses
}

func ToInvoiceResponse(invcs domain.Invoice) web.InvoiceResponse {
	return web.InvoiceResponse{
		Invoice_id:     invcs.Invoice_id,
		Invoice_Date:   invcs.Invoice_Date,
		Tax:            invcs.Tax,
		Price:          invcs.Price,
		Invoice_number: invcs.Invoice_number,
		Total:          invcs.Total,
		Created_at:     time.Now(),
		Updated_at:     time.Now(),
	}
}

func ToInvoiceResponses(invcs []domain.Invoice) []web.InvoiceResponse {
	var invoiceResponses []web.InvoiceResponse
	for _, invoices := range invcs {
		invoiceResponses = append(invoiceResponses, ToInvoiceResponse(invoices))
	}
	return invoiceResponses
}
