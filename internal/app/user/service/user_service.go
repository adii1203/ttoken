package service

import (
	"context"
	"fmt"

	"github.com/adii1203/ttoken/internal/db/repository"
	"github.com/adii1203/ttoken/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo *repository.Queries
}

func NewUserService(repo *repository.Queries) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *utils.CreateUserRequestParams) error {
	if req.Id == "" || req.EmailAddress == "" {
		return fmt.Errorf("invalid request data")
	}

	_, err := s.repo.CreateUser(ctx, repository.CreateUserParams{
		ClerkID: pgtype.Text{
			String: req.Id,
			Valid:  true,
		},
		Email:     req.EmailAddress,
		FirstName: req.FirstName,
		LastName: pgtype.Text{
			String: req.LastName,
			Valid:  true,
		},
	})

	if err != nil {
		return fmt.Errorf("error creating user: %v", err.Error())
	}

	return nil
}
