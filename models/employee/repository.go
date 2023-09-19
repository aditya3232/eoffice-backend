package employee

import (
	"context"
	"encoding/json"
	"eoffice-backend/helper"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Employee, helper.Pagination, error)
	GetOne(id int) (Employee, error)
	Create(employee Employee) (Employee, error)
	Update(employee Employee) (Employee, error)
	Delete(ID int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) *repository {
	return &repository{db, redis}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Employee, helper.Pagination, error) {
	var employees []Employee
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&employees), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return employees, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&employees).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(employees)

	return employees, pagination, nil
}

func (r *repository) GetOne(id int) (Employee, error) {
	// get to redis first to check if the data is cached
	employee := Employee{
		ID: id,
	}

	data, err := r.redis.Get(context.Background(), employee.RedisKey()).Bytes()
	if err == nil {
		// Data is cached in Redis, unmarshal it into the struct
		err = json.Unmarshal(data, &employee)
		if err != nil {
			return Employee{}, err
		}
	} else {
		// If the data is not cached, get it from the database
		err = r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&employee).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return Employee{}, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			employee = Employee{}
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(employee)
		if err != nil {
			return Employee{}, err
		}

		// Cache the data to Redis
		err = r.redis.Set(context.Background(), employee.RedisKey(), jsonData, 3*time.Minute).Err()
		if err != nil {
			return Employee{}, err
		}
	}

	return employee, nil
}

func (r *repository) Create(employee Employee) (Employee, error) {
	err := r.db.Model(&employee).Create(&employee).Error
	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (r *repository) Update(employee Employee) (Employee, error) {
	employee.UpdatedAt = time.Now()
	err := r.db.Model(&employee).Where("id = ?", employee.ID).Updates(&employee).Error
	if err != nil {
		return Employee{}, err
	}

	// Reload the updated row from the database
	err = r.db.Where("id = ?", employee.ID).First(&employee).Error
	if err != nil {
		return Employee{}, err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), employee.RedisKey()).Err()
	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (r *repository) Delete(id int) error {
	employee := Employee{
		ID: id,
	}

	err := r.db.Table(employee.TableName()).Delete(employee).Error
	if err != nil {
		return err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), employee.RedisKey()).Err()
	if err != nil {
		return err
	}

	return nil
}
