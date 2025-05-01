package http_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/KartoonYoko/problem-details-storage/internal/logger"
	"github.com/KartoonYoko/problem-details-storage/internal/usecase/buckets"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Controller struct {
	config Config

	router *chi.Mux

	usecase *buckets.Usecase
}

func New(config Config, usecase *buckets.Usecase) *Controller {
	c := new(Controller)
	c.config = config
	c.usecase = usecase

	router := chi.NewRouter()

	router.Use(middlewareLogRequestDuration)
	router.Use(middlewareLogRequestInfo)
	router.Use(middlewareLogResponseInfo)

	router.HandleFunc("/*", handleWebApp)

	apiRouter := chi.NewRouter()
	router.Mount("/api/", apiRouter)
	apiRouter.Post("/bucket", c.handlePOSTBucket)
	apiRouter.Get("/buckets", c.handleGETBuckets)
	apiRouter.Get("/buckets/{bucketID}/problem-details", c.handleGETBucketProblemDetails)
	apiRouter.Get("/buckets/{bucketID}/problem-detail/{problemDetailType}", c.handleGETBucketProblemDetailByType)
	apiRouter.Post("/buckets/{bucketID}/problem-detail", c.handlePOSTBucketProblemDetail)
	apiRouter.Put("/buckets/{bucketID}/problem-detail/{problemDetailID}", c.handlePUTBucketProblemDetail)
	apiRouter.Delete("/buckets/{bucketID}/problem-detail/{problemDetailID}", c.handleDELETEBucketProblemDetail)

	c.router = router

	return c
}

func (c *Controller) Serve(ctx context.Context) error {
	server := &http.Server{Addr: c.config.BaseAddress, Handler: c.router}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(ctx)

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// run server
	logger.Log.Info(fmt.Sprintf("server serve on %s", server.Addr))
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server serve error: %w", err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

	return nil
}
