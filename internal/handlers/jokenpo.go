package handlers

import (
	"jokenpo/internal/models"
	"jokenpo/internal/services"
	"jokenpo/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlayJokenpo(c *gin.Context) {
	var request models.PlayRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida"})
		return
	}

	computerChoice := services.GetComputerChoice()

	result, err := services.DetermineWinner(request.Choice, computerChoice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIP := c.ClientIP()

	utils.LogToFile(request.Choice, computerChoice, result, userIP)

	c.JSON(http.StatusOK, models.PlayResponse{
		PlayerChoice:   request.Choice,
		ComputerChoice: computerChoice,
		Result:         result,
	})
}
