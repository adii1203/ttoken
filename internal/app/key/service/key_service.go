package service

import (
	"context"
	"time"

	"github.com/adii1203/ttoken/internal/db/repository"
	"github.com/adii1203/ttoken/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type KeyService struct {
	repo *repository.Queries
}

func NewKeyService(repo *repository.Queries) *KeyService {
	return &KeyService{
		repo: repo,
	}
}

func (s *KeyService) CreateKey(ctx context.Context, req *utils.CreateKeyRequestParams) (repository.ApiKey, string, error) {
	b, err := utils.NewKey(req.Prefix, 10)
	if err != nil {
		return repository.ApiKey{}, "", err
	}
	newKey := b.ToString()
	// hash the key using argon2
	hashedKey, err := utils.HashApiKey(newKey)
	if err != nil {
		return repository.ApiKey{}, "", err
	}

	id, err := uuid.Parse(req.ProjectId.String())
	if err != nil {
		return repository.ApiKey{}, "", err
	}

	apiKey, err := s.repo.CreateKey(ctx, repository.CreateKeyParams{
		ProjectID:   id,
		Prefix:      NewNullableText(req.Prefix),
		KeyHash:     hashedKey,
		Scopes:      []string{},
		Environment: "",
		ExpiresAt: pgtype.Timestamp{
			Time: time.Now(),
		},
	})

	if err != nil {
		return repository.ApiKey{}, "", err
	}

	return apiKey, newKey, nil
}

func (s *KeyService) VerifyProject(ctx context.Context, id uuid.UUID) (bool, error) {
	id, err := uuid.Parse(id.String())
	if err != nil {
		return false, nil
	}

	_, err = s.repo.GetProject(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func NewNullableText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}

	return pgtype.Text{
		String: *s,
		Valid:  true,
	}
}
