package app

import (
	"fmt"
	"net/http"

	"ipfs-visualizer/config"
	"ipfs-visualizer/internal/http/middleware"
	"ipfs-visualizer/internal/http/handlers"
	"ipfs-visualizer/internal/kube"
	"ipfs-visualizer/internal/services"

	"github.com/go-chi/chi/v5"
)

type App struct {
	cfg       *config.Config
	kube      kube.KubeClient
	services  services.Container
	router    chi.Router
}

func NewApp(cfg *config.Config) (*App, error) {
	// init k8s client
	kc, err := kube.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	a := &App{
		cfg:    cfg,
		kube:   kc,
		router: chi.NewRouter(),
	}

	a.services = services.Container{
		Clusters: services.NewClusterService(kc),
		Nodes:    services.NewNodeService(kc),
		Status:   services.NewStatusService(kc),
		Logs:     services.NewLogsService(kc),
	}

	a.initRoutes()

	return a, nil
}

func (a *App) initRoutes() {
	r := a.router

	r.Use(middleware.RequestLogger)
	r.Use(middleware.Recovery)

	h := handlers.New(a.services)

	h.RegisterRoutes(r)
}

func (a *App) Start() error {
	addr := fmt.Sprintf("%s:%d", a.cfg.HTTP.Host, a.cfg.HTTP.Port)
	return http.ListenAndServe(addr, a.router)
}

func (a *App) handleGetClusters(w http.ResponseWriter, r *http.Request)        {}
func (a *App) handleCreateCluster(w http.ResponseWriter, r *http.Request)      {}
func (a *App) handleGetCluster(w http.ResponseWriter, r *http.Request)         {}
func (a *App) handleDeleteCluster(w http.ResponseWriter, r *http.Request)      {}
func (a *App) handleUpdateCluster(w http.ResponseWriter, r *http.Request)      {}
func (a *App) handleScaleCluster(w http.ResponseWriter, r *http.Request)       {}
func (a *App) handleRestartCluster(w http.ResponseWriter, r *http.Request)     {}

func (a *App) handleGetNodes(w http.ResponseWriter, r *http.Request)           {}
func (a *App) handleCreateNode(w http.ResponseWriter, r *http.Request)         {}
func (a *App) handleGetNode(w http.ResponseWriter, r *http.Request)            {}
func (a *App) handleDeleteNode(w http.ResponseWriter, r *http.Request)         {}
func (a *App) handleUpdateNode(w http.ResponseWriter, r *http.Request)         {}
func (a *App) handleGetNodeLogs(w http.ResponseWriter, r *http.Request)        {}

func (a *App) handleGetClusterStatus(w http.ResponseWriter, r *http.Request)   {}
