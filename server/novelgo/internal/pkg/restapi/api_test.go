package restapi

import (
	"net/http"
	"net/http/httptest"
	"novelgo/internal/pkg/restapi/operations"
	"strings"
	"testing"

	"github.com/go-openapi/loads"
	"github.com/stretchr/testify/assert"
)

func TestListGames(t *testing.T) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	assert.NoError(t, err)
	api := operations.NewNovelgoAPI(swaggerSpec)
	server := NewServer(api)
	defer server.Shutdown()

	// Test empty
	handler := configureAPI(api)
	req, err := http.NewRequest("GET", "/games", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	expected := `[]`
	assert.JSONEq(t, expected, rr.Body.String())

	// Test non-empty
	gameJSON := `{"Id":"1","Name":"Test game","Settings":{"BoardWidth":10,"BoardHeight":10},"Gameplay":{"PlayerMoves":[{"Row":1,"Col":1}]}}`
	req, err = http.NewRequest("POST", "/games", strings.NewReader(gameJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	req, err = http.NewRequest("GET", "/games", nil)
	assert.NoError(t, err)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `[`+gameJSON+`]`, rr.Body.String())
}

func TestCreateGame(t *testing.T) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	assert.NoError(t, err)
	api := operations.NewNovelgoAPI(swaggerSpec)
	server := NewServer(api)
	defer server.Shutdown()

	handler := configureAPI(api)
	gameJSON := `{"Id":"1","Name":"Test game","Settings":{"BoardWidth":10,"BoardHeight":10},"Gameplay":{"PlayerMoves":[{"Row":1,"Col":1}]}}`
	req, err := http.NewRequest("POST", "/games", strings.NewReader(gameJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.JSONEq(t, gameJSON, rr.Body.String())
}
