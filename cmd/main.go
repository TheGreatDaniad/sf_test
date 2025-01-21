package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"sf_test/config"
	"sf_test/internal/api"
	"sf_test/internal/core"
	"sf_test/internal/db"
	"sf_test/pkg/logger"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	appLogger, err := logger.NewLogger("app.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	appLogger.Info("Starting application...")

	// Database connection
	dbConn, err := db.NewDB(
		"postgres://" + cfg.Database.User + ":" + cfg.Database.Password +
			"@" + cfg.Database.Host + ":" + strconv.Itoa(cfg.Database.Port) +
			"/" + cfg.Database.DBName + "?sslmode=" + cfg.Database.SSLMode,
	)
	if err != nil {
		appLogger.Error(err)
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if cfg.App.DoMigrations {
		err = dbConn.MigrateDB()
		if err != nil {
			appLogger.Error(err)
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}

	defer dbConn.Close()

	appLogger.Info("Connected to the database")

	// Initialize repositories
	sequenceRepo := db.NewSequenceRepository(dbConn)
	stepRepo := db.NewStepRepository(dbConn)

	// Initialize services
	sequenceService := core.NewSequenceService(sequenceRepo)
	stepService := core.NewStepService(stepRepo)

	// Initialize handlers
	sequenceHandler := api.NewSequenceHandler(sequenceService)
	stepHandler := api.NewStepHandler(stepService)
	generalHandler := api.NewGeneralHandler(cfg.App.Version)
	// Create router and routes
	router := api.NewRouter(&api.Routes{
		SequenceHandler: sequenceHandler,
		StepHandler:     stepHandler,
		GeneralHandler:  generalHandler,
	})

	// Add Prometheus metrics endpoint if enabled
	if cfg.Metrics.Enabled {
		router.Handle(cfg.Metrics.Path, promhttp.Handler())
		appLogger.Info("Prometheus metrics enabled at " + cfg.Metrics.Path)
		go func() {
			metricsAddr := ":" + strconv.Itoa(cfg.Metrics.Port)
			appLogger.Info("Metrics server running on " + metricsAddr)
			log.Fatal(http.ListenAndServe(metricsAddr, nil))
		}()
	}

	// Start the HTTP server
	serverAddr := ":" + strconv.Itoa(cfg.App.Port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	appLogger.Info("Starting server on port " + strconv.Itoa(cfg.App.Port))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		appLogger.Error(err)
		log.Fatalf("Server error: %v", err)
	}

	appLogger.Info("Server stopped")
}
