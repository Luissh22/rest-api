package comment

import (
	"encoding/json"
	"github.com/gorilla/mux"
	httpConstants "github.com/luissh22/rest-api/internal/constants/http"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Router  *mux.Router
	service Service
}

func NewHandler(r *mux.Router, service Service) *Handler {
	return &Handler{
		Router:  r,
		service: service,
	}
}

func (h *Handler) SetupRoutes() {
	h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods(httpConstants.GET)
	h.Router.HandleFunc("/api/v1/comment", h.GetAllComments).Methods(httpConstants.GET)
	h.Router.HandleFunc("/api/v1/comment", h.PostComment).Methods(httpConstants.POST)
	h.Router.HandleFunc("/api/v1/ping", h.Ping)
}

type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Server is up"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to encode response", err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	unsignedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to parse comment id", err)
	}

	comment, err := h.service.GetComment(uint(unsignedId))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch comment", err)
		return
	}

	if err = json.NewEncoder(w).Encode(comment); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to encode response", err)
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.service.GetAllComments()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch all comments", err)
		return
	}

	if err = json.NewEncoder(w).Encode(comments); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to encode response", err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to parse request body", err)
		return
	}

	comment, err := h.service.PostComment(comment)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to store comment", err)
		return
	}

	if err = json.NewEncoder(w).Encode(comment); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to encode response", err)
	}
}

func respondWithError(w http.ResponseWriter, statusCode int, message string, error error) {
	w.WriteHeader(statusCode)
	response := Response{
		Message: message,
		Error:   error.Error(),
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal(err)
	}
}
