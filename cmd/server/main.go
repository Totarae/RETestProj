package main

import (
	"awesomeProject10/internal/config"
	"awesomeProject10/internal/middleware"
	"awesomeProject10/internal/router"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to init zap logger: %v", err)
	}
	defer logger.Sync()

	conf := &config.Config{}

	if err := conf.LoadPackSizesFromFile("internal/config/packs.json"); err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	if err := conf.WatchConfigFile("internal/config/packs.json"); err != nil {
		logger.Fatal("failed to watch config", zap.Error(err))
	}

	// logging routeas
	r := middleware.LoggingMiddleware(logger)(router.NewRouter(conf))

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("server error: %v", err)
	}

	// TODO: graceful shutdown, condigurable port, etc.
}
