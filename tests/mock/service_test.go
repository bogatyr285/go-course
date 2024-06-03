package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockUserRepository(ctrl)

	userService := NewUserService(mockRepo)
	expectedUser := &User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	// Define mock behavior
	mockRepo.EXPECT().GetUserByID(5).Return(expectedUser, nil)

	// Call the method we want to test
	user, err := userService.GetUser(5)
	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)
	newUser := &User{Name: "Jane Doe", Email: "jane@example.com"}

	// Define mock behavior
	mockRepo.EXPECT().CreateUser(newUser).Return(nil)

	// Call the method we want to test
	err := userService.AddUser(newUser)
	// Assert expectations
	assert.NoError(t, err)
}

func TestAddUserEmptyUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)
	newUser := &User{Name: "sdfsdf", Email: "jane@example.com"}

	// Define mock behavior
	mockRepo.EXPECT().CreateUser(newUser).Return(ErrInvalidUser)

	// Call the method we want to test
	err := userService.AddUser(newUser)
	assert.ErrorIs(t, err, ErrInvalidUser)
}
