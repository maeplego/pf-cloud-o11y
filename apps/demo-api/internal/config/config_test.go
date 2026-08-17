package config

import "testing"

func TestFromEnvRequiresOTLPEndpoint(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "")
	_, err := FromEnv()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestFromEnvParsesDebug(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://collector:4318")
	t.Setenv("ENABLE_DEBUG", "true")
	cfg, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.EnableDebug {
		t.Fatal("debug")
	}
}
