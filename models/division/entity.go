package division

import (
	"strconv"
	"time"
)

type Division struct {
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	ParentID  *int       `gorm:"column:parent_id" json:"parent_id"`
	CreatedAt time.Time  `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:now()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Remarks   string     `gorm:"column:remarks" json:"remarks"`
	Nama      string     `gorm:"column:nama" json:"nama"`
}

func (m *Division) TableName() string {
	return "divisions"
}

func (e *Division) RedisKey() string {
	if e.ID == 0 {
		return "divisions"
	}

	return "divisions:" + strconv.Itoa(e.ID)
}
