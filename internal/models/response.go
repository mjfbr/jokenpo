package models

type PlayResponse struct {
	PlayerChoice   string `json:"escolha_usuario"`
	ComputerChoice string `json:"escolha_computador"`
	Result         string `json:"resultado"`
}
