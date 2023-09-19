package user

import (
	"eoffice-backend/helper"
	"eoffice-backend/models/employee"
	"errors"
	"strconv"
	"time"
)

// service yang menentukan repository mana yang akan di call
// sedangkan repository isinya query2

type Service interface {
	Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]User, helper.Pagination, error)
	GetByID(id int) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id int) error
}

type service struct {
	userRepository     Repository
	employeeRepository employee.Repository
}

func NewService(userRepository Repository, employeeRepository employee.Repository) *service {
	return &service{userRepository, employeeRepository}
}

func (s *service) Get(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]User, helper.Pagination, error) {
	user, pagination, err := s.userRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return user, pagination, nil
}

func (s *service) GetByID(id int) (User, error) {
	user, err := s.userRepository.GetOne(id)
	if err != nil {
		return User{}, err
	}

	if user.ID == 0 {
		return User{}, nil
	}

	return user, nil
}

func (s *service) Create(user User) (User, error) {
	userExist, _, err := s.userRepository.GetAll(map[string]string{"employee_id": strconv.Itoa(user.EmployeeID)}, helper.NewPagination(1, 1), helper.NewSort("id", "ASC"))

	if err != nil {
		return User{}, err
	}

	if len(userExist) > 0 {
		return User{}, errors.New("User dengan employee id tersebut sudah ada")
	}

	employee, err := s.employeeRepository.GetOne(user.EmployeeID)

	if err != nil {
		return User{}, err
	}

	if employee.ID == 0 {
		return User{}, errors.New("Employee dengan id tersebut tidak ditemukan")
	}

	newUser, err := s.userRepository.Create(User{
		EmployeeID: user.EmployeeID,
		RoleID:     user.RoleID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
		LastLogin:  nil,
		Remarks:    "",
		Password:   user.Password,
	})
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (s *service) Update(user User) (User, error) {
	isExists, err := s.userRepository.GetOne(user.ID)
	if err != nil {
		return User{}, err
	}

	if isExists.ID == 0 {
		return User{}, errors.New("User dengan id tersebut tidak ditemukan")
	}

	user, err = s.userRepository.Update(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// soft delete (change deleted_at value)
func (s *service) Delete(id int) error {
	now := time.Now()
	user, err := s.userRepository.Update(
		User{
			ID:        id,
			DeletedAt: &now,
		},
	)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User dengan id tersebut tidak ditemukan")
	}

	return nil
}
