package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"ipfs-visualizer/internal/app/models"
	"ipfs-visualizer/internal/app/errors"
	"ipfs-visualizer/internal/app/services"
)

type ClusterHandler struct {
	svc *services.ClusterService
}

func NewClusterHandler(svc *services.ClusterService) *ClusterHandler {
	return &ClusterHandler{svc: svc}
}

///////////////////////////////////////////////////////////////
// GET /clusters
///////////////////////////////////////////////////////////////
func (h *ClusterHandler) GetClusters(w http.ResponseWriter, r *http.Request) {
	clusters, err := h.svc.GetClusters(r.Context())
	if err != nil {
		errors.Write(w, err)
		return
	}

	writeJSON(w, http.StatusOK, clusters)
}

///////////////////////////////////////////////////////////////
// POST /clusters
///////////////////////////////////////////////////////////////
func (h *ClusterHandler) CreateCluster(w http.ResponseWriter, r *http.Request) {
	var req models.ClusterCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteBadRequest(w, "invalid json")
		return
	}

	cluster, err := h.svc.CreateCluster(r.Context(), req)
	if err != nil {
		errors.Write(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, cluster)
}

///////////////////////////////////////////////////////////////
// GET /clusters/{clusterId}
///////////////////////////////////////////////////////////////
func (h *ClusterHandler) GetCluster(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clusterId")

	cluster, err := h.svc.GetCluster(r.Context(), id)
	if err != nil {
		errors.Write(w, err)
		return
	}

	writeJSON(w, http.StatusOK, cluster)
}

///////////////////////////////////////////////////////////////
// DELETE /clusters/{clusterId}
///////////////////////////////////////////////////////////////
func (h *ClusterHandler) DeleteCluster(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clusterId")

	if err := h.svc.DeleteCluster(r.Context(), id); err != nil {
		errors.Write(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

///////////////////////////////////////////////////////////////
// PUT /clusters/{clusterId}
///////////////////////////////////////////////////////////////
func (h *ClusterHandler) UpdateCluster(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clusterId")

	var req models.ClusterUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteBadRequest(w, "invalid json")
		return
	}

	cluster, err := h.svc.UpdateCluster(r.Context(), id, req)
	if err != nil {
		errors.Write(w, err)
		return
	}

	writeJSON(w, http.StatusOK, cluster)
}

///////////////////////////////////////////////////////////////
// POST /clusters/{clusterId}/scale
///////////////////////////////////////////////////////////////
func (h *ClusterHandler) ScaleCluster(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clusterId")

	var payload struct {
		Replicas int `json:"replicas"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		errors.WriteBadRequest(w, "invalid json")
		return
	}

	if payload.Replicas < 1 {
		errors.WriteBadRequest(w, "replicas must be >= 1")
		return
	}

	if err := h.svc.ScaleCluster(r.Context(), id, payload.Replicas); err != nil {
		errors.Write(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"status":   "scaling",
		"replicas": payload.Replicas,
	})
}

///////////////////////////////////////////////////////////////
// POST /clusters/{clusterId}/restart
///////////////////////////////////////////////////////////////
func (h *ClusterHandler) RestartCluster(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clusterId")

	if err := h.svc.RestartCluster(r.Context(), id); err != nil {
		errors.Write(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "restarting",
	})
}

///////////////////////////////////////////////////////////////
// GET /clusters/{clusterId}/status
///////////////////////////////////////////////////////////////
func (h *ClusterHandler) GetClusterStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "clusterId")

	status, err := h.svc.GetClusterStatus(r.Context(), id)
	if err != nil {
		errors.Write(w, err)
		return
	}

	writeJSON(w, http.StatusOK, status)
}

///////////////////////////////////////////////////////////////
// helpers
///////////////////////////////////////////////////////////////
func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
