package permission

import (
	"strconv"
	"time"
)

type Permission struct {
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	ParentID  *int       `gorm:"column:parent_id" json:"parent_id"`
	CreatedAt time.Time  `gorm:"column:created_at;default:now()," json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:now()," json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Remarks   string     `gorm:"column:remarks" json:"remarks"`
	Nama      string     `gorm:"column:nama" json:"nama"`
	Url       string     `gorm:"column:url" json:"url"`
	Position  int        `gorm:"column:position" json:"position"`
}

func (m *Permission) TableName() string {
	return "permissions"
}

func (m *Permission) RedisKey() string {
	if m.ID == 0 {
		return "permission"
	}

	return "permission:" + strconv.Itoa(m.ID)
}
