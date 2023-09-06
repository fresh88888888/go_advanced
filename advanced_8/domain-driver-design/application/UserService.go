package application

import (
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/domain"
)

type UserService struct {
	UserRespoitory domain.UserRepository
}

func (us *UserService) Users() ([]*domain.User, error) {
	return us.UserRespoitory.All()
}

func (us *UserService) Create(u *domain.User) error {
	return us.UserRespoitory.Create(u)
}

func (us *UserService) Delete(id int64) error {
	return us.UserRespoitory.Delete(id)
}

func (us *UserService) User(id int64) (*domain.User, error) {
	return us.UserRespoitory.GetById(id)
}
