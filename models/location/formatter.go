package location

import (
	"eoffice-backend/helper"
)

type LocationFormatter struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Remarks   string `json:"remarks"`
	Nama      string `json:"nama"`
}

func FormatLocation(location Location) LocationFormatter {
	deletedAt := ""

	if location.DeletedAt != nil {
		deletedAt = helper.DateTimeToString(*location.DeletedAt)
	}

	formatter := LocationFormatter{
		ID:        location.ID,
		CreatedAt: helper.DateTimeToString(location.CreatedAt),
		UpdatedAt: helper.DateTimeToString(location.UpdatedAt),
		DeletedAt: deletedAt,
		Nama:      location.Nama,
	}

	return formatter
}

func FormatLocations(locations []Location) []LocationFormatter {
	locationsFormatter := []LocationFormatter{}

	for _, location := range locations {
		locationFormatter := FormatLocation(location)
		locationsFormatter = append(locationsFormatter, locationFormatter)
	}

	return locationsFormatter
}
