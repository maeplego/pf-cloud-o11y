package config

import (
	"fmt"
	"os"
	"strings"
)

// Config holds process settings from the environment.
type Config struct {
	HTTPAddr     string
	ServiceName  string
	EnableDebug  bool
	OTelEndpoint string
}

// FromEnv reads DEMO_* and OTEL_* variables.
func FromEnv() (Config, error) {
	cfg := Config{
		HTTPAddr:     envOr("DEMO_HTTP_ADDR", ":8080"),
		ServiceName:  envOr("OTEL_SERVICE_NAME", "demo-api"),
		EnableDebug:  envBool("ENABLE_DEBUG", false),
		OTelEndpoint: strings.TrimSpace(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
	}
	if cfg.OTelEndpoint == "" {
		return Config{}, fmt.Errorf("OTEL_EXPORTER_OTLP_ENDPOINT is required")
	}
	return cfg, nil
}

func envOr(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func envBool(key string, fallback bool) bool {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	switch strings.ToLower(v) {
	case "1", "true", "yes", "on":
		return true
	case "0", "false", "no", "off":
		return false
	default:
		return fallback
	}
}
