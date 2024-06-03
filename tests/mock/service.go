package mock

import "errors"

var ErrInvalidUser = errors.New("invalid user data")

type User struct {
	ID    int
	Name  string
	Email string
}

//go:generate mockgen -source=service.go -destination=service_mock.go -package mock
type UserRepository interface {
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.GetUserByID(id)

}

func (s *UserService) AddUser(user *User) error {
	if user.Name == "" || user.Email == "" {
		return ErrInvalidUser
	}
	return s.repo.CreateUser(user)
}
