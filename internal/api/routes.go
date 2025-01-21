package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	SequenceHandler *SequenceHandler
	StepHandler     *StepHandler
	GeneralHandler  *GeneralHandler
}

// NewRouter creates a new router and sets up all routes.
func NewRouter(routes *Routes) *mux.Router {
	router := mux.NewRouter()

	// Swagger UI route
	router.PathPrefix("/api/v1/docs/swagger-ui/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/api/v1/docs/openapi3.yaml"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
	))
	// Serve OpenAPI spec
	router.HandleFunc("/api/v1/docs/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving OpenAPI spec")
		filePath := filepath.Join("internal/api/docs", "openapi3.yaml")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, filePath)
	}).Methods(http.MethodGet)

	// Home page
	router.HandleFunc("/", routes.GeneralHandler.HomePage).Methods(http.MethodGet)

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// General routes
	api.HandleFunc("/health", routes.GeneralHandler.HealthCheck).Methods(http.MethodGet)
	api.HandleFunc("/info", routes.GeneralHandler.GetAPIInfo).Methods(http.MethodGet)

	// Sequence routes
	api.HandleFunc("/sequences", routes.SequenceHandler.CreateSequence).Methods(http.MethodPost)
	api.HandleFunc("/sequences/{id}", routes.SequenceHandler.UpdateTracking).Methods(http.MethodPut)
	api.HandleFunc("/sequences/{id}", routes.SequenceHandler.GetSequence).Methods(http.MethodGet)

	// Step routes
	api.HandleFunc("/steps", routes.StepHandler.CreateStep).Methods(http.MethodPost)
	api.HandleFunc("/steps/{id}", routes.StepHandler.UpdateStep).Methods(http.MethodPut)
	api.HandleFunc("/steps/{id}", routes.StepHandler.DeleteStep).Methods(http.MethodDelete)
	api.HandleFunc("/steps", routes.StepHandler.ListSteps).Methods(http.MethodGet)

	// Middleware (optional, e.g., logging)
	router.Use(LoggingMiddleware)

	return router
}
