package api

import (
	"net/http"

	"testmongo/dbmongo"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router     *gin.Engine
	mongoStore dbmongo.MongoStore
}

func NewServer(mongoStore dbmongo.MongoStore) *Server {
	server := &Server{
		mongoStore: mongoStore,
	}

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	router.POST("/person", server.CreatePerson)
	router.GET("/person", server.GetListPerson)

	server.Router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address) // pass err to output
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
