package rolepermission

type CreateInput struct {
	RoleID       int `json:"role_id" form:"role_id"`
	PermissionID int `json:"permission_id" form:"permission_id"`
}

type UpdateInput struct {
	RoleID       int `json:"role_id" form:"role_id"`
	PermissionID int `json:"permission_id" form:"permission_id"`
}
