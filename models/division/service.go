package division

import (
	"eoffice-backend/helper"
	"errors"
	"time"
)

type Service interface {
	Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Division, helper.Pagination, error)
	GetByID(id int) (Division, error)
	Create(division Division) (Division, error)
	Update(division Division) (Division, error)
	Delete(id int) error
}

type service struct {
	divisionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Division, helper.Pagination, error) {
	divisions, pagination, err := s.divisionRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return divisions, pagination, nil
}

func (s *service) GetByID(id int) (Division, error) {
	division, err := s.divisionRepository.GetOne(id)
	if err != nil {
		return Division{}, err
	}

	if division.ID == 0 {
		return Division{}, nil
	}

	return division, nil
}

func (s *service) Create(division Division) (Division, error) {
	newDivision, err := s.divisionRepository.Create(
		Division{
			DeletedAt: nil,
			Nama:      division.Nama,
			ParentID:  division.ParentID,
		},
	)
	if err != nil {
		return Division{}, err
	}

	return newDivision, nil
}

func (s *service) Update(division Division) (Division, error) {
	divisionExist, err := s.divisionRepository.GetOne(division.ID)
	if err != nil {
		return Division{}, err
	}

	if divisionExist.ID == 0 {
		return Division{}, errors.New("Division dengan id tersebut tidak ditemukan")
	}

	division, err = s.divisionRepository.Update(division)
	if err != nil {
		return Division{}, err
	}

	return division, nil
}

// soft delete (change deleted_at value)
func (s *service) Delete(id int) error {
	now := time.Now()
	division, err := s.divisionRepository.Update(
		Division{
			ID:        id,
			DeletedAt: &now,
		},
	)
	if err != nil {
		return err
	}

	if division.ID == 0 {
		return errors.New("Division dengan id tersebut tidak ditemukan")
	}

	return nil
}
