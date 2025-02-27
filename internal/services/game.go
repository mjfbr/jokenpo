package services

import (
	"errors"
	"jokenpo/pkg/utils"
)

var choices = []string{"pedra", "papel", "tesoura"}

func IsValidChoice(choice string) bool {
	for _, validChoice := range choices {
		if choice == validChoice {
			return true
		}
	}
	return false
}

func GetComputerChoice() string {
	return utils.RandomChoice(choices)
}

func DetermineWinner(playerChoice, computerChoice string) (string, error) {
	if !IsValidChoice(playerChoice) {
		return "", errors.New("escolha inválida, use 'pedra', 'papel' ou 'tesoura'")
	}

	if playerChoice == computerChoice {
		return "Empate!", nil
	}

	winConditions := map[string]string{
		"pedra":   "tesoura",
		"papel":   "pedra",
		"tesoura": "papel",
	}

	if winConditions[playerChoice] == computerChoice {
		return "Você venceu!", nil
	}

	return "Computador venceu!", nil
}
