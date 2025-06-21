package service

import (
	"context"

	"github.com/adii1203/ttoken/internal/db/repository"
	"github.com/adii1203/ttoken/utils"
)

type ProjectService struct {
	repo *repository.Queries
}

func NewProjectService(repo *repository.Queries) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *utils.CreateProjectRequestParams) (repository.Project, error) {
	project, err := s.repo.CreateProject(ctx, req.Name)
	if err != nil {
		return repository.Project{}, err
	}
	return project, err
}
