package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rootxrishabh/PasteContainer/update/handlers"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := handlers.NewHandler(db)
	r.POST("/update/:id", h.UpdatePaste)

	return r
}