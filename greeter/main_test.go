package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Alice", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var got Greeting
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if got.Message != "Hello, Alice!" {
		t.Errorf("expected greeting to include Alice, got %q", got.Message)
	}
}

func TestHelloWithoutName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var got Greeting
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if got.Message != "Hello, World!" {
		t.Errorf("expected generic greeting, got %q", got.Message)
	}
}

func TestHelloWithEmptyName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var got Greeting
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if got.Message != "Hello, World!" {
		t.Errorf("expected generic greeting, got %q", got.Message)
	}
}
