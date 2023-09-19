package profile

import (
	"eoffice-backend/models/employee"
	"eoffice-backend/models/user"
)

type Profile struct {
	User     user.User         `json:"user"`
	Employee employee.Employee `json:"employee"`
}
