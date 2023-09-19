package user

import (
	"context"
	"encoding/json"
	"eoffice-backend/helper"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]User, helper.Pagination, error)
	GetOne(id int) (User, error)
	Create(User) (User, error)
	Update(User) (User, error)
	Delete(id int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) *repository {
	return &repository{db, redis}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]User, helper.Pagination, error) {
	var users []User
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&users), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return users, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&users).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(users)

	return users, pagination, nil
}

func (r *repository) GetOne(id int) (User, error) {
	// get to redis first to check if the data is cached
	user := User{
		ID: id,
	}

	data, err := r.redis.Get(context.Background(), user.RedisKey()).Bytes()
	if err == nil {
		// Data is cached in Redis, unmarshal it into the struct
		err = json.Unmarshal(data, &user)
		if err != nil {
			return User{}, err
		}
	} else {
		// If the data is not cached, get it from the database
		err = r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = User{}
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(user)
		if err != nil {
			return User{}, err
		}

		// Cache the data to Redis
		err = r.redis.Set(context.Background(), user.RedisKey(), jsonData, 3*time.Minute).Err()
		if err != nil {
			return User{}, err
		}
	}

	return user, nil
}

func (r *repository) Create(user User) (User, error) {
	if user.Password == "" {
		return User{}, errors.New("password is required")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	user.Password = string(password)

	err = r.db.Model(&user).Create(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	if user.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			return User{}, err
		}

		user.Password = string(password)
	}

	user.UpdatedAt = time.Now()
	err := r.db.Model(&user).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return User{}, err
	}

	// Reload the updated row from the database
	err = r.db.Where("id = ?", user.ID).First(&user).Error
	if err != nil {
		return User{}, err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), user.RedisKey()).Err()
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *repository) Delete(id int) error {
	user := User{
		ID: id,
	}

	err := r.db.Table(user.TableName()).Delete(user).Error
	if err != nil {
		return err
	}

	// Delete the cached data from Redis
	err = r.redis.Del(context.Background(), user.RedisKey()).Err()
	if err != nil {
		return err
	}

	return nil
}
