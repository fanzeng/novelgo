package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"novelgo/internal/pkg/models"
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
	// Post 1 item to server
	gameJSON := `{"Id":"1","Name":"Test game","Settings":{"BoardWidth":10,"BoardHeight":10},"Gameplay":{"PlayerMoves":[{"Row":1,"Col":1}]}}`
	req, err = http.NewRequest("POST", "/games", strings.NewReader(gameJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	// Call the endpoint again, expecting the posted item
	req, err = http.NewRequest("GET", "/games", nil)
	assert.NoError(t, err)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	// Remove the ID fields before comparison
	gameJSON, _ = rmID(gameJSON)
	resJSON, err := rmID(rr.Body.String())
	fmt.Println(rr.Body.String())
	assert.NoError(t, err)
	// Wrap in array before comparison
	gameJSON = `[` + gameJSON + `]`
	assert.JSONEq(t, gameJSON, resJSON)
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
	// Remove the ID fields before comparison
	gameJSON, _ = rmID(gameJSON)
	resJSON, err := rmID(rr.Body.String())
	assert.NoError(t, err)
	assert.JSONEq(t, gameJSON, resJSON)
}

// Helper function to remove ID fields from JSON strings
func rmID(j string) (string, error) {
	empty := ""
	// Try unmarhsal into single object first
	var game models.Game
	err := json.Unmarshal([]byte(j), &game)
	if err != nil {
		// Check for case of array
		var games []models.Game
		err := json.Unmarshal([]byte(j), &games)
		if err != nil {
			return "", err
		}
		for i := range games {
			games[i].ID = &empty
		}
		s, err := json.Marshal(games)
		return string(s), nil
	}
	// Single object marshal succeded
	game.ID = &empty
	s, err := json.Marshal(game)
	return string(s), nil
}
