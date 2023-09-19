package permission

import (
	"eoffice-backend/helper"
	"errors"
	"time"
)

type Service interface {
	Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Permission, helper.Pagination, error)
	GetByID(id int) (Permission, error)
	Create(permission Permission) (Permission, error)
	Update(permission Permission) (Permission, error)
	Delete(id int) error
}

type service struct {
	permissionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Permission, helper.Pagination, error) {
	permissions, pagination, err := s.permissionRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return permissions, pagination, nil
}

func (s *service) GetByID(id int) (Permission, error) {
	permission, err := s.permissionRepository.GetOne(id)
	if err != nil {
		return Permission{}, err
	}

	if permission.ID == 0 {
		return Permission{}, nil
	}

	return permission, nil
}

func (s *service) Create(permission Permission) (Permission, error) {
	newPermission, err := s.permissionRepository.Create(
		Permission{
			DeletedAt: nil,
			Nama:      permission.Nama,
			ParentID:  permission.ParentID,
			Url:       permission.Url,
			Position:  permission.Position,
		},
	)
	if err != nil {
		return Permission{}, err
	}

	return newPermission, nil
}

func (s *service) Update(permission Permission) (Permission, error) {
	permissionExist, err := s.permissionRepository.GetOne(permission.ID)
	if err != nil {
		return Permission{}, err
	}

	if permissionExist.ID == 0 {
		return Permission{}, errors.New("Permission dengan id tersebut tidak ditemukan")
	}

	permission, err = s.permissionRepository.Update(permission)
	if err != nil {
		return Permission{}, err
	}

	return permission, nil
}

// soft delete (change deleted_at value)
func (s *service) Delete(id int) error {
	now := time.Now()
	permission, err := s.permissionRepository.Update(
		Permission{
			ID:        id,
			DeletedAt: &now,
		},
	)
	if err != nil {
		return err
	}

	if permission.ID == 0 {
		return errors.New("Permission dengan id tersebut tidak ditemukan")
	}

	return nil
}
