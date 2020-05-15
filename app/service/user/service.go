package user

import "github.com/mmgopher/user-service/app/dao"

// Provider provides and interface to work with User service
type Provider interface {
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
