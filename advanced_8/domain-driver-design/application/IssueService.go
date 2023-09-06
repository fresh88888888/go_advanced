package application

import (
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/domain"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/persistence/db"
)

type IssueService struct {
	IssueRespository db.IssueRepository
}

func (i *IssueService) Issues() ([]*domain.Issue, error) {
	return i.IssueRespository.All()
}

func (i *IssueService) Create(u *domain.Issue) error {
	return i.IssueRespository.Create(u)
}

func (i *IssueService) Delete(id int64) error {
	return i.IssueRespository.Delete(id)
}

func (i *IssueService) Issue(id int64) (*domain.Issue, error) {
	return i.IssueRespository.GetById(id)
}
