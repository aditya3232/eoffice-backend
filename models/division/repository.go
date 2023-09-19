package division

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
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Division, helper.Pagination, error)
	GetOne(id int) (Division, error)
	Create(division Division) (Division, error)
	Update(division Division) (Division, error)
	Delete(ID int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) *repository {
	return &repository{db, redis}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Division, helper.Pagination, error) {
	var divisions []Division
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&divisions), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return divisions, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&divisions).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(divisions)

	return divisions, pagination, nil
}

func (r *repository) GetOne(id int) (Division, error) {
	// get to redis first to check if the data is cached
	division := Division{
		ID: id,
	}

	data, err := r.redis.Get(context.Background(), division.RedisKey()).Bytes()
	if err == nil {
		// Data is cached in Redis, unmarshal it into the struct
		err = json.Unmarshal(data, &division)
		if err != nil {
			return Division{}, err
		}
	} else {
		// If the data is not cached, get it from the database
		err = r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&division).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return Division{}, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			division = Division{}
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(division)
		if err != nil {
			return Division{}, err
		}

		// Cache the data to Redis
		err = r.redis.Set(context.Background(), division.RedisKey(), jsonData, 3*time.Minute).Err()
		if err != nil {
			return Division{}, err
		}
	}

	return division, nil
}

func (r *repository) Create(division Division) (Division, error) {
	err := r.db.Model(&division).Create(&division).Error
	if err != nil {
		return Division{}, err
	}

	return division, nil
}

func (r *repository) Update(division Division) (Division, error) {
	division.UpdatedAt = time.Now()
	err := r.db.Model(&division).Where("id = ?", division.ID).Updates(&division).Error
	if err != nil {
		return Division{}, err
	}

	// Reload the updated row from the database
	err = r.db.Where("id = ?", division.ID).First(&division).Error
	if err != nil {
		return Division{}, err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), division.RedisKey()).Err()
	if err != nil {
		return Division{}, err
	}

	return division, nil
}

func (r *repository) Delete(id int) error {
	division := Division{
		ID: id,
	}

	err := r.db.Table(division.TableName()).Delete(division).Error
	if err != nil {
		return err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), division.RedisKey()).Err()
	if err != nil {
		return err
	}

	return nil
}
