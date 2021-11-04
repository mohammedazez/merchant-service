package schema

import (
	"database/sql"
	"strconv"
	"time"
)

type BaseLegacy struct {
	Id        string         `gorm:"type:int(11) AUTO_INCREMENT;primaryKey"`
	CreatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedBy sql.NullString `gorm:"type:varchar(255)"`
	UpdatedBy sql.NullString `gorm:"type:varchar(255)"`
	DeletedAt *time.Time     `gorm:"type:datetime;index"`
	Cursor    string         `gorm:"-"`
}

func (b BaseLegacy) GetLegacyId() int64 {
	id, _ := strconv.ParseInt(b.Id, 10, 64)
	return id
}

func (b *BaseLegacy) SetId(id interface{}) {
	switch idi := id.(type) {
	case string:
		b.Id = idi
	case int:
		b.Id = strconv.Itoa(idi)
	case int32:
		b.Id = strconv.FormatInt(int64(idi), 10)
	case int64:
		b.Id = strconv.FormatInt(idi, 10)
	}
}
