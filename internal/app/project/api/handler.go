package api

import (
	"net/http"

	"github.com/adii1203/ttoken/internal/app/project/service"
	"github.com/adii1203/ttoken/pkg/validator"
	"github.com/adii1203/ttoken/utils"
	"github.com/google/uuid"
)

type ProjectHandler struct {
	Service   *service.ProjectService
	Validator *validator.Validator
}

type CreateProjectResponse struct {
	ProjectId uuid.UUID `json:"project_id"`
}

func NewProjectHandler(service *service.ProjectService, validator *validator.Validator) *ProjectHandler {
	return &ProjectHandler{
		Service:   service,
		Validator: validator,
	}
}

func (p *ProjectHandler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	req := &utils.CreateProjectRequestParams{}

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

	err = p.Validator.ValidateStruct(req)
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

	project, err := p.Service.CreateProject(r.Context(), req)
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

	utils.Success(w, 200, CreateProjectResponse{
		ProjectId: project.ID,
	})

}
