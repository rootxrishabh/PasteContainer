package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rootxrishabh/PasteContainer/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// create a new paste and writes it to database
func (h *Handler) CreatePaste(c *gin.Context){
p := &models.Paste{}
json.NewDecoder(c.Request.Body).Decode(p)
h.HandlerEnv.DB.Create(p)
c.Redirect(http.StatusSeeOther, "/" + p.ID.String())
}