package memory

import (
	"errors"

	"github.com/patrickmn/go-cache"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/domain"
)

const (
	UserAllKey  = "Project:all"
	UsertLastId = "Project:lastId"
)

type UserRepository struct {
	db *cache.Cache
}

func NewUserRepository() *UserRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(UsertLastId, int64(0))
	db.SetDefault(UserAllKey, []*domain.User{})

	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) All() ([]*domain.User, error) {
	result, ok := r.db.Get(UserAllKey)
	if ok {
		return result.([]*domain.User), nil
	} else {
		return nil, errors.New("Empty list.")
	}
}

func (r *UserRepository) GetById(id int64) (*domain.User, error) {
	result, ok := r.db.Get(UserAllKey)
	if ok {
		items := result.([]*domain.User)
		for _, user := range items {
			if id == user.Id {
				return user, nil
			}
		}

		return nil, errors.New("Not Found")
	}
	return nil, errors.New("Not Found")
}

func (r *UserRepository) Create(p *domain.User) error {
	id, _ := r.db.IncrementInt64(UsertLastId, int64(1))
	p.Id = id

	result, ok := r.db.Get(UserAllKey)
	if ok {
		result = append(result.([]*domain.User), p)
		r.db.Set(UserAllKey, result, cache.NoExpiration)
	}

	return nil
}

func (r *UserRepository) Delete(id int64) error {
	result, ok := r.db.Get(UserAllKey)
	if ok {
		items := result.([]*domain.User)
		for i, user := range items {
			if user.Id == id {
				items = append(items[:i], items[i+1:]...)
				r.db.Set(UserAllKey, items, cache.NoExpiration)
			}
		}
		return errors.New("Not Found")
	}
	return errors.New("Not Found")
}
