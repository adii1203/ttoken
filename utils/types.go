package utils

import "github.com/google/uuid"

type CreateKeyRequestParams struct {
	Prefix    *string   `json:"prefix"`
	ProjectId uuid.UUID `json:"project_id" validate:"required"`
}

type CreateProjectRequestParams struct {
	Name string `json:"name" validate:"required"`
}

type ClerkUserCreated struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Data   struct {
		Id             string         `json:"id"`
		FirstName      string         `json:"first_name"`
		LastName       string         `json:"last_name"`
		EmailAddresses []EmailAddress `json:"email_addresses"`
	} `json:"data"`
}

type EmailAddress struct {
	EmailAddress string `json:"email_address"`
}

type CreateUserRequestParams struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
}
