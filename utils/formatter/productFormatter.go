package formatter

import "time"

type ProductDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatDeleteProduct(msg string) ProductDeleteFormat {
	var deleteFormat = ProductDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
