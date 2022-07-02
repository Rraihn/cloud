package web

type InvoiceUpdateRequest struct {
	Invoice_id int     `validate:"required"`
	Tax        float64 `validate:"required,min=1" json:"tax"`
	Price      float64 `validate:"required,min=1" json:"price"`
	Total      float64 `validate:"required,min=1" json:"total"`
}
