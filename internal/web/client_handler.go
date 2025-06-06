package web

import (
	"encoding/json"
	"net/http"

	"github.com/own4rd/ms-wallet-core/internal/usecase/create_client"
)

type WebClientHandler struct {
	CreateClientUseCase create_client.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase create_client.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		CreateClientUseCase: createClientUseCase,
	}
}
func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto create_client.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		println("Error decoding request body:", err.Error())
		return
	}

	output, err := h.CreateClientUseCase.Execute(dto)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		println("Error executing use case:", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		println("Error encoding response:", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}
