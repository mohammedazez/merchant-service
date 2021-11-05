package outlet

import (
	"merchant-service/entity"
	"time"
)

type OutletFormat struct {
	ID         int    `json:"id"`
	OutletName string `json:"outlet_name"`
	Picture    string `json:"picture"`
	UserID     string `json:"user_id"`
}

type OutletDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func Format(outlet entity.Outlet) OutletFormat {
	var formatOutlet = OutletFormat{
		ID:         outlet.ID,
		OutletName: outlet.OutletName,
		Picture:    outlet.Picture,
		UserID:     outlet.UserID,
	}

	return formatOutlet
}

func FormatDelete(msg string) OutletDeleteFormat {
	var deleteFormat = OutletDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
