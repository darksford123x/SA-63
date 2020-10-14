package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/darksford123x/app/ent"
	"github.com/darksford123x/app/ent/status"
	"github.com/gin-gonic/gin"
)

// StatusController defines the struct for the status controller
type StatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateStatus handles POST requests for adding status entities
// @Summary Create status
// @Description Create status
// @ID create-status
// @Accept   json
// @Produce  json
// @Param status body ent.Status true "Status entity"
// @Success 200 {object} ent.Status
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuses [post]
func (ctl *StatusController) CreateStatus(c *gin.Context) {
	obj := ent.Status{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "status binding failed",
		})
		return
	}

	u, err := ctl.client.Status.
		Create().
		SetAge(obj.Age).
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetStatus handles GET requests to retrieve a status entity
// @Summary Get a Status entity by ID
// @Description get Status by ID
// @ID get-Status
// @Produce  json
// @Param id path int true "Status ID"
// @Success 200 {object} ent.Status
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuses/{id} [get]
func (ctl *StatusController) GetStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Status.
		Query().
		Where(status.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListStatus handles request to get a list of status entities
// @Summary List status entities
// @Description list status entities
// @ID list-status
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Status
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuses [get]
func (ctl *StatusController) ListStatus(c *gin.Context) {
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

	statuses, err := ctl.client.Status.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, statuses)
}

// DeleteStatus handles DELETE requests to delete a status entity
// @Summary Delete a status entity by ID
// @Description get status by ID
// @ID delete-status
// @Produce  json
// @Param id path int true "Status ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuses/{id} [delete]
func (ctl *StatusController) DeleteStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Status.
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

// UpdateStatus handles PUT requests to update a status entity
// @Summary Update a status entity by ID
// @Description update status by ID
// @ID update-status
// @Accept   json
// @Produce  json
// @Param id path int true "Status ID"
// @Param status body ent.Status true "Status entity"
// @Success 200 {object} ent.Status
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuses/{id} [put]
func (ctl *StatusController) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Status{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "status binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Status.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewStatusController creates and registers handles for the status controller
func NewStatusController(router gin.IRouter, client *ent.Client) *StatusController {
	uc := &StatusController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitStatusController registers routes to the main engine
func (ctl *StatusController) register() {
	statuses := ctl.router.Group("/statuses")

	statuses.GET("", ctl.ListStatus)

	// CRUD
	statuses.POST("", ctl.CreateStatus)
	statuses.GET(":id", ctl.GetStatus)
	statuses.PUT(":id", ctl.UpdateStatus)
	statuses.DELETE(":id", ctl.DeleteStatus)
}
