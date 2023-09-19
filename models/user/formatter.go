package user

import (
	"eoffice-backend/helper"
)

type UserFormatter struct {
	ID         int    `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
	Remarks    string `json:"remarks"`
	EmployeeID int    `json:"employee_id"`
	Password   string `json:"password"`
	Token      string `json:"token"`
	RoleID     int    `json:"role_id"`
	LastLogin  string `json:"last_login"`
}

func FormatUser(user User) UserFormatter {
	deletedAt := ""
	lastLogin := ""

	if user.LastLogin != nil {
		lastLogin = helper.DateTimeToString(*user.LastLogin)
	}

	if user.DeletedAt != nil {
		deletedAt = helper.DateTimeToString(*user.DeletedAt)
	}

	formatter := UserFormatter{
		ID:         user.ID,
		CreatedAt:  helper.DateTimeToString(user.CreatedAt),
		UpdatedAt:  helper.DateTimeToString(user.UpdatedAt),
		DeletedAt:  deletedAt,
		Remarks:    user.Remarks,
		EmployeeID: user.EmployeeID,
		Password:   user.Password,
		Token:      user.Token,
		RoleID:     user.RoleID,
		LastLogin:  lastLogin,
	}

	return formatter
}

func FormatUsers(users []User) []UserFormatter {
	usersFormatter := []UserFormatter{}

	for _, user := range users {
		userFormatter := FormatUser(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
