package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GRRashid/lamoda"
	"github.com/gin-gonic/gin"
)

// @Summary Create product
// @Tags products
// @Description create product
// @ID create-product
// @Accept  json
// @Produce  json
// @Param product body lamoda.RawProduct true "product info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/products/create [post]
func (h *Handler) createProduct(c *gin.Context) {
	var product lamoda.RawProduct
	err := c.BindJSON(&product)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(product)
	if err != nil {
		fmt.Print(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Reserve product
// @Tags products
// @Description reserve product
// @ID reserved-product
// @Accept  json
// @Produce  json
// @Param ids body lamoda.ProductIds true "product info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/products/reserve [put]
func (h *Handler) reserveProduct(c *gin.Context) {
	var ids lamoda.ProductIds
	err := c.BindJSON(&ids)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.ReservedProduct(ids.IDs)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})

}

// @Summary Unreserved product
// @Tags products
// @Description unreserved product
// @ID unreserved-product
// @Accept  json
// @Produce  json
// @Param ids body lamoda.ProductIds true "product info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/products/unreserved [put]
func (h *Handler) unreservedProduct(c *gin.Context) {
	var ids lamoda.ProductIds
	err := c.BindJSON(&ids)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UnreservedProduct(ids.IDs)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

type getAvailableResponse struct {
	Data []lamoda.Product `json:"data"`
}

// @Summary Get available products
// @Tags products
// @Description get all lists
// @ID get-products
// @Accept json
// @Produce json
// @Param storageId path int true "storage id"
// @Success 200 {object} getAvailableResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/storages/{storageId}/products/unreserved [get]
func (h *Handler) getAvailableProducts(c *gin.Context) {
	storageId, err := strconv.Atoi(c.Param("storageId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	products, err := h.services.GetLast(storageId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAvailableResponse{
		Data: products,
	})
}
