package web

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/portfolio/pf-cloud-o11y/demo-api/internal/route"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Server is the demo API HTTP surface.
type Server struct {
	EnableDebug bool
	mux         *http.ServeMux
}

// NewServer wires routes with otelhttp wrappers.
func NewServer(enableDebug bool) *Server {
	s := &Server{EnableDebug: enableDebug, mux: http.NewServeMux()}
	s.mux.Handle("GET /health", otelhttp.NewHandler(http.HandlerFunc(s.health), "GET /health"))
	s.mux.Handle("GET /ready", otelhttp.NewHandler(http.HandlerFunc(s.ready), "GET /ready"))
	s.mux.Handle("GET /work/{id}", otelhttp.NewHandler(http.HandlerFunc(s.work), "GET /work/{id}"))
	if enableDebug {
		s.mux.Handle("POST /debug/slow", otelhttp.NewHandler(http.HandlerFunc(s.debugSlow), "POST /debug/slow"))
		s.mux.Handle("POST /debug/fail", otelhttp.NewHandler(http.HandlerFunc(s.debugFail), "POST /debug/fail"))
	}
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) ready(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

func (s *Server) work(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/work/")
	pattern := route.Pattern(r.Method, r.URL.Path)
	slog.InfoContext(r.Context(), "work handled", "http.route", pattern, "work_id", id)
	time.Sleep(25 * time.Millisecond)
	writeJSON(w, http.StatusOK, map[string]string{"id": id, "status": "done"})
}

func (s *Server) debugSlow(w http.ResponseWriter, r *http.Request) {
	if !s.EnableDebug {
		http.NotFound(w, r)
		return
	}
	ms, _ := strconv.Atoi(r.URL.Query().Get("ms"))
	if ms <= 0 {
		ms = 500
	}
	if ms > 5000 {
		ms = 5000
	}
	slog.InfoContext(r.Context(), "debug slow", "delay_ms", ms)
	time.Sleep(time.Duration(ms) * time.Millisecond)
	writeJSON(w, http.StatusOK, map[string]any{"delayed_ms": ms})
}

func (s *Server) debugFail(w http.ResponseWriter, r *http.Request) {
	if !s.EnableDebug {
		http.NotFound(w, r)
		return
	}
	slog.WarnContext(r.Context(), "debug fail injected")
	http.Error(w, "injected failure", http.StatusInternalServerError)
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
