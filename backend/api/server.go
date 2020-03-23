package api

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"radar/dataprovider/sql"
	"radar/initializer"
	"time"
)

type IRadarServer interface {
	Setup()
	Run()
	Close()
}

type RadarServer struct {
	sql    sql.Client
	engine *gin.Engine
}

func NewRadarServer() IRadarServer {
	gin.ForceConsoleColor()
	logger, _ := zap.NewProduction()

	engine := gin.Default()
	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, true))

	return &RadarServer{engine: engine}
}

func (s *RadarServer) Setup() {
	s.setupDB()
	s.migrate()
	s.setupRoutes()
}

func (s *RadarServer) setupDB() {

	client, err := sql.NewPostgresClient()

	if err != nil {
		panic(err)
	}

	s.sql = client

}

func (s *RadarServer) migrate() {
	init := initializer.NewInitializer(s.sql)
	init.CreateTables()
}

func (s *RadarServer) setupRoutes() {
	server := s.engine.Group("/v1/")

	pc := NewProfileController(server, s.sql)
	pc.Setup()
}

func (s *RadarServer) Run() {
	s.engine.Run(":8080")
}

func (s *RadarServer) Close() {
	db := s.sql.GetConnection()
	db.Close()
}
