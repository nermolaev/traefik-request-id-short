package traefik_plugin_request_id_short

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddRequestIdInHeaderIfNoneExist(t *testing.T) {
	cfg := CreateConfig()
	cfg.HeaderName = "X-Custom-Request-ID" // Настраиваем имя заголовка

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "sw-request-id-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	// Проверка заголовка запроса
	cid := req.Header.Get(cfg.HeaderName)
	if cid == "" {
		t.Errorf("Request ID has not been generated")
	} else if strings.Contains(cid, "-") {
		t.Errorf("Request ID should not contain dashes, got: %s", cid)
	}

	// Проверка заголовка ответа
	responseCid := recorder.Header().Get(cfg.HeaderName)
	if responseCid == "" {
		t.Errorf("Request ID has not been added to the response")
	} else if strings.Contains(responseCid, "-") {
		t.Errorf("Response Request ID should not contain dashes, got: %s", responseCid)
	}
}

func TestKeepRequestIdInHeaderIfOneExist(t *testing.T) {
	cfg := CreateConfig()
	cfg.HeaderName = "X-Custom-Request-ID" // Настраиваем имя заголовка

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "sw-request-id-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	existingCid := "some-existing-request-id"
	req.Header.Set(cfg.HeaderName, existingCid)

	handler.ServeHTTP(recorder, req)

	// Проверка, что существующий Request ID остался в запросе
	cid := req.Header.Get(cfg.HeaderName)
	if cid != existingCid {
		t.Errorf("Existing Request ID has not been kept in the request, got: %s", cid)
	}

	// Проверка, что существующий Request ID остался в ответе
	responseCid := recorder.Header().Get(cfg.HeaderName)
	if responseCid != existingCid {
		t.Errorf("Existing Request ID in the response has not been kept, got: %s", responseCid)
	}
}
