package employee

import (
	"eoffice-backend/helper"
)

type EmployeeFormatter struct {
	ID             int    `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
	Remarks        string `json:"remarks"`
	Nama           string `json:"nama"`
	Nip            int    `json:"nip"`
	TempatLahir    string `json:"tempat_lahir"`
	TanggalLahir   string `json:"tanggal_lahir"`
	Alamat         string `json:"alamat"`
	NoHp           string `json:"no_hp"`
	EmailPersonal  string `json:"email_personal"`
	EmailCorporate string `json:"email_corporate"`
	DivisionID     int    `json:"division_id"`
	PositionID     int    `json:"position_id"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Avatar         string `json:"avatar"`
}

func FormatEmployee(employee Employee) EmployeeFormatter {
	deletedAt := ""
	endDate := ""

	if employee.DeletedAt != nil {
		deletedAt = helper.DateTimeToString(*employee.DeletedAt)
	}

	if employee.EndDate != nil {
		endDate = helper.DateTimeToString(*employee.EndDate)
	}

	formatter := EmployeeFormatter{
		ID:             employee.ID,
		CreatedAt:      helper.DateTimeToString(employee.CreatedAt),
		UpdatedAt:      helper.DateTimeToString(employee.UpdatedAt),
		DeletedAt:      deletedAt,
		Nama:           employee.Nama,
		Nip:            employee.Nip,
		TempatLahir:    employee.TempatLahir,
		TanggalLahir:   helper.DateTimeToString(employee.TanggalLahir),
		Alamat:         employee.Alamat,
		NoHp:           employee.NoHp,
		EmailPersonal:  employee.EmailPersonal,
		EmailCorporate: employee.EmailCorporate,
		DivisionID:     employee.DivisionID,
		PositionID:     employee.PositionID,
		StartDate:      helper.DateTimeToString(employee.StartDate),
		EndDate:        endDate,
		Avatar:         employee.Avatar,
	}

	return formatter
}

func FormatEmployees(employees []Employee) []EmployeeFormatter {
	employeesFormatter := []EmployeeFormatter{}

	for _, employee := range employees {
		employeeFormatter := FormatEmployee(employee)
		employeesFormatter = append(employeesFormatter, employeeFormatter)
	}

	return employeesFormatter
}
