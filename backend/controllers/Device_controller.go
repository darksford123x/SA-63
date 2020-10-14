package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/darksford123x/app/ent"
	"github.com/gin-gonic/gin"
)

// DeviceController defines the struct for the device controller
type DeviceController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDevice handles POST requests for adding device entities
// @Summary Create device
// @Description Create device
// @ID create-device
// @Accept   json
// @Produce  json
// @Param device body ent.Device true "Device entity"
// @Success 200 {object} ent.Device
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /devices [post]
func (ctl *DeviceController) CreateDevice(c *gin.Context) {
	obj := ent.Device{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "device binding failed",
		})
		return
	}

	u, err := ctl.client.Device.
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

// GetDevice handles GET requests to retrieve a device entity
// @Summary Get a device entity by ID
// @Description get device by ID
// @ID get-device
// @Produce  json
// @Param id path int true "Device ID"
// @Success 200 {object} ent.Device
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /devices/{id} [get]
func (ctl *DeviceController) GetDevice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Device.
		Query().
		Where(Device.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListDevice handles request to get a list of device entities
// @Summary List device entities
// @Description list device entities
// @ID list-device
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Device
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /devices [get]
func (ctl *DeviceController) ListDevice(c *gin.Context) {
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

	devices, err := ctl.client.Device.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, devices)
}

// DeleteDevice handles DELETE requests to delete a device entity
// @Summary Delete a device entity by ID
// @Description get device by ID
// @ID delete-device
// @Produce  json
// @Param id path int true "Device ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /devices/{id} [delete]
func (ctl *DeviceController) DeleteDevice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Device.
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

// UpdateDevice handles PUT requests to update a device entity
// @Summary Update a device entity by ID
// @Description update device by ID
// @ID update-device
// @Accept   json
// @Produce  json
// @Param id path int true "Device ID"
// @Param device body ent.Device true "Device entity"
// @Success 200 {object} ent.Device
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /devices/{id} [put]
func (ctl *DeviceController) UpdateDevice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Device{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "device binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Device.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewDeviceController creates and registers handles for the device controller
func NewDeviceController(router gin.IRouter, client *ent.Client) *DeviceController {
	uc := &DeviceController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDeviceController registers routes to the main engine
func (ctl *DeviceController) register() {
	devices := ctl.router.Group("/devices")

	devices.GET("", ctl.ListDevice)

	// CRUD
	devices.POST("", ctl.CreateDevice)
	devices.GET(":id", ctl.GetDevice)
	devices.PUT(":id", ctl.UpdateDevice)
	devices.DELETE(":id", ctl.DeleteDevice)
}
