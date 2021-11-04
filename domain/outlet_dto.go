package domain

type GetOutletsResponse struct {
	Status  int64           `json:"status"`
	Meta    MetaDetail      `json:"meta"`
	Data    []GetAllOutlets `json:"data"`
	Message string          `json:"message"`
}

type GetAllOutlets struct {
	Id         string `gorm:"primaryKey" json:"id"`
	OutletName string `json:"outlet_name"`
	UserId     string `json:"user_id"`
}

type Outlet struct {
	Id         string    `gorm:"primaryKey" json:"id"`
	OutletName string    `json:"outlet_name"`
	UserId     string    `json:"user_id"`
	Products   []Product `gorm:"ForeignKey:OutletId"`
}

type InputOutlet struct {
	OutletName string `json:"outlet_name" binding:"required"`
}

type UpdateOutlet struct {
	OutletName string `json:"outlet_name"`
}
