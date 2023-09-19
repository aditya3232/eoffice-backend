package role

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
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Role, helper.Pagination, error)
	GetOne(id int) (Role, error)
	Create(role Role) (Role, error)
	Update(role Role) (Role, error)
	Delete(ID int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) *repository {
	return &repository{db, redis}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Role, helper.Pagination, error) {
	var roles []Role
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&roles), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return roles, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&roles).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(roles)

	return roles, pagination, nil
}

func (r *repository) GetOne(id int) (Role, error) {
	// get to redis first to check if the data is cached
	role := Role{
		ID: id,
	}

	data, err := r.redis.Get(context.Background(), role.RedisKey()).Bytes()
	if err == nil {
		// Data is cached in Redis, unmarshal it into the struct
		err = json.Unmarshal(data, &role)
		if err != nil {
			return Role{}, err
		}
	} else {
		// If the data is not cached, get it from the database
		err = r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&role).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return Role{}, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			role = Role{}
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(role)
		if err != nil {
			return Role{}, err
		}

		// Cache the data to Redis
		err = r.redis.Set(context.Background(), role.RedisKey(), jsonData, 3*time.Minute).Err()
		if err != nil {
			return Role{}, err
		}
	}

	return role, nil
}

func (r *repository) Create(role Role) (Role, error) {
	err := r.db.Model(&role).Create(&role).Error
	if err != nil {
		return Role{}, err
	}

	return role, nil
}

func (r *repository) Update(role Role) (Role, error) {
	role.UpdatedAt = time.Now()
	err := r.db.Model(&role).Where("id = ?", role.ID).Updates(&role).Error
	if err != nil {
		return Role{}, err
	}

	// Reload the updated row from the database
	err = r.db.Where("id = ?", role.ID).First(&role).Error
	if err != nil {
		return Role{}, err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), role.RedisKey()).Err()
	if err != nil {
		return Role{}, err
	}

	return role, nil
}

func (r *repository) Delete(id int) error {
	role := Role{
		ID: id,
	}

	err := r.db.Table(role.TableName()).Delete(role).Error
	if err != nil {
		return err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), role.RedisKey()).Err()
	if err != nil {
		return err
	}

	return nil
}
