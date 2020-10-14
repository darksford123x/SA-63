package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/darksford123x/app/ent"
	"github.com/darksford123x/app/ent/user"
	"github.com/gin-gonic/gin"
)

// SymptomController defines the struct for the symptom controller
type SymptomController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSymptom handles POST requests for adding symptom entities
// @Summary Create symptom
// @Description Create symptom
// @ID create-symptom
// @Accept   json
// @Produce  json
// @Param symptom body ent.Symptom true "Symptom entity"
// @Success 200 {object} ent.Symptom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptoms [post]
func (ctl *SymptomController) CreateSymptom(c *gin.Context) {
	obj := ent.Symptom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Symptom binding failed",
		})
		return
	}

	u, err := ctl.client.Symptom.
		Create().
		SetSymtomId(obj.Symptom).
		SetSymptomName(obj.Symptom).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetSymptom handles GET requests to retrieve a symptom entity
// @Summary Get a symptom entity by ID
// @Description get symptom by ID
// @ID get-symptom
// @Produce  json
// @Param id path int true "Symptom ID"
// @Success 200 {object} ent.Symptom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptoms/{id} [get]
func (ctl *SymptomController) GetSymptom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSymptom handles request to get a list of symptom entities
// @Summary List symptom entities
// @Description list symptom entities
// @ID list-symptom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Symptom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptoms [get]
func (ctl *SymptomController) ListSymptom(c *gin.Context) {
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

	users, err := ctl.client.User.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, symptoms)
}

// DeleteSymptom handles DELETE requests to delete a symptom entity
// @Summary Delete a symptom entity by ID
// @Description get symptom by ID
// @ID delete-symptom
// @Produce  json
// @Param id path int true "Symptom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptoms/{id} [delete]
func (ctl *SymptomController) DeleteSymptom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Symptom.
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

// UpdateSymptom handles PUT requests to update a symptom entity
// @Summary Update a symptom entity by ID
// @Description update symptom by ID
// @ID update-symptom
// @Accept   json
// @Produce  json
// @Param id path int true "Symptom ID"
// @Param symptom body ent.Symptom true "Symptom entity"
// @Success 200 {object} ent.Symptom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptoms/{id} [put]
func (ctl *SymptomController) UpdateSymptom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Symptom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "symptom binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Symptom.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewSymptomController creates and registers handles for the symptom controller
func NewSymptomController(router gin.IRouter, client *ent.Client) *UserController {
	uc := &SymptomController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitSymptomController registers routes to the main engine
func (ctl *SymptomController) register() {
	symptoms := ctl.router.Group("/symptoms")

	symptoms.GET("", ctl.ListSymptom)

	// CRUD
	symptoms.POST("", ctl.CreateSymptom)
	symptoms.GET(":id", ctl.GetSymptom)
	symptoms.PUT(":id", ctl.UpdateSymptom)
	symptoms.DELETE(":id", ctl.DeleteSymptom)
}
