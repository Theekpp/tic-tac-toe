package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"tictactoe/internal/domain"
)

type Handler struct {
	service domain.Service
	repo    domain.Repository
}

func NewHandler(service domain.Service, repo domain.Repository) *Handler {
	return &Handler{service: service, repo: repo}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/game/") {
		h.makeMove(w, r)
		return
	}
	h.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
}

func (h *Handler) makeMove(w http.ResponseWriter, r *http.Request) {
	gameID := extractIDFromURL(r.URL.Path)
	if gameID == "" {
		h.writeError(w, http.StatusBadRequest, "invalid game ID in URL")
		return
	}

	parsedID, err := uuid.Parse(gameID)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid UUID: "+err.Error())
		return
	}

	var req MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid JSON body: "+err.Error())
		return
	}

	oldGame, err := h.repo.Get(r.Context(), parsedID.String())
	if err != nil {
		oldGame = &domain.Game{ID: parsedID, Board: domain.Board{}}
	}

	newGame := &domain.Game{ID: parsedID, Board: req.Board}

	if err = h.service.ValidateBoard(r.Context(), oldGame, newGame); err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	move, err := h.service.ComputeNextMove(r.Context(), newGame)
	if err != nil {
		if err.Error() == "game is over" {
			h.writeError(w, http.StatusConflict, "game is over")
			return
		}
		h.writeError(w, http.StatusBadRequest, "invalid move: "+err.Error())
		return
	}

	if move != nil {
		newGame.Board[move.Row][move.Col] = domain.CellO
	}

	if err := h.repo.Save(r.Context(), newGame); err != nil {
		h.writeError(w, http.StatusInternalServerError, "failed to save game")
		return
	}

	isFinish, winner, err := h.service.CheckCompletion(r.Context(), newGame)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, "failed to check game status")
		return
	}

	response := FromDomainGame(newGame, isFinish, winner)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func extractIDFromURL(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) != 2 || parts[0] != "game" {
		return ""
	}
	return parts[1]
}

func (h *Handler) writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
