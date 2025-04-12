package restapi

import (
	"encoding/json"
	"errors"
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
	gameJSON := `{"Id":"1","Name":"Test game","Settings":{"BoardWidth":1,"BoardHeight":1},"Gameplay":{"PlayerMoves":[{"Row":1,"Col":1}],"BoardGridPoints":[0]}}`
	_, err = createTestGame(gameJSON, &handler)
	assert.NoError(t, err)
	// Call the endpoint again, expecting the posted item
	req, err = http.NewRequest("GET", "/games", nil)
	assert.NoError(t, err)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	// Remove the ID fields before comparison
	gameJSON, _ = rmID(gameJSON)
	resJSON, err := rmID(rr.Body.String())
	assert.NoError(t, err)
	// Wrap in array before comparison
	gameJSON = `[` + gameJSON + `]`
	assert.JSONEq(t, gameJSON, resJSON)
}

func TestGetGameByID(t *testing.T) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	assert.NoError(t, err)
	api := operations.NewNovelgoAPI(swaggerSpec)
	server := NewServer(api)
	defer server.Shutdown()

	// Test getting non-existent item
	handler := configureAPI(api)
	req, err := http.NewRequest("GET", "/games/1", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)

	// Test getting existing
	// Post 1 item to server
	gameJSON := `{"Id":"1","Name":"Test game","Settings":{"BoardWidth":1,"BoardHeight":1},"Gameplay":{"PlayerMoves":[{"Row":1,"Col":1}],"BoardGridPoints":[0]}}`
	gameID, err := createTestGame(gameJSON, &handler)
	assert.NoError(t, err)
	// Call the endpoint again, getting the posted item
	req, err = http.NewRequest("GET", "/games/"+gameID, nil)
	assert.NoError(t, err)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	// Remove the ID fields before comparison
	gameJSON, _ = rmID(gameJSON)
	resJSON, err := rmID(rr.Body.String())
	assert.NoError(t, err)
	// Wrap in array before comparison
	assert.JSONEq(t, gameJSON, resJSON)
}

func TestCreateGame(t *testing.T) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	assert.NoError(t, err)
	api := operations.NewNovelgoAPI(swaggerSpec)
	server := NewServer(api)
	defer server.Shutdown()

	handler := configureAPI(api)
	gameJSON := `{"Id":"1","Name":"Test game","Settings":{"BoardWidth":1,"BoardHeight":1},"Gameplay":{"PlayerMoves":[{"Row":1,"Col":1}],"BoardGridPoints":[0]}}`
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

func TestPutGame(t *testing.T) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	assert.NoError(t, err)
	api := operations.NewNovelgoAPI(swaggerSpec)
	server := NewServer(api)
	defer server.Shutdown()

	// Update non-existent game
	// Expect error
	handler := configureAPI(api)
	gameJSON := `{"Id":"1","Name":"Test game","Settings":{"BoardWidth":2,"BoardHeight":1},"Gameplay":{"PlayerMoves":[{"Row":0,"Col":0}]}}`
	req, err := http.NewRequest("PUT", "/games/1", strings.NewReader(gameJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)

	// Create a game for update
	gameID, err := createTestGame(gameJSON, &handler)
	assert.NoError(t, err)

	// Update the game by appending to player moves
	// Expect success
	updatedGameJSON := `{"Id":"1","Name":"Test game","Settings":{"BoardWidth":2,"BoardHeight":1},"Gameplay":{"PlayerMoves":[{"Row":0,"Col":0},{"Row":0,"Col":1}],"BoardGridPoints":[1,3]}}`
	req, err = http.NewRequest("PUT", "/games/"+gameID, strings.NewReader(updatedGameJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	// Remove the ID fields before comparison
	updatedGameJSON, _ = rmID(updatedGameJSON)
	resJSON, err := rmID(rr.Body.String())
	assert.NoError(t, err)
	assert.JSONEq(t, updatedGameJSON, resJSON)
}

func TestDeleteGame(t *testing.T) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	assert.NoError(t, err)
	api := operations.NewNovelgoAPI(swaggerSpec)
	server := NewServer(api)
	defer server.Shutdown()

	// Test deleting non-existent item
	handler := configureAPI(api)
	req, err := http.NewRequest("DELETE", "/games/1", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)

	// Test deleting existing
	// Post 1 item to server
	gameJSON := `{"Id":"1","Name":"Test game","Settings":{"BoardWidth":1,"BoardHeight":1},"Gameplay":{"PlayerMoves":[{"Row":1,"Col":1}],"BoardGridPoints":[0]}}`
	gameID, err := createTestGame(gameJSON, &handler)
	assert.NoError(t, err)
	// Call the endpoint again, deleting the posted item
	req, err = http.NewRequest("DELETE", "/games/"+gameID, nil)
	assert.NoError(t, err)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

// Helper function to remove ID fields from JSON strings
//
// Parameters:
// - j: Serialised string of the game to remove ID, in JSON format.
//
// Returns:
// - string: Serialised string of the game, with ID field removed.
// - error: Error if any.
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
			return "", errors.New("failed to unmarshal to single game object or array of game objects")
		}
		for i := range games {
			games[i].ID = &empty
		}
		s, err := json.Marshal(games)
		return string(s), nil
	}
	// Single object marshal succeeded
	game.ID = &empty
	s, err := json.Marshal(game)
	return string(s), nil
}

// Helper function to create a test game on the server when testing other operations
//
// Parameters:
// - gameJSON: Serialised string of the game to create, in JSON format.
// - h: The http handler to create the new game on.
//
// Returns:
// - string: The ID of the new game returned from server.
// - error: Error if any.
func createTestGame(gameJSON string, h *http.Handler) (string, error) {
	req, err := http.NewRequest("POST", "/games", strings.NewReader(gameJSON))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	(*h).ServeHTTP(rr, req)
	if http.StatusCreated != rr.Code {
		return "", errors.New("failed to create test game, server did not return success")
	}
	// Get the ID of the posted item
	var game models.Game
	err = json.Unmarshal([]byte(rr.Body.String()), &game)
	if err != nil {
		return "", err
	}
	return *game.ID, nil
}
