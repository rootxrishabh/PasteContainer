package router

import (
	"github.com/rootxxrishabh/PasteContainer/delete/handlers"
	"github.com/go-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := handlers.NewHandler(db)

	r.DELETE(":/id", h.DeletePaste)
	return r

}