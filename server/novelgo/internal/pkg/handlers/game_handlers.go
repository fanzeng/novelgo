package handlers

import (
	"errors"
	"fmt"
	"novelgo/internal/pkg/cyclic"
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

	b := cyclic.NewBoard(int(game.Settings.BoardHeight), int(game.Settings.BoardWidth), true)
	arr := b.GetGridPointsAsArray()
	game.Gameplay.BoardGridPoints = make([]int64, len(arr))
	for i, s := range arr {
		game.Gameplay.BoardGridPoints[i] = int64(s)
	}
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
func UpdateGame(id string, game *models.Game) (*models.Game, error) {
	old, exists := games[id]
	if !exists {
		return nil, errors.New("game not found")
	}
	oldMoves := &old.Gameplay.PlayerMoves
	moves := game.Gameplay.PlayerMoves
	lastMove := moves[len(moves)-1]
	fmt.Printf("last move = %v", lastMove)
	*oldMoves = append(*oldMoves, lastMove)
	fmt.Printf("oldMoves = %v", *oldMoves)
	b := cyclic.NewBoard(int(game.Settings.BoardHeight), int(game.Settings.BoardWidth), true)
	for i, move := range *oldMoves {
		fmt.Printf("oves = %v", move)
		var color cyclic.GridPointState
		if i % 2 == 0 {
			color = cyclic.Black
		} else {
			color = cyclic.White
		}
		b.Put(int(*move.Row), int(*move.Col), color)
	}
	arr := b.GetGridPointsAsArray()
	fmt.Printf("arr = %v", arr)
	old.Gameplay.BoardGridPoints = make([]int64, len(arr))
	for i, s := range arr {
		old.Gameplay.BoardGridPoints[i] = int64(s)
	}
	games[id] = old
	return old, nil
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
