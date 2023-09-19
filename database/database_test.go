package database_test

import (
	"eoffice-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	// Test
	db, err := database.ConnectDB()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Cleanup
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func TestDB(t *testing.T) {
	// Test
	db := database.DB()
	assert.NotNil(t, db)
}

func TestClose(t *testing.T) {
	// Test
	database.Close()

	// Assert
	assert.Nil(t, database.DB())
}
