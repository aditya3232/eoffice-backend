package database

import (
	"context"
	"eoffice-backend/config"
	"fmt"
	"sync"

	"github.com/olivere/elastic"
	"github.com/redis/go-redis/v9"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	db    *gorm.DB
	redis *redis.Client
	es    *elastic.Client
}

var (
	debug    int = config.ENV.DEBUG
	database Database
	initOnce sync.Once
)

func init() {
	initOnce.Do(func() {
		db, err := ConnectDB()
		if err != nil {
			panic(err)
		}
		redis, err := ConnectRedis()
		if err != nil {
			panic(err)
		}
		es, err := ConnectElastic()
		if err != nil {
			panic(err)
		}
		database = Database{
			db:    db,
			redis: redis,
			es:    es,
		}
	})
}

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.ENV.DB1_HOST, config.ENV.DB1_USERNAME, config.ENV.DB1_PASSWORD, config.ENV.DB1_DATABASE, config.ENV.DB1_PORT)
	logMode := logger.Silent
	if debug == 1 {
		logMode = logger.Info
		fmt.Println("Database connection string: ", dsn)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return db, nil
}

func ConnectRedis() (*redis.Client, error) {
	if config.ENV.REDIS_HOST != "" {
		redis := redis.NewClient(&redis.Options{
			Addr:     config.ENV.REDIS_HOST + ":" + config.ENV.REDIS_PORT,
			Password: config.ENV.REDIS_PASSWORD,
			DB:       0,
		})

		_, err := redis.Ping(context.Background()).Result()
		if err != nil {
			return nil, err
		}

		return redis, nil
	}

	return nil, nil
}

func ConnectElastic() (*elastic.Client, error) {
	// Create a new Elasticsearch client
	client, err := elastic.NewClient(
		elastic.SetURL(config.ENV.ELASTICSEARCH_HOST+":"+config.ENV.ELASTICSEARCH_PORT), // Update with your Elasticsearch URL
		elastic.SetSniff(false), // Disable sniffing of the cluster
	)
	if err != nil {
		return nil, err
	}

	// Ping the Elasticsearch server to check if it's reachable
	info, code, err := client.Ping(config.ENV.ELASTICSEARCH_HOST + ":" + config.ENV.ELASTICSEARCH_PORT).Do(context.Background())
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("Elasticsearch server return code: %d", code)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	return client, nil
}

func DB() *gorm.DB {
	return database.db
}

func Redis() *redis.Client {
	return database.redis
}

func ES() *elastic.Client {
	return database.es
}

func Close() {
	if database.db != nil {
		sqlDB, _ := database.db.DB()
		sqlDB.Close()
		database.db = nil
	}

	if database.redis != nil {
		database.redis.Close()
		database.redis = nil
	}

	if database.es != nil {
		database.es.Stop()
		database.es = nil
	}
}
