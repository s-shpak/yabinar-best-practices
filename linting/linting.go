package linting

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.uber.org/zap"
)

type Response struct {
	//Text2     string
	Text      string
	IsTextSet bool
	Success   bool
}

// DoSmth does something.
func DoSmth() error {
	resp, err := http.Get("localhost:8080/v1/some_op")
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}
	defer func() {
		resp.Body.Close()
	}()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	var req Response

	if err := json.Unmarshal(b, &req); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	log.Printf("text: %s (%t)", req.Text, req.IsTextSet)
	return nil
}

func LogHello() {
	zap.S().Info("hi!")
}
