package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetCommits testa o endpoint /metrics/:user/:repo
func TestGetCommitsRoute(t *testing.T) {
	// 1. Confi the router
	router := SetupRouter()

	// 2. Create a Recorder (record the response to later)
	w := httptest.NewRecorder()

	// 3. Create a mock requisition
	req, _ := http.NewRequest("GET", "/metrics/usuario-teste/repo-teste", nil)

	// 4. Ex router requisition
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusBadRequest {
		t.Errorf("Esperava status 200 ou 400, recebeu %d", w.Code)
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Não foi possível fazer o parse do JSON de resposta: %v", err)
	}
}
