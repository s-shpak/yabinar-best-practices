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
	IsTextSet bool
	Text      string

	Success bool
}

func DoSmth() error {
	resp, err := http.Get("localhost:8080/v1/some_op")
	if err != nil {
		return err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	var req Response

	json.Unmarshal(b, &req)

	log.Printf("text: %s (%t)", req.Text, req.IsTextSet)
	return nil
}

func LogHello() {
	zap.S().Info("hi!")
}
