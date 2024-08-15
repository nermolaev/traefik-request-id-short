// Package traefik_plugin_request_id_short a Traefik plugin to add request ID to incoming HTTP requests.
package traefik_plugin_request_id_short

import (
  "context"
  "github.com/google/uuid"
  "net/http"
  "strings"
)

const defaultHeader = "X-Request-ID"
const defaultEnabled = true

type Config struct {
  HeaderName string `json:"headerName,omitempty"`
  Enabled    bool   `json:"enabled,omitempty"`
}

func CreateConfig() *Config {
  return &Config{
    HeaderName: defaultHeader,
    Enabled:    defaultEnabled,
  }
}

func New(ctx context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
  return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
    if config.Enabled {
      existingID := request.Header.Get(config.HeaderName)
      if existingID == "" {
        value := strings.ReplaceAll(uuid.Must(uuid.NewRandom()).String(), "-", "")
        request.Header.Add(config.HeaderName, value)
        writer.Header().Add(config.HeaderName, value)
      } else {
        writer.Header().Add(config.HeaderName, existingID)
      }
    }
    next.ServeHTTP(writer, request)
  }), nil
}
