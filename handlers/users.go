package handlers

import (
	"encoding/json"
	"net/http"
	dto "waysbuck/dto/result"
	"waysbuck/repositories"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}

		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}

	json.NewEncoder(w).Encode(response)
}
