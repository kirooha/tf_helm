package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/kirooha/kuber-practice/internal/pkg/dbmodel"

	"github.com/gofiber/fiber/v2"
)

type SaveHandler struct {
	queries *dbmodel.Queries
	apiKey  string
}

func NewSaveHandler(queries *dbmodel.Queries, apiKey string) *SaveHandler {
	return &SaveHandler{
		queries: queries,
		apiKey:  apiKey,
	}
}

func (h *SaveHandler) Handle(fiberCtx *fiber.Ctx) error {
	var (
		ctx          = fiberCtx.Context()
		msgPrefix    = "app.handler.SaveHandler.Handle"
		headerApiKey = fiberCtx.Get("Authorization")
	)

	if h.apiKey != headerApiKey {
		return fiberCtx.SendStatus(http.StatusForbidden)
	}

	multipartForm, err := fiberCtx.MultipartForm()
	if err != nil {
		return fiberCtx.SendStatus(http.StatusInternalServerError)
	}

	file, ok := multipartForm.File["file"]
	if !ok {
		return fiberCtx.SendStatus(http.StatusBadRequest)
	}
	if len(file) != 1 {
		return fiberCtx.SendStatus(http.StatusBadRequest)
	}

	var (
		fileHeader = file[0]
		filename   = fileHeader.Filename
	)

	f, err := fileHeader.Open()
	if err != nil {
		log.Printf("%s: fileHeader.Open error - %v", msgPrefix, err)
		return fiberCtx.SendStatus(http.StatusInternalServerError)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Printf("%s: io.ReadAll error - %v", msgPrefix, err)
		return fiberCtx.SendStatus(http.StatusInternalServerError)
	}

	_, err = h.queries.AddFile(ctx, dbmodel.AddFileParams{
		Name:    filename,
		Content: string(b),
	})
	if err != nil {
		log.Printf("%s: h.queries.AddFile error - %v", msgPrefix, err)
		return fiberCtx.SendStatus(http.StatusInternalServerError)
	}

	return fiberCtx.SendString("file saved")
}
