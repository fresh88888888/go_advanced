package memory

import (
	"errors"

	"github.com/patrickmn/go-cache"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/domain"
)

const (
	ProjectAllKey = "Project:all"
	ProjectLastId = "Project:lastId"
)

type ProjectRespoitory struct {
	db *cache.Cache
}

func NewProjectRepository() *ProjectRespoitory {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(ProjectLastId, int64(0))
	db.SetDefault(ProjectAllKey, []*domain.Issue{})

	return &ProjectRespoitory{
		db: db,
	}
}

func (r *ProjectRespoitory) All() ([]*domain.Project, error) {
	result, ok := r.db.Get(ProjectAllKey)
	if ok {
		return result.([]*domain.Project), nil
	} else {
		return nil, errors.New("Empty list.")
	}
}

func (r *ProjectRespoitory) GetById(id int64) (*domain.Project, error) {
	result, ok := r.db.Get(ProjectAllKey)
	if ok {
		items := result.([]*domain.Project)
		for _, project := range items {
			if id == project.Id {
				return project, nil
			}
		}

		return nil, errors.New("Not Found")
	}
	return nil, errors.New("Not Found")
}

func (r *ProjectRespoitory) Create(p *domain.Project) error {
	id, _ := r.db.IncrementInt64(ProjectLastId, int64(1))
	p.Id = id

	result, ok := r.db.Get(IssueAllKey)
	if ok {
		result = append(result.([]*domain.Project), p)
		r.db.Set(ProjectAllKey, result, cache.NoExpiration)
	}

	return nil
}

func (r *ProjectRespoitory) Delete(id int64) error {
	result, ok := r.db.Get(ProjectAllKey)
	if ok {
		items := result.([]*domain.Project)
		for i, project := range items {
			if project.Id == id {
				items = append(items[:i], items[i+1:]...)
				r.db.Set(ProjectAllKey, items, cache.NoExpiration)
			}
		}
		return errors.New("Not Found")
	}
	return errors.New("Not Found")
}
