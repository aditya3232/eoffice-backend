package user

import (
	"strconv"
	"time"
)

type User struct {
	ID         int        `gorm:"column:id;primary_key" json:"id"`
	CreatedAt  time.Time  `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;default:now()" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Remarks    string     `gorm:"column:remarks" json:"remarks"`
	EmployeeID int        `gorm:"column:employee_id" json:"employee_id"`
	Password   string     `gorm:"column:password" json:"password"`
	Token      string     `gorm:"column:token" json:"token"`
	RoleID     int        `gorm:"column:role_id" json:"role_id"`
	LastLogin  *time.Time `gorm:"column:last_login" json:"last_login"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) RedisKey() string {
	if u.ID == 0 {
		return "users"
	}

	return "users:" + strconv.Itoa(u.ID)
}
