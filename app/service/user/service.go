package user

import (
	"github.com/mmgopher/user-service/app/api/request"
	"github.com/mmgopher/user-service/app/dao"
	"github.com/mmgopher/user-service/app/httperrors"
	"github.com/mmgopher/user-service/app/model"
)

// Provider provides and interface to work with User service
type Provider interface {
	// GetUser returns User based on user ID
	GetUser(userID int) (*model.User, error)
	// DeleteUser deletes user from databse
	DeleteUser(userID int) error
}

// Service represents User service
type Service struct {
	userRepository dao.UserRepositoryProvider
}

// NewService creates new instance of Payment service.
func NewService(
	userRepository dao.UserRepositoryProvider,
) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

// GetUser returns User based on user ID
func (s Service) GetUser(userID int) (*model.User, error) {
	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, httperrors.InternalServerError.WithCause(err)
	}

	if user == nil {
		return nil, httperrors.EntityNotFoundError("user")
	}

	return user, nil
}

// DeleteUser deletes user from databse
func (s Service) DeleteUser(userID int) error {
	deleted, err := s.userRepository.Delete(userID)
	if err != nil {
		return httperrors.InternalServerError.WithCause(err)
	}

	if !deleted {
		return httperrors.EntityNotFoundError("user")
	}

	return nil
}

// CreateUser creates new user
func (s Service) CreateUser(request *request.CreateUser) (int, error) {
	userID, err := s.userRepository.Create(&model.User{
		Name:    request.Name,
		Surname: request.Surname,
		Gender:  request.Gender,
		Age:     request.Age,
		Address: request.Address,
	})

	if err != nil {
		return 0, httperrors.InternalServerError.WithCause(err)
	}

	return userID, nil
}
