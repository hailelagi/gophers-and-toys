// Package rps contains simple game logic and data structures for a game of rock-paper-scissors.
// Rules:
// Rock -> scissors.
// Scissors -> paper.
// Paper -> rock.
package rps

import (
	"errors"
	"strings"
)

type gameState struct {
	clientA string
	clientB string
}

// sanitiseInput properly formats input
func sanitiseInput(input string) (string, error) {
	var formatInput = strings.TrimSpace(strings.ToLower(input))
	var hands = [3]string{"rock", "paper", "scissors"}

	for _, hand := range hands {
		if hand == formatInput {
			return input, nil
		}
	}
	return "", errors.New("invalid hand supplied")
}

// validateRules checks for the winner of a game
func validateRules(state gameState) string {
	if state.clientA == state.clientB {
		return "draw"
	}

	switch state.clientA {
	case "rock":
		if state.clientB == "scissors" {
			return "client A wins"
		}
	case "paper":
		if state.clientB == "rock" {
			return "client A wins"
		}
	case "scissors":
		if state.clientB == "paper" {
			return "client A wins"
		}
	}
	return "client B wins"
}
