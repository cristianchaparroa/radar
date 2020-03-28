package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"radar/adapters/presenter"
	"radar/presenters"
	"radar/providers/sql"
	"radar/services"
)

type StatusController struct {
	r *gin.RouterGroup
	service services.IStatus
	locationService services.ILocation
}

func NewStatusController(r *gin.RouterGroup, client sql.Client) *StatusController {

	service := services.NewStatus(client)
	locationService := services.NewLocation(client)

	controller := &StatusController{
		r:r,
		service: service,
		locationService:locationService,
	}

	return controller
}

func (c *StatusController) Setup() {
	c.setupRoutes()
}

func (c *StatusController) setupRoutes() {
	c.r.POST("statuses/", c.Create)
}

func (c *StatusController) Create(ctx *gin.Context) {

	var request presenters.CreateStatus
	err := ctx.BindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	location := presenter.MapLocationPresenterToEntity(request.Location)

	l, err := c.locationService.RegisterLocation(location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}


	s, err := c.service.Create(request.Status.ProfileID, l.ID, request.Status.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := presenter.MapCreateStatusEntityToPresenter(*s,l)
	ctx.JSON(http.StatusOK, response)
}