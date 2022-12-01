package helpers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mahendrafathan/go-cake/common/logger"
)

type AppError struct {
	Error   error
	Message string
	Code    int
}

func RespondWithError(ctx context.Context, w http.ResponseWriter, code int, message string) {
	logger.Errorf(ctx, "Got Error : %+v", message)
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
