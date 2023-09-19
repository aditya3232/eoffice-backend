package location

import (
	"eoffice-backend/helper"
	"errors"
	"time"
)

type Service interface {
	Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Location, helper.Pagination, error)
	GetByID(id int) (Location, error)
	Create(location Location) (Location, error)
	Update(location Location) (Location, error)
	Delete(id int) error
}

type service struct {
	locationRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Location, helper.Pagination, error) {
	locations, pagination, err := s.locationRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return locations, pagination, nil
}

func (s *service) GetByID(id int) (Location, error) {
	location, err := s.locationRepository.GetOne(id)
	if err != nil {
		return Location{}, err
	}

	if location.ID == 0 {
		return Location{}, nil
	}

	return location, nil
}

func (s *service) Create(location Location) (Location, error) {
	newLocation, err := s.locationRepository.Create(
		Location{
			DeletedAt: nil,
			Nama:      location.Nama,
		},
	)
	if err != nil {
		return Location{}, err
	}

	return newLocation, nil
}

func (s *service) Update(location Location) (Location, error) {
	locationExist, err := s.locationRepository.GetOne(location.ID)
	if err != nil {
		return Location{}, err
	}

	if locationExist.ID == 0 {
		return Location{}, errors.New("Location dengan id tersebut tidak ditemukan")
	}

	location, err = s.locationRepository.Update(location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}

// soft delete (change deleted_at value)
func (s *service) Delete(id int) error {
	now := time.Now()
	location, err := s.locationRepository.Update(
		Location{
			ID:        id,
			DeletedAt: &now,
		},
	)
	if err != nil {
		return err
	}

	if location.ID == 0 {
		return errors.New("Location dengan id tersebut tidak ditemukan")
	}

	return nil
}
