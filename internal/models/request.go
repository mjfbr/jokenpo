package models

type PlayRequest struct {
	Choice string `json:"tipo_jogada" binding:"required"`
}
