package employee

import (
	"strconv"
	"time"
)

type Employee struct {
	ID             int        `gorm:"column:id;primary_key" json:"id"`
	CreatedAt      time.Time  `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;default:now()" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Remarks        string     `gorm:"column:remarks" json:"remarks"`
	Nama           string     `gorm:"column:nama" json:"nama"`
	Nip            int        `gorm:"column:nip" json:"nip"`
	TempatLahir    string     `gorm:"column:tempat_lahir" json:"tempat_lahir"`
	TanggalLahir   time.Time  `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	Alamat         string     `gorm:"column:alamat" json:"alamat"`
	NoHp           string     `gorm:"column:no_hp" json:"no_hp"`
	EmailPersonal  string     `gorm:"column:email_personal" json:"email_personal"`
	EmailCorporate string     `gorm:"column:email_corporate" json:"email_corporate"`
	DivisionID     int        `gorm:"column:division_id" json:"division_id"`
	PositionID     int        `gorm:"column:position_id" json:"position_id"`
	StartDate      time.Time  `gorm:"column:start_date" json:"start_date"`
	EndDate        *time.Time `gorm:"column:end_date" json:"end_date"`
	Avatar         string     `gorm:"column:avatar" json:"avatar"`
}

func (e *Employee) TableName() string {
	return "employees"
}

func (e *Employee) RedisKey() string {
	if e.ID == 0 {
		return "employees"
	}

	return "employees:" + strconv.Itoa(e.ID)
}
