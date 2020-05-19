// +build unit

package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mmgopher/user-service/app/api/request"
	"github.com/mmgopher/user-service/app/dao"
	"github.com/mmgopher/user-service/app/httperrors"
	"github.com/mmgopher/user-service/app/model"
)

func TestGetUserOK(t *testing.T) {
	userID := 5001
	model := model.User{
		ID:      userID,
		Name:    "name",
		Surname: "surname",
		Gender:  "male",
		Age:     10,
		Address: "address",
	}
	mockUserRepository := dao.MockUserRepositoryProvider{}
	mockUserRepository.On("GetByID", userID).Return(&model, nil)
	service := NewService(&mockUserRepository)
	user, err := service.GetUser(userID)
	assert.Nil(t, err)
	assert.Equal(t, model, *user)
}

func TestGetUserNotFound(t *testing.T) {
	userID := 5001
	mockUserRepository := dao.MockUserRepositoryProvider{}
	mockUserRepository.On("GetByID", userID).Return(nil, nil)
	service := NewService(&mockUserRepository)
	user, err := service.GetUser(userID)
	require.NotNil(t, err)
	assert.Nil(t, user)
	assert.EqualError(t, httperrors.EntityNotFoundError("user"), err.Error())
}

func TestDeleteUserOK(t *testing.T) {
	userID := 5001
	mockUserRepository := dao.MockUserRepositoryProvider{}
	mockUserRepository.On("Delete", userID).Return(true, nil)
	service := NewService(&mockUserRepository)
	err := service.DeleteUser(userID)
	assert.Nil(t, err)
}

func TestDeleteUserNotFound(t *testing.T) {
	userID := 5001
	mockUserRepository := dao.MockUserRepositoryProvider{}
	mockUserRepository.On("Delete", userID).Return(false, nil)
	service := NewService(&mockUserRepository)
	err := service.DeleteUser(userID)
	require.NotNil(t, err)
	assert.EqualError(t, httperrors.EntityNotFoundError("user"), err.Error())
}

func TestCreateUserOK(t *testing.T) {
	newUserID := 1
	request := request.CreateUser{
		Name:    "name",
		Surname: "surname",
		Gender:  "male",
		Age:     10,
		Address: "address",
	}

	model := model.User{
		Name:    "name",
		Surname: "surname",
		Gender:  "male",
		Age:     10,
		Address: "address",
	}

	mockUserRepository := dao.MockUserRepositoryProvider{}
	mockUserRepository.On("Create", &model).Return(newUserID, nil)
	mockUserRepository.On("CheckIfExistWithNameAndSurname", request.Name, request.Surname).Return(false, nil)
	service := NewService(&mockUserRepository)
	id, err := service.CreateUser(&request)
	assert.Equal(t, newUserID, id)
	assert.Nil(t, err)
}

func TestCreateUserAlreadyRegistered(t *testing.T) {

	request := request.CreateUser{
		Name:    "name",
		Surname: "surname",
		Gender:  "male",
		Age:     10,
		Address: "address",
	}

	mockUserRepository := dao.MockUserRepositoryProvider{}
	mockUserRepository.On("CheckIfExistWithNameAndSurname", request.Name, request.Surname).Return(true, nil)
	service := NewService(&mockUserRepository)
	id, err := service.CreateUser(&request)
	require.NotNil(t, err)
	assert.Equal(t, 0, id)
	assert.EqualError(t, httperrors.UserAlreadyRegistered, err.Error())
}

func TestCreateUserValidationError(t *testing.T) {

	request := request.CreateUser{
		Name:    "",
		Surname: "surname",
		Gender:  "male",
		Age:     10,
		Address: "address",
	}
	service := NewService(&dao.MockUserRepositoryProvider{})
	id, err := service.CreateUser(&request)
	require.NotNil(t, err)
	assert.Equal(t, 0, id)
	assert.EqualError(t, httperrors.UserNameEmpty, err.Error())
}
