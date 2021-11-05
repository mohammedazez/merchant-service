package product

import "time"

type ProductDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatDelete(msg string) ProductDeleteFormat {
	var deleteFormat = ProductDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
