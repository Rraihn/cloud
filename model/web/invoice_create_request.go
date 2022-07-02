package web

type InvoiceCreateRequest struct {
	Tax            float64 `validate:"required,min=1" json:"tax"`
	Price          float64 `validate:"required,min=1" json:"price"`
	Invoice_number int     `validate:"required,min=1" json:"invoice_Number"`
	Total          float64 `validate:"required,min=1" json:"total"`
}
