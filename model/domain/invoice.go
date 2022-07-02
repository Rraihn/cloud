package domain

import "time"

type Invoice struct {
	Invoice_id     int
	Invoice_Date   string
	Tax            float64
	Price          float64
	Invoice_number int
	Total          float64
	Created_at     time.Time
	Updated_at     time.Time
}
