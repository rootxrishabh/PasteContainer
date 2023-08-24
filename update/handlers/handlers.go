package handlers

import (
	"github.com/rootxrishabh/PasteContainer/pkg/common/models"
	"gorm.io/gorm"
)

type Handler struct {
	HandlerEnv *models.HandlerEnv
}

func NewHandler(db *gorm.DB) Handler {
	hEnv := models.HandlerEnv{
		DB: db,
	}

	return Handler{
		HandlerEnv: &hEnv,
	}
}