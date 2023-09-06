package application

import "umbrella.github.com/advanced_go/advanced_8/domain-driver-design/domain"

type ProjectService struct {
	ProjectRespoitory domain.ProjectRespority
}

func (p *ProjectService) Projects() ([]*domain.Project, error) {
	return p.ProjectRespoitory.All()
}

func (p *ProjectService) Create(u *domain.Project) error {
	return p.ProjectRespoitory.Create(u)
}

func (p *ProjectService) Delete(id int64) error {
	return p.ProjectRespoitory.Delete(id)
}

func (p *ProjectService) Project(id int64) (*domain.Project, error) {
	return p.ProjectRespoitory.GetById(id)
}
