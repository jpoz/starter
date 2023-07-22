package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/jpoz/starter/pkg/config"
	"github.com/jpoz/starter/pkg/server"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.SetLevel(log.ParseLevel(os.Getenv("LOG_LEVEL")))

	// Gorm
	newLogger := logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	gormdb, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("error opening db: %w", err))
	}
	sqlDB, err := gormdb.DB()
	if err != nil {
		log.Fatal(fmt.Errorf("error getting db: %w", err))
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(fmt.Errorf("error pinging db: %w", err))
	}

	// Redis
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(fmt.Errorf("error getting redis url: %w", err))
	}
	r := redis.NewClient(opt)
	_, err = r.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	// Config
	cfg := config.Server{
		Addr:      os.Getenv("ADDR"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	log.Info("Starting server...")

	s := server.NewServer(cfg)
	go s.ListenAndServe()

	// Create a channel to listen for the termination signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for the termination signal
	sig := <-signalChan
	log.Warnf("Received signal: %s. Shutting down...", sig.String())

	log.Info("Goodbye!")
}
