package web

import "time"

type InvoiceResponse struct {
	Invoice_id     int       `json:"invoice_Id"`
	Invoice_Date   string    `json:"invoice_Date"`
	Tax            float64   `json:"tax"`
	Price          float64   `json:"price"`
	Invoice_number int       `json:"invoice_Number"`
	Total          float64   `json:"total"`
	Created_at     time.Time `json:"created_At"`
	Updated_at     time.Time `json:"updated_At"`
}
