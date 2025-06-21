package utils

import "github.com/google/uuid"

type CreateKeyRequestParams struct {
	Prefix    *string   `json:"prefix"`
	ProjectId uuid.UUID `json:"project_id" validate:"required"`
}

type CreateProjectRequestParams struct {
	Name string `json:"name" validate:"required"`
}
