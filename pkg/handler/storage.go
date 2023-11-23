package handler

import (
	"net/http"

	"github.com/GRRashid/lamoda"
	"github.com/gin-gonic/gin"
)

// @Summary Create storage
// @Tags storage
// @Description create storage
// @ID create-storage
// @Accept  json
// @Produce  json
// @Param storage body lamoda.RawStorage true "storage info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
func (h *Handler) createStorage(c *gin.Context) {
	var storage lamoda.RawStorage
	err := c.BindJSON(&storage)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateStorage(storage)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
