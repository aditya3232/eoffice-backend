package position

import (
	"eoffice-backend/helper"
)

type PositionFormatter struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Remarks   string `json:"remarks"`
	Nama      string `json:"nama"`
}

func FormatPosition(position Position) PositionFormatter {
	deletedAt := ""

	if position.DeletedAt != nil {
		deletedAt = helper.DateTimeToString(*position.DeletedAt)
	}

	formatter := PositionFormatter{
		ID:        position.ID,
		CreatedAt: helper.DateTimeToString(position.CreatedAt),
		UpdatedAt: helper.DateTimeToString(position.UpdatedAt),
		DeletedAt: deletedAt,
		Nama:      position.Nama,
	}

	return formatter
}

func FormatPositions(positions []Position) []PositionFormatter {
	positionsFormatter := []PositionFormatter{}

	for _, position := range positions {
		positionFormatter := FormatPosition(position)
		positionsFormatter = append(positionsFormatter, positionFormatter)
	}

	return positionsFormatter
}
