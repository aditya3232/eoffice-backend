package rolepermission

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
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]RolePermission, helper.Pagination, error)
	GetOne(id int) (RolePermission, error)
	Create(rolePermission RolePermission) (RolePermission, error)
	Update(rolePermission RolePermission) (RolePermission, error)
	Delete(ID int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) *repository {
	return &repository{db, redis}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]RolePermission, helper.Pagination, error) {
	var rolePermissions []RolePermission
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&rolePermissions), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return rolePermissions, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&rolePermissions).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(rolePermissions)

	return rolePermissions, pagination, nil
}

func (r *repository) GetOne(id int) (RolePermission, error) {
	// get to redis first to check if the data is cached
	rolePermission := RolePermission{
		ID: id,
	}

	data, err := r.redis.Get(context.Background(), rolePermission.RedisKey()).Bytes()
	if err == nil {
		// Data is cached in Redis, unmarshal it into the struct
		err = json.Unmarshal(data, &rolePermission)
		if err != nil {
			return RolePermission{}, err
		}
	} else {
		// If the data is not cached, get it from the database
		err = r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&rolePermission).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return RolePermission{}, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			rolePermission = RolePermission{}
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(rolePermission)
		if err != nil {
			return RolePermission{}, err
		}

		// Cache the data to Redis
		err = r.redis.Set(context.Background(), rolePermission.RedisKey(), jsonData, 3*time.Minute).Err()
		if err != nil {
			return RolePermission{}, err
		}
	}

	return rolePermission, nil
}

func (r *repository) Create(rolePermission RolePermission) (RolePermission, error) {
	err := r.db.Model(&rolePermission).Create(&rolePermission).Error
	if err != nil {
		return RolePermission{}, err
	}

	return rolePermission, nil
}

func (r *repository) Update(rolePermission RolePermission) (RolePermission, error) {
	rolePermission.UpdatedAt = time.Now()
	err := r.db.Model(&rolePermission).Where("id = ?", rolePermission.ID).Updates(&rolePermission).Error
	if err != nil {
		return RolePermission{}, err
	}

	// Reload the updated row from the database
	err = r.db.Where("id = ?", rolePermission.ID).First(&rolePermission).Error
	if err != nil {
		return RolePermission{}, err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), rolePermission.RedisKey()).Err()
	if err != nil {
		return RolePermission{}, err
	}

	return rolePermission, nil
}

func (r *repository) Delete(id int) error {
	rolePermission := RolePermission{
		ID: id,
	}

	err := r.db.Table(rolePermission.TableName()).Delete(rolePermission).Error
	if err != nil {
		return err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), rolePermission.RedisKey()).Err()
	if err != nil {
		return err
	}

	return nil
}
