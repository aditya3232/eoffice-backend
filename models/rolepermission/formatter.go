package rolepermission

import (
	"eoffice-backend/helper"
)

type RolePermissionFormatter struct {
	ID           int    `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
	Remarks      string `json:"remarks"`
	RoleID       int    `json:"role_id"`
	PermissionID int    `json:"permission_id"`
}

func FormatRolePermission(rolePermission RolePermission) RolePermissionFormatter {
	deletedAt := ""

	if rolePermission.DeletedAt != nil {
		deletedAt = helper.DateTimeToString(*rolePermission.DeletedAt)
	}

	formatter := RolePermissionFormatter{
		ID:           rolePermission.ID,
		CreatedAt:    helper.DateTimeToString(rolePermission.CreatedAt),
		UpdatedAt:    helper.DateTimeToString(rolePermission.UpdatedAt),
		DeletedAt:    deletedAt,
		Remarks:      rolePermission.Remarks,
		RoleID:       rolePermission.RoleID,
		PermissionID: rolePermission.PermissionID,
	}

	return formatter
}

func FormatRolePermissions(rolePermissions []RolePermission) []RolePermissionFormatter {
	rolePermissionsFormatter := []RolePermissionFormatter{}

	for _, rolePermission := range rolePermissions {
		rolePermissionFormatter := FormatRolePermission(rolePermission)
		rolePermissionsFormatter = append(rolePermissionsFormatter, rolePermissionFormatter)
	}

	return rolePermissionsFormatter
}
