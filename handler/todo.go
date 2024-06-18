package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// ServeHTTP handles HTTP requests to the /todos endpoint.
func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.handlePost(w, r)
	} else if r.Method == http.MethodPut {
		h.handlePut(w, r)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// handlePost handles POST requests to create a new TODO.
func (h *TODOHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTODORequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request: failed to decode request body", http.StatusBadRequest)
		fmt.Printf("Error decoding request body: %v\n", err)
		return
	}

	if req.Subject == "" {
		http.Error(w, "bad request: subject cannot be empty", http.StatusBadRequest)
		fmt.Println("Subject cannot be empty")
		return
	}

	ctx := r.Context()
	todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	if err != nil {
		http.Error(w, "internal server error: "+err.Error(), http.StatusInternalServerError)
		fmt.Printf("Error creating TODO: %v\n", err)
		return
	}

	res := &model.CreateTODOResponse{
		TODO: *todo,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "internal server error: failed to encode response", http.StatusInternalServerError)
		fmt.Printf("Error encoding response: %v\n", err)
		return
	}
}

// handlePut handles PUT requests to update an existing TODO.
func (h *TODOHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	var req model.UpdateTODORequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request: failed to decode request body", http.StatusBadRequest)
		fmt.Printf("Error decoding request body: %v\n", err)
		return
	}

	if req.ID == 0 {
		http.Error(w, "bad request: ID cannot be zero", http.StatusBadRequest)
		fmt.Println("ID cannot be zero")
		return
	}

	if req.Subject == "" {
		http.Error(w, "bad request: subject cannot be empty", http.StatusBadRequest)
		fmt.Println("Subject cannot be empty")
		return
	}

	ctx := r.Context()
	todo, err := h.svc.UpdateTODO(ctx, req.ID, req.Subject, req.Description)
	if err != nil {
		if err == model.NewErrNotFound() {
			http.Error(w, "not found: TODO not found", http.StatusNotFound)
			fmt.Println("TODO not found")
			return
		}
		http.Error(w, "internal server error: "+err.Error(), http.StatusInternalServerError)
		fmt.Printf("Error updating TODO: %v\n", err)
		return
	}

	res := &model.UpdateTODOResponse{
		TODO: *todo,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "internal server error: failed to encode response", http.StatusInternalServerError)
		fmt.Printf("Error encoding response: %v\n", err)
		return
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
