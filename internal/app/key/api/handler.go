package api

import (
	"context"
	"net/http"

	"github.com/adii1203/ttoken/internal/app/key/service"
	v "github.com/adii1203/ttoken/pkg/validator"
	"github.com/adii1203/ttoken/utils"
	"github.com/google/uuid"
)

type KeyHandler struct {
	Service   *service.KeyService
	Validator *v.Validator
}

func NewKeyHandler(service *service.KeyService, validator *v.Validator) *KeyHandler {
	return &KeyHandler{
		Service:   service,
		Validator: validator,
	}
}

type CreateKeyResponse struct {
	KeyId uuid.UUID `json:"key_id"`
	Key   string    `json:"key"`
}

func (h *KeyHandler) CreateKeyHandler(w http.ResponseWriter, r *http.Request) {
	req := &utils.CreateKeyRequestParams{}

	err := utils.DecodeJSON(w, r, req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: utils.ErrorM{
				Code:      "",
				Message:   err.Error(),
				RequestId: "",
			},
		})
		return
	}

	err = h.Validator.ValidateStruct(req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: utils.ErrorM{
				Code:      "",
				Message:   err.Error(),
				RequestId: "",
			},
		})
		return
	}

	// TODO: verify the projecrId
	isValidProject, err := h.Service.VerifyProject(r.Context(), req.ProjectId)
	if err != nil || !isValidProject {
		utils.Error(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: utils.ErrorM{
				Code:      "BAD_REQUEST",
				Message:   "invalid project id",
				RequestId: "",
			},
		})
		return
	}

	newKey, key, err := h.Service.CreateKey(context.Background(), req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, utils.ErrorResponse{
			Error: utils.ErrorM{
				Code:      "",
				Message:   err.Error(),
				RequestId: "",
			},
		})
		return
	}

	utils.Success(w, http.StatusOK, CreateKeyResponse{
		KeyId: newKey.ID,
		Key:   key,
	})
}
