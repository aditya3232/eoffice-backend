package position

import (
	"eoffice-backend/helper"
	"errors"
	"time"
)

type Service interface {
	Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Position, helper.Pagination, error)
	GetByID(id int) (Position, error)
	Create(position Position) (Position, error)
	Update(position Position) (Position, error)
	Delete(id int) error
}

type service struct {
	positionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Position, helper.Pagination, error) {
	positions, pagination, err := s.positionRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return positions, pagination, nil
}

func (s *service) GetByID(id int) (Position, error) {
	position, err := s.positionRepository.GetOne(id)
	if err != nil {
		return Position{}, err
	}

	if position.ID == 0 {
		return Position{}, nil
	}

	return position, nil
}

func (s *service) Create(position Position) (Position, error) {
	newPosition, err := s.positionRepository.Create(
		Position{
			DeletedAt: nil,
			Nama:      position.Nama,
		},
	)
	if err != nil {
		return Position{}, err
	}

	return newPosition, nil
}

func (s *service) Update(position Position) (Position, error) {
	positionExist, err := s.positionRepository.GetOne(position.ID)
	if err != nil {
		return Position{}, err
	}

	if positionExist.ID == 0 {
		return Position{}, errors.New("Position dengan id tersebut tidak ditemukan")
	}

	position, err = s.positionRepository.Update(position)
	if err != nil {
		return Position{}, err
	}

	return position, nil
}

// soft delete (change deleted_at value)
func (s *service) Delete(id int) error {
	now := time.Now()
	position, err := s.positionRepository.Update(
		Position{
			ID:        id,
			DeletedAt: &now,
		},
	)
	if err != nil {
		return err
	}

	if position.ID == 0 {
		return errors.New("Position dengan id tersebut tidak ditemukan")
	}

	return nil
}
