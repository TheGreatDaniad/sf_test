package api

import (
	"html/template"
	"net/http"
	"os"
	"runtime"
	"time"
)

type GeneralHandler struct {
	startTime time.Time
	version   string
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type InfoResponse struct {
	Version      string `json:"version"`
	GoVersion    string `json:"goVersion"`
	Uptime       string `json:"uptime"`
	Environment  string `json:"environment"`
	TotalMemory  uint64 `json:"totalMemory"`
	NumGoroutine int    `json:"numGoroutine"`
}

func NewGeneralHandler(version string) *GeneralHandler {
	return &GeneralHandler{
		startTime: time.Now(),
		version:   version,
	}
}

// HealthCheck handles the health check endpoint
func (h *GeneralHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	health := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	WriteResponse(w, http.StatusOK, SuccessResponse(health, "Health check successful"))
}

// GetAPIInfo returns general information about the API
func (h *GeneralHandler) GetAPIInfo(w http.ResponseWriter, r *http.Request) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	info := InfoResponse{
		Version:      h.version,
		GoVersion:    runtime.Version(),
		Uptime:       time.Since(h.startTime).String(),
		Environment:  getEnvironment(),
		TotalMemory:  mem.TotalAlloc,
		NumGoroutine: runtime.NumGoroutine(),
	}

	WriteResponse(w, http.StatusOK, SuccessResponse(info, "API information retrieved successfully"))
}

// HomePage serves the API home page
func (h *GeneralHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Sequence Flow API</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            line-height: 1.6;
            max-width: 800px;
            margin: 0 auto;
            padding: 2rem;
            color: #333;
        }
        h1 {
            color: #2c3e50;
            border-bottom: 2px solid #eee;
            padding-bottom: 0.5rem;
        }
        .info-box {
            background: #f8f9fa;
            border: 1px solid #dee2e6;
            border-radius: 4px;
            padding: 1rem;
            margin: 1rem 0;
        }
        .links {
            margin: 2rem 0;
        }
        a {
            color: #007bff;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
        .version {
            color: #6c757d;
            font-size: 0.9rem;
        }
    </style>
</head>
<body>
    <h1>Sequence Flow API</h1>
    <div class="info-box">
        <p>Welcome to the Sequence Flow API. This API provides endpoints for managing sequences and steps in a workflow.</p>
        <p class="version">Version: {{.Version}}</p>
    </div>
    <div class="links">
        <h2>Quick Links</h2>
        <ul>
            <li><a href="/api/v1/docs/swagger-ui/">API Documentation (Swagger UI)</a></li>
            <li><a href="/api/v1/health">Health Check</a></li>
            <li><a href="/api/v1/info">API Information</a></li>
        </ul>
    </div>

</body>
</html>
`

	t, err := template.New("home").Parse(tmpl)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, ErrorResponse(err.Error(), "Failed to render home page"))
		return
	}

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	data := struct {
		Version     string
		Environment string
		GoVersion   string
		Uptime      string
	}{
		Version:     h.version,
		Environment: getEnvironment(),
		GoVersion:   runtime.Version(),
		Uptime:      time.Since(h.startTime).String(),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.Execute(w, data); err != nil {
		WriteResponse(w, http.StatusInternalServerError, ErrorResponse(err.Error(), "Failed to render home page"))
		return
	}
}

func getEnvironment() string {
	// This should be set through environment variables in production
	env := "development"
	if envVar := os.Getenv("APP_ENV"); envVar != "" {
		env = envVar
	}
	return env
}
