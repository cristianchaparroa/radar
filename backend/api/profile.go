package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"radar/providers/sql"
	"radar/entities"
	"radar/services"
)

type ProfileController struct {
	r       *gin.RouterGroup
	service services.IProfile
}

func NewProfileController(r *gin.RouterGroup, client sql.Client) *ProfileController {

	service := services.NewProfile(client)
	controller := &ProfileController{service: service, r: r}
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

	var profile entities.Profile

	err := ctx.BindJSON(&profile)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	p, err := c.service.Create(profile)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, p)
}
