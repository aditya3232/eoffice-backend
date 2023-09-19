package location

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
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Location, helper.Pagination, error)
	GetOne(id int) (Location, error)
	Create(location Location) (Location, error)
	Update(location Location) (Location, error)
	Delete(ID int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) *repository {
	return &repository{db, redis}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Location, helper.Pagination, error) {
	var locations []Location
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&locations), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return locations, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&locations).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(locations)

	return locations, pagination, nil
}

func (r *repository) GetOne(id int) (Location, error) {
	// get to redis first to check if the data is cached
	location := Location{
		ID: id,
	}

	data, err := r.redis.Get(context.Background(), location.RedisKey()).Bytes()
	if err == nil {
		// Data is cached in Redis, unmarshal it into the struct
		err = json.Unmarshal(data, &location)
		if err != nil {
			return Location{}, err
		}
	} else {
		// If the data is not cached, get it from the database
		err = r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&location).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return Location{}, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			location = Location{}
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(location)
		if err != nil {
			return Location{}, err
		}

		// Cache the data to Redis
		err = r.redis.Set(context.Background(), location.RedisKey(), jsonData, 3*time.Minute).Err()
		if err != nil {
			return Location{}, err
		}
	}

	return location, nil
}

func (r *repository) Create(location Location) (Location, error) {
	err := r.db.Model(&location).Create(&location).Error
	if err != nil {
		return Location{}, err
	}

	return location, nil
}

func (r *repository) Update(location Location) (Location, error) {
	location.UpdatedAt = time.Now()
	err := r.db.Model(&location).Where("id = ?", location.ID).Updates(&location).Error
	if err != nil {
		return Location{}, err
	}

	// Reload the updated row from the database
	err = r.db.Where("id = ?", location.ID).First(&location).Error
	if err != nil {
		return Location{}, err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), location.RedisKey()).Err()
	if err != nil {
		return Location{}, err
	}

	return location, nil
}

func (r *repository) Delete(id int) error {
	location := Location{
		ID: id,
	}

	err := r.db.Table(location.TableName()).Delete(location).Error
	if err != nil {
		return err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), location.RedisKey()).Err()
	if err != nil {
		return err
	}

	return nil
}
