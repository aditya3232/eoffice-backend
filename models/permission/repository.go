package permission

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
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Permission, helper.Pagination, error)
	GetOne(id int) (Permission, error)
	Create(permission Permission) (Permission, error)
	Update(permission Permission) (Permission, error)
	Delete(ID int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) *repository {
	return &repository{db, redis}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Permission, helper.Pagination, error) {
	var permissions []Permission
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&permissions), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return permissions, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&permissions).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(permissions)

	return permissions, pagination, nil
}

func (r *repository) GetOne(id int) (Permission, error) {
	// get to redis first to check if the data is cached
	permission := Permission{
		ID: id,
	}

	data, err := r.redis.Get(context.Background(), permission.RedisKey()).Bytes()
	if err == nil {
		// Data is cached in Redis, unmarshal it into the struct
		err = json.Unmarshal(data, &permission)
		if err != nil {
			return Permission{}, err
		}
	} else {
		// If the data is not cached, get it from the database
		err = r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&permission).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return Permission{}, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			permission = Permission{}
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(permission)
		if err != nil {
			return Permission{}, err
		}

		// Cache the data to Redis
		err = r.redis.Set(context.Background(), permission.RedisKey(), jsonData, 3*time.Minute).Err()
		if err != nil {
			return Permission{}, err
		}
	}

	return permission, nil
}

func (r *repository) Create(permission Permission) (Permission, error) {
	err := r.db.Model(&permission).Create(&permission).Error
	if err != nil {
		return Permission{}, err
	}

	return permission, nil
}

func (r *repository) Update(permission Permission) (Permission, error) {
	permission.UpdatedAt = time.Now()
	err := r.db.Model(&permission).Where("id = ?", permission.ID).Updates(&permission).Error
	if err != nil {
		return Permission{}, err
	}

	// Reload the updated row from the database
	err = r.db.Where("id = ?", permission.ID).First(&permission).Error
	if err != nil {
		return Permission{}, err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), permission.RedisKey()).Err()
	if err != nil {
		return Permission{}, err
	}

	return permission, nil
}

func (r *repository) Delete(id int) error {
	permission := Permission{
		ID: id,
	}

	err := r.db.Table(permission.TableName()).Delete(permission).Error
	if err != nil {
		return err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), permission.RedisKey()).Err()
	if err != nil {
		return err
	}

	return nil
}
