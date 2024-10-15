package handlers

import (
	"errors"
	"novelgo/internal/pkg/models"
)

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
func CreateGame(game *models.Game) (*models.Game, error) {
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
