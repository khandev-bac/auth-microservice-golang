package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/khandev-bac/lemon/internals/db/dto"
	"github.com/khandev-bac/lemon/internals/db/model"
	"github.com/khandev-bac/lemon/internals/middleware"
	"github.com/khandev-bac/lemon/internals/service"
	jwttoken "github.com/khandev-bac/lemon/utils/jwtToken"
)

type UserHandler struct {
	service *service.UserService
}

func NewHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user dto.SignupDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	createdUser, err := h.service.Signup(r.Context(), model.User{
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		Password:  user.Password,
	})
	if err != nil {
		http.Error(w, "Fialed to signup user"+err.Error(), http.StatusInternalServerError)
		return
	}
	tokens, _ := jwttoken.GenerateTokens(createdUser.ID, createdUser.UserEmail)
	_ = h.service.UpdateRefreshToken(r.Context(), createdUser.ID, tokens.RefreshToken)
	response := map[string]interface{}{
		"error":        false,
		"message":      "successfully sign up",
		"id":           createdUser.ID,
		"accessToken":  tokens.AccessToken,
		"refreshToken": tokens.RefreshToken,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request type", http.StatusBadRequest)
		return
	}
	newUser, err := h.service.Login(r.Context(), req.UserEmail, req.Password)
	if err != nil {
		http.Error(w, "Invalid password/email", http.StatusUnauthorized)
		return
	}
	tokens, _ := jwttoken.GenerateTokens(newUser.ID, newUser.UserEmail)
	_ = h.service.UpdateRefreshToken(r.Context(), newUser.ID, tokens.RefreshToken)
	response := map[string]interface{}{
		"error":        false,
		"message":      "successfully login",
		"accessToken":  tokens.AccessToken,
		"refreshToken": tokens.RefreshToken,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	userIDstr, err := middleware.GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	uid, err := uuid.Parse(userIDstr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	err = h.service.UpdateRefreshToken(r.Context(), uid, "")
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged out successfully",
	})
}
