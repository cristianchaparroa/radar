package api

import (
	"github.com/gin-gonic/gin"
	"radar/dataprovider/sql"
	"radar/pools"
	"radar/services"
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

}

func (c *LocationController) PubSubLocation() {

}
