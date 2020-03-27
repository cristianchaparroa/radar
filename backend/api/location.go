package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"radar/pools"
	"radar/providers/sql"
	"radar/providers/websocket"
	"radar/services"

	domainAdapters "radar/adapters/entities"
)

type LocationController struct {
	r   *gin.RouterGroup
	sql sql.Client

	pool    pools.ILocationPool
	service services.ILocation
}

func NewLocationController(r *gin.RouterGroup, client sql.Client) *LocationController {

	service := services.NewLocation(client)
	pool := pools.NewLocationPool()

	return &LocationController{
		r:       r,
		sql:     client,
		service: service,
		pool:    pool}
}

func (c *LocationController) Setup() {
	c.setupRoutes()
}

func (c *LocationController) setupRoutes() {
	c.r.GET("locations", c.GetLocations)
}

func (c *LocationController) GetLocations(ctx *gin.Context) {

	connection, err := websocket.Upgrade(ctx.Writer, ctx.Request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	location, err := domainAdapters.GetLocation(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	gconnection := websocket.NewConnection(connection)
	locationClient := pools.NewLocationClient(c.pool, gconnection, location)

	registerChannel := c.pool.GetRegisterChannel()
	registerChannel <- locationClient
	go locationClient.Read()
}

func (c *LocationController) StartPool() {
	go c.pool.Start()
}
