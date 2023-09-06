package memory

import (
	"errors"

	"github.com/patrickmn/go-cache"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/domain"
)

const (
	IssueAllKey = "Issue:all"
	IssueLastId = "Issue:lastId"
)

type IssueRepository struct {
	db *cache.Cache
}

func NewIssueRepository() *IssueRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(IssueLastId, int64(0))
	db.SetDefault(IssueAllKey, []*domain.Issue{})

	return &IssueRepository{
		db: db,
	}
}

func (r *IssueRepository) ALL() ([]*domain.Issue, error) {
	result, ok := r.db.Get(IssueAllKey)
	if ok {
		return result.([]*domain.Issue), nil
	} else {
		return nil, errors.New("Empty list.")
	}
}

func (r *IssueRepository) GetById(id int64) (*domain.Issue, error) {
	result, ok := r.db.Get(IssueAllKey)
	if ok {
		items := result.([]*domain.Issue)
		for _, issue := range items {
			if id == issue.IssueId {
				return issue, nil
			}
		}

		return nil, errors.New("Not Found")
	}
	return nil, errors.New("Not Found")
}

func (r *IssueRepository) Create(u *domain.Issue) error {
	id, _ := r.db.IncrementInt64(IssueLastId, int64(1))
	u.IssueId = id

	result, ok := r.db.Get(IssueAllKey)
	if ok {
		result = append(result.([]*domain.Issue), u)
		r.db.Set(IssueAllKey, result, cache.NoExpiration)
	}

	return nil
}

func (r *IssueRepository) Delete(id int64) error {
	result, ok := r.db.Get(IssueAllKey)
	if ok {
		items := result.([]*domain.Issue)
		for i, issue := range items {
			if issue.IssueId == id {
				items = append(items[:i], items[i+1:]...)
				r.db.Set(IssueAllKey, items, cache.NoExpiration)
			}
		}
		return errors.New("Not Found")
	}
	return errors.New("Not Found")
}
