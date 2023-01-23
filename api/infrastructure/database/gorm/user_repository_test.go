package gorm

import (
	"testing"

	"github.com/Watson-Sei/face-to-face/domain/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserRepo_List(t *testing.T) {
	// mock db
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	// make repository
	userRepo := NewUserRepository(db)

	// insert test data
	db.Create(&models.User{
		Name: "John",
		Age:  20,
	})

	// run test
	users, err := userRepo.List()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "John", users[0].Name)
	assert.Equal(t, 20, users[0].Age)
}

func TestUserRepo_Get(t *testing.T) {
	// mock db
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	// make repository
	userRepo := NewUserRepository(db)

	// insert test data
	db.Create(&models.User{
		Name: "John",
		Age:  20,
	})

	// run test
	user, err := userRepo.Get("1")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, 20, user.Age)
}

func TestUserRepo_Create(t *testing.T) {
	// mock db
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	// make repository
	userRepo := NewUserRepository(db)

	// run test
	user, err := userRepo.Create(models.User{
		Name: "John",
		Age:  20,
	})

	// assert
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, 20, user.Age)
}

func TestUserRepo_Delete(t *testing.T) {
	// mock db
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	// make repository
	userRepo := NewUserRepository(db)

	// insert test data
	db.Create(&models.User{
		Name: "John",
		Age:  20,
	})

	// run test
	err = userRepo.Delete("1")

	// assert
	assert.NoError(t, err)
}
