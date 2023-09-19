package division

import (
	"eoffice-backend/helper"
)

type DivisionFormatter struct {
	ID        int    `json:"id"`
	ParentID  int    `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Remarks   string `json:"remarks"`
	Nama      string `json:"nama"`
}

func FormatDivision(division Division) DivisionFormatter {
	deletedAt := ""
	parentID := 0

	if division.DeletedAt != nil {
		deletedAt = helper.DateTimeToString(*division.DeletedAt)
	}

	if division.ParentID != nil {
		parentID = *division.ParentID
	}

	formatter := DivisionFormatter{
		ID:        division.ID,
		ParentID:  parentID,
		CreatedAt: helper.DateTimeToString(division.CreatedAt),
		UpdatedAt: helper.DateTimeToString(division.UpdatedAt),
		DeletedAt: deletedAt,
		Nama:      division.Nama,
	}

	return formatter
}

func FormatDivisions(divisions []Division) []DivisionFormatter {
	divisionsFormatter := []DivisionFormatter{}

	for _, division := range divisions {
		divisionFormatter := FormatDivision(division)
		divisionsFormatter = append(divisionsFormatter, divisionFormatter)
	}

	return divisionsFormatter
}
