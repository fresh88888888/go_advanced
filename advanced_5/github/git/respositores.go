package git

import (
	"context"
	"fmt"
	"net/http"
)

const (
	reposPathWithUser = "users/%v/repos"
	defaultReposPath  = "user/repos"
)

type RepositoriesService struct {
	client *Client
}

// Repository represent a git repository.
type Repository struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	GitURL      string `json:"git_url,omitempty"`
}

// List of Repositories for a user
func (rs *RepositoriesService) List(ctx context.Context, user string) ([]*Repository, *http.Response, error) {
	var path string
	if user != "" {
		path = fmt.Sprintf(reposPathWithUser, user)
	} else {
		path = defaultReposPath
	}
	req, err := rs.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	var repos []*Repository
	resp, err := rs.client.Do(ctx, req, &repos)
	if err != nil {
		return nil, resp, err
	}

	return repos, resp, nil
}
