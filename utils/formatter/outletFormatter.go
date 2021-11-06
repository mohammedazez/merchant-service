package formatter

import (
	"merchant-service/domain/dto"
	"time"
)

type OutletFormat struct {
	ID         string `json:"id"`
	OutletName string `json:"outlet_name"`
	Picture    string `json:"picture"`
	UserID     string `json:"user_id"`
}

type OutletDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatOutlet(outlet dto.Outlet) OutletFormat {
	var formatOutlet = OutletFormat{
		ID:         outlet.ID,
		OutletName: outlet.OutletName,
		Picture:    outlet.Picture,
		UserID:     outlet.UserID,
	}

	return formatOutlet
}

func FormatDeleteOutlet(msg string) OutletDeleteFormat {
	var deleteFormat = OutletDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
