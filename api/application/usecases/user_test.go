package usecases

import (
	"testing"

	"github.com/Watson-Sei/face-to-face/domain/models"
	"github.com/Watson-Sei/face-to-face/domain/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetUsersUsecase(t *testing.T) {
	// mock db
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	// make repository
	userRepo := repositories.NewUserRepository(db)
	// make usecase
	userUsecase := NewUserUsecase(userRepo)

	// insert test data
	users := []models.User{{
		Name: "John",
		Age:  20,
	}, {
		Name: "Mike",
		Age:  30,
	}}
	for _, user := range users {
		userRepo.Create(user)
	}

	// run test
	users, err = userUsecase.GetUsers()
	// assert
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "John", users[0].Name)
	assert.Equal(t, 20, users[0].Age)
	assert.Equal(t, "Mike", users[1].Name)
	assert.Equal(t, 30, users[1].Age)
}

func TestGetUserUsecase(t *testing.T) {
	// mock db
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	// make repository
	userRepo := repositories.NewUserRepository(db)
	// make usecase
	userUsecase := NewUserUsecase(userRepo)

	// insert test data
	userRepo.Create(models.User{
		Name: "John",
		Age:  20,
	})

	// run test
	user, err := userUsecase.GetUser("1")
	// assert
	assert.NoError(t, err)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, 20, user.Age)
}

func TestCreateUserUsecase(t *testing.T) {
	// mock db
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	// make repository
	userRepo := repositories.NewUserRepository(db)
	// make usecase
	userUsecase := NewUserUsecase(userRepo)

	// run test
	user, err := userUsecase.CreateUser(models.User{
		Name: "John",
		Age:  20,
	})

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, 20, user.Age)
}

func TestDeleteUserUsecase(t *testing.T) {
	// mock db
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	// make repository
	userRepo := repositories.NewUserRepository(db)
	// make usecase
	userUsecase := NewUserUsecase(userRepo)

	// insert test data
	userRepo.Create(models.User{
		Name: "John",
		Age:  20,
	})

	// run test
	err = userUsecase.DeleteUser("1")

	// assert
	assert.NoError(t, err)
}
