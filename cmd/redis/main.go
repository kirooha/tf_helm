package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kirooha/kuber-practice/internal/pkg/dbmodel"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

func main() {
	var (
		ctx = context.Background()
	)

	if os.Getenv("APP_DB_USER") == "" {
		log.Fatal("APP_DB_USER must be set")
	}
	if os.Getenv("APP_DB_PASSWORD") == "" {
		log.Fatal("APP_DB_PASSWORD must be set")
	}
	if os.Getenv("APP_DB_HOST") == "" {
		log.Fatal("APP_DB_HOST must be set")
	}
	if os.Getenv("APP_DB_PORT") == "" {
		log.Fatal("APP_DB_PORT must be set")
	}
	if os.Getenv("APP_DB_NAME") == "" {
		log.Fatal("APP_DB_NAME must be set")
	}
	if os.Getenv("APP_REDIS_HOST") == "" {
		log.Fatal("APP_REDIS_HOST must be set")
	}
	if os.Getenv("APP_REDIS_PASSWORD") == "" {
		log.Fatal("APP_REDIS_PASSWORD must be set")
	}

	var (
		dbUser     = os.Getenv("APP_DB_USER")
		dbPassword = os.Getenv("APP_DB_PASSWORD")
		dbHost     = os.Getenv("APP_DB_HOST")
		dbPort     = os.Getenv("APP_DB_PORT")
		dbName     = os.Getenv("APP_DB_NAME")

		redisHost     = os.Getenv("APP_REDIS_HOST")
		redisPassword = os.Getenv("APP_REDIS_PASSWORD")
	)

	conn, err := pgx.Connect(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("pgx.Connect error - %v", err)
	}
	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("conn.Ping error - %v", err)
	}

	queries := dbmodel.New(conn)

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     redisHost,
			Password: redisPassword,
			DB:       0,
		},
	)
	defer redisClient.Close()

	files, err := queries.ListFiles(ctx)
	if err != nil {
		log.Fatalf("queries.ListFiles error - %v", err)
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name)
	}

	err = redisClient.Set(ctx, "filenames", strings.Join(filenames, ","), time.Hour).Err()
	if err != nil {
		log.Fatalf("redisClient.Set error - %v", err)
	}

	log.Println("filenames successfully set")
}
