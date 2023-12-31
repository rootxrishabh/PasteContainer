package router

import (
	"github.com/rootxrishabh/PasteContainer/create/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine{
	r := gin.Default()
	h := handlers.NewHandler(db)
	r.POST("/create", h.CreatePaste)
	return r
}
