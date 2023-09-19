package permission

import (
	"eoffice-backend/helper"
)

type PermissionFormatter struct {
	ID        int    `json:"id"`
	ParentID  int    `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Remarks   string `json:"remarks"`
	Nama      string `json:"nama"`
	Url       string `json:"url"`
	Position  int    `json:"position"`
}

func FormatPermission(permission Permission) PermissionFormatter {
	deletedAt := ""
	parentID := 0

	if permission.DeletedAt != nil {
		deletedAt = helper.DateTimeToString(*permission.DeletedAt)
	}

	if permission.ParentID != nil {
		parentID = *permission.ParentID
	}

	formatter := PermissionFormatter{
		ID:        permission.ID,
		ParentID:  parentID,
		CreatedAt: helper.DateTimeToString(permission.CreatedAt),
		UpdatedAt: helper.DateTimeToString(permission.UpdatedAt),
		DeletedAt: deletedAt,
		Nama:      permission.Nama,
		Url:       permission.Url,
		Position:  permission.Position,
	}

	return formatter
}

func FormatPermissions(permissions []Permission) []PermissionFormatter {
	permissionsFormatter := []PermissionFormatter{}

	for _, permission := range permissions {
		permissionFormatter := FormatPermission(permission)
		permissionsFormatter = append(permissionsFormatter, permissionFormatter)
	}

	return permissionsFormatter
}
