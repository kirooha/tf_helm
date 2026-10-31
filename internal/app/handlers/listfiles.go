package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/kirooha/kuber-practice/internal/pkg/dbmodel"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type ListHandler struct {
	queries     *dbmodel.Queries
	redisClient *redis.Client
	apiKey      string
}

func NewListHandler(queries *dbmodel.Queries, redisClient *redis.Client, apiKey string) *ListHandler {
	return &ListHandler{
		queries:     queries,
		redisClient: redisClient,
		apiKey:      apiKey,
	}
}

func (h *ListHandler) Handle(fiberCtx *fiber.Ctx) error {
	var (
		ctx          = fiberCtx.Context()
		msgPrefix    = "app.handler.ListHandler.Handle"
		headerApiKey = fiberCtx.Get("Authorization")
	)

	if h.apiKey != headerApiKey {
		return fiberCtx.SendStatus(http.StatusForbidden)
	}

	value, err := h.redisClient.Get(ctx, "filenames").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Printf("%s: h.redisClient.Get error - %v", msgPrefix, err)
		return fiberCtx.SendStatus(http.StatusInternalServerError)
	}
	filenames := strings.Split(value, ",")

	return json.NewEncoder(fiberCtx.Response().BodyWriter()).Encode(filenames)
}
