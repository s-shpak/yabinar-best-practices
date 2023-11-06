package errors

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if err := handleRequest(r); err != nil {
		zap.S().Error("failed to handle users request", zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func handleRequest(r *http.Request) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	zap.S().Infof("received body: %s", b)
	return nil
}
