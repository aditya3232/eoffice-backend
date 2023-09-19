package position

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
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Position, helper.Pagination, error)
	GetOne(id int) (Position, error)
	Create(position Position) (Position, error)
	Update(position Position) (Position, error)
	Delete(ID int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) *repository {
	return &repository{db, redis}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Position, helper.Pagination, error) {
	var positions []Position
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&positions), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return positions, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&positions).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(positions)

	return positions, pagination, nil
}

func (r *repository) GetOne(id int) (Position, error) {
	// get to redis first to check if the data is cached
	position := Position{
		ID: id,
	}

	data, err := r.redis.Get(context.Background(), position.RedisKey()).Bytes()
	if err == nil {
		// Data is cached in Redis, unmarshal it into the struct
		err = json.Unmarshal(data, &position)
		if err != nil {
			return Position{}, err
		}
	} else {
		// If the data is not cached, get it from the database
		err = r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&position).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return Position{}, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			position = Position{}
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(position)
		if err != nil {
			return Position{}, err
		}

		// Cache the data to Redis
		err = r.redis.Set(context.Background(), position.RedisKey(), jsonData, 3*time.Minute).Err()
		if err != nil {
			return Position{}, err
		}
	}

	return position, nil
}

func (r *repository) Create(position Position) (Position, error) {
	err := r.db.Model(&position).Create(&position).Error
	if err != nil {
		return Position{}, err
	}

	return position, nil
}

func (r *repository) Update(position Position) (Position, error) {
	position.UpdatedAt = time.Now()
	err := r.db.Model(&position).Where("id = ?", position.ID).Updates(&position).Error
	if err != nil {
		return Position{}, err
	}

	// Reload the updated row from the database
	err = r.db.Where("id = ?", position.ID).First(&position).Error
	if err != nil {
		return Position{}, err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), position.RedisKey()).Err()
	if err != nil {
		return Position{}, err
	}

	return position, nil
}

func (r *repository) Delete(id int) error {
	position := Position{
		ID: id,
	}

	err := r.db.Table(position.TableName()).Delete(position).Error
	if err != nil {
		return err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), position.RedisKey()).Err()
	if err != nil {
		return err
	}

	return nil
}
