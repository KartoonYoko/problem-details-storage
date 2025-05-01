package server

import (
	"context"
	"fmt"
	"github.com/KartoonYoko/problem-details-storage/internal/controller/http_server"
	"github.com/KartoonYoko/problem-details-storage/internal/logger"
	"github.com/KartoonYoko/problem-details-storage/internal/storage"
	"github.com/KartoonYoko/problem-details-storage/internal/usecase/buckets"
	"go.uber.org/zap"
	"log"
)

func Run() {
	ctx := context.Background()

	if err := logger.Initialize("Info"); err != nil {
		log.Fatal(fmt.Errorf("logger init error: %w", err))
	}

	defer func(Log *zap.Logger) {
		err := Log.Sync()
		if err != nil {
			log.Fatal(fmt.Errorf("logger sync error: %w", err))
		}
	}(logger.Log)

	controller := http_server.New(http_server.Config{BaseAddress: ":8080"})

	err := configUsecases(ctx, controller)

	if err != nil {
		log.Fatal(fmt.Errorf("config usecase error: %w", err))
	}

	err = controller.Serve(ctx)

	if err != nil {
		log.Fatal(fmt.Errorf("server serving error: %w", err))
	}
}

func configUsecases(
	ctx context.Context,
	controller *http_server.Controller) error {
	s, err := storage.New(ctx, storage.Config{
		ConnectionString: "host=localhost user=problem_details_storage password=problem_details_storage dbname=problem_details_storage_db port=11000 sslmode=disable",
	})

	if err != nil {
		return fmt.Errorf("storage initilisation error: %w", err)
	}

	buckets.New(s)

	return nil
}
