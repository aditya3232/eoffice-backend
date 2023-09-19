package rolepermission

import (
	"strconv"
	"time"
)

type RolePermission struct {
	ID           int        `gorm:"column:id;primary_key" json:"id"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Remarks      string     `gorm:"column:remarks" json:"remarks"`
	RoleID       int        `gorm:"column:role_id" json:"role_id"`
	PermissionID int        `gorm:"column:permission_id" json:"permission_id"`
}

func (m *RolePermission) TableName() string {
	return "role_permissions"
}

func (m *RolePermission) RedisKey() string {
	if m.ID == 0 {
		return "role_permissions"
	}

	return "role_permissions:" + strconv.Itoa(m.ID)
}
