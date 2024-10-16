package handlers

import (
	"errors"
	"novelgo/internal/pkg/models"

	"github.com/google/uuid"
)

// Use a map to simulate a DB
// TODO: Replace with an actual DB
var games = make(map[string]*models.Game)

// ListGames returns all games
func ListGames() []*models.Game {
	var result []*models.Game
	for _, game := range games {
		result = append(result, game)
	}
	return result
}

// CreateGame adds a new game
// The ID field from the request is ignored
// A new uuid will be generated for the new game
func CreateGame(game *models.Game) (*models.Game, error) {
	ID := uuid.New().String()
	game.ID = &ID
	games[*game.ID] = game
	return game, nil
}

// GetGameByID returns a game by ID
func GetGameByID(id string) (*models.Game, error) {
	game, exists := games[id]
	if !exists {
		return nil, errors.New("game not found")
	}
	return game, nil
}

// UpdateGame updates an existing game
func UpdateGame(id string, game *models.Game) error {
	_, exists := games[id]
	if !exists {
		return errors.New("game not found")
	}
	games[id] = game
	return nil
}

// DeleteGame removes a game by ID
func DeleteGame(id string) error {
	_, exists := games[id]
	if !exists {
		return errors.New("game not found")
	}
	delete(games, id)
	return nil
}
