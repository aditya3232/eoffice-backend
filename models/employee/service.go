package employee

import (
	"eoffice-backend/helper"
	"errors"
	"time"
)

type Service interface {
	Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Employee, helper.Pagination, error)
	GetByID(id int) (Employee, error)
	Create(employee Employee) (Employee, error)
	Update(employee Employee) (Employee, error)
	Delete(id int) error
}

type service struct {
	employeeRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Employee, helper.Pagination, error) {
	employees, pagination, err := s.employeeRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return employees, pagination, nil
}

func (s *service) GetByID(id int) (Employee, error) {
	employee, err := s.employeeRepository.GetOne(id)
	if err != nil {
		return Employee{}, err
	}

	if employee.ID == 0 {
		return Employee{}, nil
	}

	return employee, nil
}

func (s *service) Create(employee Employee) (Employee, error) {
	newEmployee, err := s.employeeRepository.Create(Employee{
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
		Remarks:        employee.Remarks,
		Nama:           employee.Nama,
		Nip:            employee.Nip,
		TempatLahir:    employee.TempatLahir,
		TanggalLahir:   employee.TanggalLahir,
		Alamat:         employee.Alamat,
		NoHp:           employee.NoHp,
		EmailPersonal:  employee.EmailPersonal,
		EmailCorporate: employee.EmailCorporate,
		DivisionID:     employee.DivisionID,
		PositionID:     employee.PositionID,
		StartDate:      employee.StartDate,
		EndDate:        employee.EndDate,
		Avatar:         employee.Avatar,
	})
	if err != nil {
		return Employee{}, err
	}

	return newEmployee, nil
}

func (s *service) Update(employee Employee) (Employee, error) {
	employeeExist, err := s.employeeRepository.GetOne(employee.ID)
	if err != nil {
		return Employee{}, err
	}

	if employeeExist.ID == 0 {
		return Employee{}, errors.New("Employee dengan id tersebut tidak ditemukan")
	}

	employee, err = s.employeeRepository.Update(employee)
	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

// soft delete (change deleted_at value)
func (s *service) Delete(id int) error {
	now := time.Now()
	role, err := s.employeeRepository.Update(
		Employee{
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
