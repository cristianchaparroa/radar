package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"radar/initializer"
	"radar/providers/sql"
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

	pc *ProfileController
	lc *LocationController
	sc *StatusController
}

func NewRadarServer() IRadarServer {
	gin.ForceConsoleColor()
	logger, _ := zap.NewProduction()
	engine := gin.Default()
	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, true))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	engine.Use(cors.New(config))

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
	ws := s.engine.Group("/ws/")

	s.pc = NewProfileController(server, s.sql)
	s.pc.Setup()

	s.lc = NewLocationController(ws, s.sql)
	s.lc.Setup()

	s.sc = NewStatusController(server, s.sql)
	s.sc.Setup()
}

func (s *RadarServer) Run() {
	s.lc.StartPool()
	s.engine.Run(":8080")
}

func (s *RadarServer) Close() {
	db := s.sql.GetConnection()
	db.Close()
}
