package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adii1203/ttoken/internal/app/user/service"
	"github.com/adii1203/ttoken/pkg/svix"
	"github.com/adii1203/ttoken/utils"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) ClerkHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	wh := svix.InitSvix()

	err = wh.Verify(payload, headers)
	if err != nil {
		fmt.Println("varification error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var clerkEvt utils.ClerkUserCreated

	if err := json.Unmarshal(payload, &clerkEvt); err != nil {
		fmt.Println("json error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// fmt.Println(clerkEvt)
	// Todo: Add usre to database

	if clerkEvt.Type == "user.created" {
		err = h.Service.CreateUser(r.Context(), &utils.CreateUserRequestParams{
			Id:           clerkEvt.Data.Id,
			FirstName:    clerkEvt.Data.FirstName,
			LastName:     clerkEvt.Data.LastName,
			EmailAddress: clerkEvt.Data.EmailAddresses[0].EmailAddress,
		})

		if err != nil {
			fmt.Println("db error: ", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
