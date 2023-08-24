package handlers

import (
	"github.com/rootxrishabh/PasteContainer/pkg/common/models"
	"gorm.io/gorm"
)

type Handlers struct {
	HandlerEnv *models.HandlerEnv
}

func NewHandler(db *gorm.DB) Handlers{
	hEnv := models.HandlerEnv{
		DB: db,
	}
	
	return Handler {
		HandlerEnv: &hEnv,
	}
}