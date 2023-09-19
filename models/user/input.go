package user

type CreateInput struct {
	EmployeeID int    `json:"employee_id" form:"employee_id" binding:"required" validate:"required"`
	Password   string `json:"password" form:"password" binding:"required" validate:"required"`
	RoleID     int    `json:"role_id" form:"role_id" binding:"required" validate:"required"`
}

type UpdateInput struct {
	Remarks  string `json:"remarks" form:"remarks"`
	Password string `json:"password" form:"password"`
	RoleID   int    `json:"role_id" form:"role_id" binding:"numeric"`
}
