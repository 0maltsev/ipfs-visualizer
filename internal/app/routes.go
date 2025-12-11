package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) Routes() http.Handler {
	r := chi.NewRouter()

	// Base middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// -----------------------------------------------------
	// Clusters
	// -----------------------------------------------------
	r.Route("/clusters", func(r chi.Router) {

		// GET /clusters
		r.Get("/", a.handleGetClusters)

		// POST /clusters
		r.Post("/", a.handleCreateCluster)

		// /clusters/{clusterId}
		r.Route("/{clusterId}", func(r chi.Router) {

			r.Get("/", a.handleGetCluster)
			r.Delete("/", a.handleDeleteCluster)
			r.Put("/", a.handleUpdateCluster)

			// POST /clusters/{clusterId}/scale
			r.Post("/scale", a.handleScaleCluster)

			// POST /clusters/{clusterId}/restart
			r.Post("/restart", a.handleRestartCluster)

			// -----------------------------------------------------
			// Nodes in cluster
			// -----------------------------------------------------
			r.Route("/nodes", func(r chi.Router) {

				// GET /clusters/{clusterId}/nodes
				r.Get("/", a.handleGetNodes)

				// POST /clusters/{clusterId}/nodes
				r.Post("/", a.handleCreateNode)

				// /clusters/{clusterId}/nodes/{nodeId}
				r.Route("/{nodeId}", func(r chi.Router) {

					r.Get("/", a.handleGetNode)
					r.Delete("/", a.handleDeleteNode)
					r.Put("/", a.handleUpdateNode)

					// GET /clusters/{clusterId}/nodes/{nodeId}/logs
					r.Get("/logs", a.handleGetNodeLogs)
				})
			})

			// -----------------------------------------------------
			// Status
			// -----------------------------------------------------
			r.Get("/status", a.handleGetClusterStatus)
		})
	})

	return r
}
