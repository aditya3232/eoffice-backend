package profile

import (
	"eoffice-backend/helper"
	"eoffice-backend/models/employee"
	"eoffice-backend/models/permission"
	"eoffice-backend/models/role"
	"eoffice-backend/models/rolepermission"
	"eoffice-backend/models/user"
	"strconv"
	"sync"
)

type Service interface {
	GetProfile(userID int) (Profile, error)
	GetPermission(userID int) ([]permission.Permission, error)
}

type service struct {
	userRepository           user.Repository
	employeeRepository       employee.Repository
	roleRepository           role.Repository
	permissionRepository     permission.Repository
	rolePermissionRepository rolepermission.Repository
}

func NewService(userRepository user.Repository, employeeRepository employee.Repository, roleRepository role.Repository, permissionRepository permission.Repository, rolePermissionRepository rolepermission.Repository) *service {
	return &service{userRepository, employeeRepository, roleRepository, permissionRepository, rolePermissionRepository}
}

func (s *service) GetProfile(userID int) (Profile, error) {
	user, err := s.userRepository.GetOne(userID)
	if err != nil {
		return Profile{}, err
	}

	if user.DeletedAt != nil || user.EmployeeID == 0 {
		return Profile{}, err
	}

	employeeData, err := s.employeeRepository.GetOne(user.EmployeeID)

	if err != nil {
		return Profile{}, err
	}

	return Profile{
		User:     user,
		Employee: employeeData,
	}, nil
}

func (s *service) GetPermission(userID int) ([]permission.Permission, error) {
	user, err := s.userRepository.GetOne(userID)
	if err != nil {
		return nil, err
	}

	if user.DeletedAt != nil || user.EmployeeID == 0 {
		return nil, err
	}

	rolePermissionData, _, err := s.rolePermissionRepository.GetAll(
		map[string]string{
			"role_id": strconv.Itoa(user.RoleID),
		},
		helper.NewPagination(1, 1000),
		helper.NewSort("id", "asc"),
	)

	if err != nil {
		return nil, err
	}

	permissionData := make([]permission.Permission, len(rolePermissionData))
	errChan := make(chan error, len(rolePermissionData))

	var wg sync.WaitGroup
	wg.Add(len(rolePermissionData))

	for i, rolePermission := range rolePermissionData {
		go func(i int, rolePermission rolepermission.RolePermission) {
			defer wg.Done()

			permissionAppend, err := s.permissionRepository.GetOne(rolePermission.PermissionID)

			if err != nil {
				errChan <- err
				return
			}

			permissionData[i] = permissionAppend
		}(i, rolePermission)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return permissionData, nil
}
