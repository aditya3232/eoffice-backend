package role

import (
	"eoffice-backend/helper"
	"errors"
	"time"
)

type Service interface {
	Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Role, helper.Pagination, error)
	GetByID(id int) (Role, error)
	Create(role Role) (Role, error)
	Update(role Role) (Role, error)
	Delete(id int) error
}

type service struct {
	roleRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Role, helper.Pagination, error) {
	roles, pagination, err := s.roleRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return roles, pagination, nil
}

func (s *service) GetByID(id int) (Role, error) {
	role, err := s.roleRepository.GetOne(id)
	if err != nil {
		return Role{}, err
	}

	if role.ID == 0 {
		return Role{}, nil
	}

	return role, nil
}

func (s *service) Create(role Role) (Role, error) {
	newRole, err := s.roleRepository.Create(
		Role{
			DeletedAt: nil,
			Nama:      role.Nama,
		},
	)
	if err != nil {
		return Role{}, err
	}

	return newRole, nil
}

func (s *service) Update(role Role) (Role, error) {
	roleExist, err := s.roleRepository.GetOne(role.ID)
	if err != nil {
		return Role{}, err
	}

	if roleExist.ID == 0 {
		return Role{}, errors.New("Role dengan id tersebut tidak ditemukan")
	}

	role, err = s.roleRepository.Update(role)
	if err != nil {
		return Role{}, err
	}

	return role, nil
}

// soft delete (change deleted_at value)
func (s *service) Delete(id int) error {
	now := time.Now()
	role, err := s.roleRepository.Update(
		Role{
			ID:        id,
			DeletedAt: &now,
		},
	)
	if err != nil {
		return err
	}

	if role.ID == 0 {
		return errors.New("Role dengan id tersebut tidak ditemukan")
	}

	return nil
}
