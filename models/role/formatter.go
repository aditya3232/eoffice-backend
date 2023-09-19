package role

import (
	"eoffice-backend/helper"
)

type RoleFormatter struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Remarks   string `json:"remarks"`
	Nama      string `json:"nama"`
}

func FormatRole(role Role) RoleFormatter {
	deletedAt := ""

	if role.DeletedAt != nil {
		deletedAt = helper.DateTimeToString(*role.DeletedAt)
	}

	formatter := RoleFormatter{
		ID:        role.ID,
		CreatedAt: helper.DateTimeToString(role.CreatedAt),
		UpdatedAt: helper.DateTimeToString(role.UpdatedAt),
		DeletedAt: deletedAt,
		Nama:      role.Nama,
	}

	return formatter
}

func FormatRoles(roles []Role) []RoleFormatter {
	rolesFormatter := []RoleFormatter{}

	for _, role := range roles {
		roleFormatter := FormatRole(role)
		rolesFormatter = append(rolesFormatter, roleFormatter)
	}

	return rolesFormatter
}
