package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/darksford123x/app/ent"
	"github.com/darksford123x/app/ent/repairinvoice"
	"github.com/gin-gonic/gin"
)

// RepairInvoiceController defines the struct for the RepairInvoice controller
type RepairInvoiceController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateRepairInvoice handles POST requests for adding RepairInvoice entities
// @Summary Create RepairInvoice
// @Description Create RepairInvoice
// @ID create-RepairInvoice
// @Accept   json
// @Produce  json
// @Param RepairInvoice body ent.RepairInvoice true "RepairInvoice entity"
// @Success 200 {object} ent.RepairInvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /RepairInvoice [post]
func (ctl *RepairInvoiceController) CreateRepairInvoice(c *gin.Context) {
	obj := ent.RepairInvoice{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "RepairInvoice binding failed",
		})
		return
	}

	u, err := ctl.client.RepairInvoice.
		Create().
		SetRepairInvoceId(obj.Age).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetRepairInvoice handles GET requests to retrieve a repairinvoce entity
// @Summary Get a repairinvoce entity by ID
// @Description get repairinvoce by ID
// @ID get-repairinvoce
// @Produce  json
// @Param id path int true "RepairInvoice ID"
// @Success 200 {object} ent.RepairInvoice
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /RepairInvoices/{id} [get]
func (ctl *RepairInvoiceController) GetRepairInvoice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.RepairInvoice.
		Query().
		Where(RepairInvoice.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListRepairInvoice handles request to get a list of RepairInvoice entities
// @Summary List RepairInvoice entities
// @Description list RepairInvoice entities
// @ID list-RepairInvoice
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.RepairInvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /RepairInvoices [get]
func (ctl *RepairInvoiceController) ListRepairInvoice(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	RepairInvoices, err := ctl.client.RepairInvoice.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, RepairInvoices)
}

// DeleteRepairInvoice handles DELETE requests to delete a RepairInvoice entity
// @Summary Delete a RepairInvoice entity by ID
// @Description get RepairInvoice by ID
// @ID delete-RepairInvoice
// @Produce  json
// @Param id path int true "RepairInvoice ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /RepairInvoices/{id} [delete]
func (ctl *RepairInvoiceController) DeleteRepairInvoice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.RepairInvoice.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateRepairInvoice handles PUT requests to update a RepairInvoice entity
// @Summary Update a RepairInvoice entity by ID
// @Description update RepairInvoice by ID
// @ID update-RepairInvoice
// @Accept   json
// @Produce  json
// @Param id path int true "RepairInvoice ID"
// @Param RepairInvoice body ent.RepairInvoice true "RepairInvoice entity"
// @Success 200 {object} ent.RepairInvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /RepairInvoices/{id} [put]
func (ctl *RepairInvoiceController) UpdateRepairInvoice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.RepairInvoice{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "RepairInvoice binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.RepairInvoice.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewRepairInvoiceController creates and registers handles for the RepairInvoice controller
func NewRepairInvoiceController(router gin.IRouter, client *ent.Client) *RepairInvoiceController {
	uc := &RepairInvoiceController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRepairInvoiceController registers routes to the main engine
func (ctl *RepairInvoiceController) register() {
	RepairInvoices := ctl.router.Group("/RepairInvoices")

	RepairInvoices.GET("", ctl.ListRepairInvoice)

	// CRUD
	RepairInvoices.POST("", ctl.CreateRepairInvoice)
	RepairInvoices.GET(":id", ctl.GetRepairInvoice)
	RepairInvoices.PUT(":id", ctl.UpdateRepairInvoice)
	RepairInvoices.DELETE(":id", ctl.DeleteRepairInvoice)
}
