package location

import (
	"strconv"
	"time"
)

type Location struct {
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:now()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Remarks   string     `gorm:"column:remarks" json:"remarks"`
	Nama      string     `gorm:"column:nama" json:"nama"`
}

func (m *Location) TableName() string {
	return "locations"
}

func (e *Location) RedisKey() string {
	if e.ID == 0 {
		return "locations"
	}

	return "locations:" + strconv.Itoa(e.ID)
}
