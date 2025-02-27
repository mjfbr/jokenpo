package utils

import (
	"fmt"
	"os"
	"time"
)

func LogToFile(playerChoice, computerChoice, result, userIP string) {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo de log:", err)
		return
	}
	defer file.Close()

	logEntry := fmt.Sprintf("[%s] IP: %s | Jogador: %s | Computador: %s | Resultado: %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		userIP,
		playerChoice,
		computerChoice,
		result,
	)

	if _, err := file.WriteString(logEntry); err != nil {
		fmt.Println("Erro ao escrever no log:", err)
	}
}
