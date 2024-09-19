package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	constants "github.com/turk2003/workflow/constant"
	"github.com/turk2003/workflow/models"
	"github.com/turk2003/workflow/services"
)

type ItemController struct {
	Service *services.ItemService
}

func NewItemController(service *services.ItemService) *ItemController {
	return &ItemController{Service: service}
}

func (c *ItemController) CreateItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if item.Status == "" {
		item.Status = constants.ItemPendingStatus
	}
	if err := c.Service.CreateItem(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, item)
}

func (c *ItemController) GetAllItems(ctx *gin.Context) {
	items, err := c.Service.GetAllItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *ItemController) GetItemByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	item, err := c.Service.GetItemByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

// Implement other methods (Update, Delete, etc.) here...
func (c *ItemController) UpdateItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = uint(id)

	if err := c.Service.UpdateItem(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *ItemController) PatchItemStatus(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var updateStatus struct {
		Status string `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&updateStatus); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่า status ที่รับเข้ามาตรงกับค่าที่กำหนดไว้ใน constants หรือไม่
	if updateStatus.Status != string(constants.ItemPendingStatus) &&
		updateStatus.Status != string(constants.ItemApprovedStatus) &&
		updateStatus.Status != string(constants.ItemRejectedStatus) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	item, err := c.Service.GetItemByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// แปลงค่า string เป็น constants.ItemStatus
	item.Status = constants.ItemStatus(updateStatus.Status)
	if err := c.Service.UpdateItem(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (c *ItemController) DeleteItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.Service.DeleteItem(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
