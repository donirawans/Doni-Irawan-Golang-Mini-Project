package repository

import (
	"log"
	"testing"
	"warningfloodsystem/domain/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

// Fungsi untuk membuat database sementara
func NewTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to SQLite in-memory database: %v", err)
	}

	// Auto-migrate the schema
	db.AutoMigrate(&model.User{})
	return db
}

func TestAuthRepository(t *testing.T) {
	// Setup database dan repository
	db := NewTestDB()
	repo := NewAuthRepository(db)

	// Test case: CreateUser
	t.Run("CreateUser", func(t *testing.T) {
		user := &model.User{
			Email:    "test@example.com",
			Password: "hashed_password",
		}

		err := repo.CreateUser(user)
		assert.NoError(t, err, "Expected no error when creating user")

		// Verify user is saved
		var savedUser model.User
		result := db.First(&savedUser, "email = ?", "test@example.com")
		assert.NoError(t, result.Error, "Expected user to be found in database")
		assert.Equal(t, user.Email, savedUser.Email)
	})

	// Test case: FindByEmail
	t.Run("FindByEmail", func(t *testing.T) {
		// Existing user
		user, err := repo.FindByEmail("test@example.com")
		assert.NoError(t, err, "Expected no error when finding user")
		assert.Equal(t, "test@example.com", user.Email)

		// Non-existing user
		_, err = repo.FindByEmail("nonexistent@example.com")
		assert.Error(t, err, "Expected error when finding non-existing user")
		assert.Contains(t, err.Error(), "tidak ditemukan", "Expected error message to contain 'tidak ditemukan'")
	})
}
