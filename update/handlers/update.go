package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/rootxrishabh/PasteContainer/pkg/common/models"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func (h *Handler) UpdatePaste(c *gin.Context) {
	p := &models.Paste{}
	json.NewDecoder(c.Request.Body).Decode(p)
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	p.ID = id
	h.HandlerEnv.DB.Model(p).Updates(p)
	c.String(http.StatusOK, "%s", p.Body)
}