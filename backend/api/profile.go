package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"radar/adapters/presenter"
	"radar/entities"
	"radar/presenters"
	"radar/providers/sql"
	"radar/services"
)

type ProfileController struct {
	r               *gin.RouterGroup
	service         services.IProfile
	statusService   services.IStatus
	locationService services.ILocation
}

func NewProfileController(r *gin.RouterGroup, client sql.Client) *ProfileController {

	service := services.NewProfile(client)
	statusService := services.NewStatus(client)
	locationService := services.NewLocation(client)

	controller := &ProfileController{r: r,
		service:         service,
		statusService:   statusService,
		locationService: locationService,
	}

	return controller
}

func (c *ProfileController) Setup() {
	c.setupRoutes()
}

func (c *ProfileController) setupRoutes() {
	c.r.GET("profiles/:id", c.GetProfile)
	c.r.POST("profiles/", c.CreateProfile)
}

func (c *ProfileController) GetProfile(ctx *gin.Context) {
	profileID := ctx.Param("id")

	p, err := c.service.GetProfile(profileID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, p)
}

func (c *ProfileController) CreateProfile(ctx *gin.Context) {

	var profileRequest presenters.CreateProfileRequest

	err := ctx.BindJSON(&profileRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	profile := presenter.MapProfilePresenterToEntity(profileRequest.Profile)
	p, err := c.service.Create(profile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	location := presenter.MapLocationPresenterToEntity(profileRequest.Location)
	l, err := c.locationService.RegisterLocation(location)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	s, err := c.statusService.Create(p.ID, l.ID, entities.NegativeStatus)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := presenter.NewProfileResponse(p, s)
	ctx.JSON(http.StatusOK, response)
}
