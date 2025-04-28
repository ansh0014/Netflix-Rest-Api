package handlers

import (
	"net/http"
	"encoding/json"
	"monkeytype-backend/internal/models"
)

type TypingHandler struct{}

func (h *TypingHandler) HandleTypingRequest(w http.ResponseWriter, r *http.Request) {
	var typingSession models.Typing
	if err := json.NewDecoder(r.Body).Decode(&typingSession); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process the typing session (business logic would go here)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(typingSession)
}