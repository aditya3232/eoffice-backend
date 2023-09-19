package auth

import (
	"eoffice-backend/helper"
	"eoffice-backend/library/JWT"
	"eoffice-backend/models/employee"
	"eoffice-backend/models/user"
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(user user.User) (user.User, error)
	Logout(token string) error
}

type service struct {
	userRepository     user.Repository
	employeeRepository employee.Repository
}

func NewService(userRepository user.Repository, employeeRepository employee.Repository) *service {
	return &service{userRepository, employeeRepository}
}

func (s *service) Login(userStruct user.User) (user.User, error) {
	employeeData, err := s.employeeRepository.GetOne(userStruct.EmployeeID)
	if err != nil {
		return user.User{}, err
	}

	users, _, err := s.userRepository.GetAll(map[string]string{
		"employee_id": strconv.Itoa(employeeData.ID)},
		helper.NewPagination(1, 1),
		helper.NewSort("id", "asc"))
	if err != nil {
		return user.User{}, err
	}

	if len(users) == 0 {
		return user.User{}, errors.New("User tersebut tidak ditemukan")
	}

	if users[0].DeletedAt != nil {
		return user.User{}, errors.New("User tersebut sudah dihapus")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(userStruct.Password))
	if err != nil {
		return user.User{}, errors.New("Password yang anda masukkan salah")
	}

	token, err := JWT.GenerateToken(users[0].ID, 30)
	if err != nil {
		return user.User{}, err
	}

	now := time.Now()
	userUpdate, err := s.userRepository.Update(user.User{
		ID:        users[0].ID,
		LastLogin: &now,
		Token:     token,
	})

	if err != nil {
		return user.User{}, err
	}

	return userUpdate, nil
}

func (s *service) Logout(token string) error {
	userID, err := JWT.GetUserIDFromToken(token)
	if err != nil {
		return err
	}

	users, _, err := s.userRepository.GetAll(map[string]string{"id": strconv.Itoa(userID)}, helper.NewPagination(1, 1), helper.NewSort("id", "ASC"))
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return errors.New("User tidak ditemukan")
	}

	if users[0].Token != token {
		return errors.New("Token tidak valid")
	}

	_, err = s.userRepository.Update(
		user.User{
			ID:    userID,
			Token: " ",
		},
	)

	if err != nil {
		return err
	}

	return nil
}
