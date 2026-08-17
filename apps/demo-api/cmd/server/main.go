package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/portfolio/pf-cloud-o11y/demo-api/internal/config"
	"github.com/portfolio/pf-cloud-o11y/demo-api/internal/telemetry"
	"github.com/portfolio/pf-cloud-o11y/demo-api/internal/web"
)

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	shutdown, err := telemetry.Init(ctx, cfg.ServiceName, cfg.OTelEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: web.NewServer(cfg.EnableDebug),
	}

	go func() {
		log.Printf("demo-api listening on %s debug=%v (learning stack, not for production)", cfg.HTTPAddr, cfg.EnableDebug)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	_ = shutdown(ctx)
}
