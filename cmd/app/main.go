package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kirooha/kuber-practice/internal/app/handlers"
	"github.com/kirooha/kuber-practice/internal/pkg/dbmodel"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if os.Getenv("APP_PORT") == "" {
		log.Fatal("APP_PORT must be set")
	}
	if os.Getenv("APP_API_KEY") == "" {
		log.Fatal("APP_API_KEY must be set")
	}
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
	if os.Getenv("APP_DB_MIGRATIONS_DIRECTORY") == "" {
		log.Fatal("APP_DB_MIGRATIONS_DIRECTORY must be set")
	}
	if os.Getenv("APP_REDIS_HOST") == "" {
		log.Fatal("APP_REDIS_HOST must be set")
	}
	if os.Getenv("APP_REDIS_PASSWORD") == "" {
		log.Fatal("APP_REDIS_PASSWORD must be set")
	}

	var (
		dbUser             = os.Getenv("APP_DB_USER")
		dbPassword         = os.Getenv("APP_DB_PASSWORD")
		dbHost             = os.Getenv("APP_DB_HOST")
		dbPort             = os.Getenv("APP_DB_PORT")
		dbName             = os.Getenv("APP_DB_NAME")
		dbMigrationsFolder = os.Getenv("APP_DB_MIGRATIONS_DIRECTORY")

		redisHost     = os.Getenv("APP_REDIS_HOST")
		redisPassword = os.Getenv("APP_REDIS_PASSWORD")

		apiKey = os.Getenv("APP_API_KEY")

		port = os.Getenv("APP_PORT")
	)

	conn, err := pgx.Connect(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("pgx.Connect error - %v", err)
	}
	defer conn.Close(ctx)

	pingCtx, pingCancel := context.WithTimeout(ctx, 1*time.Second)
	defer pingCancel()

	if err := conn.Ping(pingCtx); err != nil {
		log.Fatalf("conn.Ping error - %v", err)
	}

	runMigrations(ctx, dbUser, dbPassword, dbHost, dbPort, dbName, dbMigrationsFolder)

	queries := dbmodel.New(conn)

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     redisHost,
			Password: redisPassword,
			DB:       0,
		},
	)
	defer redisClient.Close()

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("redisClient.Ping().Result() error - %v", err)
	}

	app := fiber.New()

	app.Get("/files", handlers.NewListHandler(queries, redisClient, apiKey).Handle)
	app.Get("/healthcheck", handlers.NewHealthcheckHandler().Handle)
	app.Post("/file", handlers.NewSaveHandler(queries, apiKey).Handle)
	app.Post("/foo", handlers.NewFooHandler().Handle)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

func runMigrations(ctx context.Context, dbUser, dbPassword, dbHost, dbPort, dbName, dbMigrationsFolder string) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("sql.Open error - %v", err)
	}
	defer db.Close()

	pingCtx, pingCancel := context.WithTimeout(ctx, 1*time.Second)
	defer pingCancel()

	if err := db.PingContext(pingCtx); err != nil {
		log.Fatalf("db.Ping error - %v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("goose.SetDialect error - %v", err)
	}
	if err := goose.Up(db, dbMigrationsFolder); err != nil {
		log.Fatalf("goose.Up error - %v", err)
	}
}
