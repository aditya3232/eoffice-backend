package rolepermission

import (
	"eoffice-backend/helper"
	"errors"
	"time"
)

type Service interface {
	Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]RolePermission, helper.Pagination, error)
	GetByID(id int) (RolePermission, error)
	Create(rolePermission RolePermission) (RolePermission, error)
	Update(rolePermission RolePermission) (RolePermission, error)
	Delete(id int) error
}

type service struct {
	rolePermissionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]RolePermission, helper.Pagination, error) {
	rolePermissions, pagination, err := s.rolePermissionRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return rolePermissions, pagination, nil
}

func (s *service) GetByID(id int) (RolePermission, error) {
	rolePermission, err := s.rolePermissionRepository.GetOne(id)
	if err != nil {
		return RolePermission{}, err
	}

	if rolePermission.ID == 0 {
		return RolePermission{}, nil
	}

	return rolePermission, nil
}

func (s *service) Create(rolePermission RolePermission) (RolePermission, error) {
	newRolePermission, err := s.rolePermissionRepository.Create(
		RolePermission{
			DeletedAt:    nil,
			RoleID:       rolePermission.RoleID,
			PermissionID: rolePermission.PermissionID,
		},
	)
	if err != nil {
		return RolePermission{}, err
	}

	return newRolePermission, nil
}

func (s *service) Update(rolePermission RolePermission) (RolePermission, error) {
	rolePermissionExist, err := s.rolePermissionRepository.GetOne(rolePermission.ID)
	if err != nil {
		return RolePermission{}, err
	}

	if rolePermissionExist.ID == 0 {
		return RolePermission{}, errors.New("RolePermission dengan id tersebut tidak ditemukan")
	}

	rolePermission, err = s.rolePermissionRepository.Update(rolePermission)
	if err != nil {
		return RolePermission{}, err
	}

	return rolePermission, nil
}

// soft delete (change deleted_at value)
func (s *service) Delete(id int) error {
	now := time.Now()
	rolePermission, err := s.rolePermissionRepository.Update(
		RolePermission{
			ID:        id,
			DeletedAt: &now,
		},
	)
	if err != nil {
		return err
	}

	if rolePermission.ID == 0 {
		return errors.New("RolePermission dengan id tersebut tidak ditemukan")
	}

	return nil
}
