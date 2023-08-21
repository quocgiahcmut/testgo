package api

import (
	"errors"
	"net/http"
	"strconv"

	"testmongo/dbmongo"

	"github.com/gin-gonic/gin"
)

type CreatePersonRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Bio  string `json:"bio"`
}

func (server *Server) CreatePerson(ctx *gin.Context) {
	var req CreatePersonRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := &dbmongo.CreatePersonParams{
		Name: req.Name,
		Age:  req.Age,
		Bio:  req.Bio,
	}

	person, err := server.mongoStore.InsertPerson(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, person)
}

func (server *Server) GetListPerson(ctx *gin.Context) {
	page, err := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Invalid page parameter")))
		return
	}

	perPage, err := strconv.ParseInt(ctx.DefaultQuery("per_page", "10"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Invalid per_page parameter")))
		return
	}

	people, err := server.mongoStore.GetListPerson(ctx, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, people)
}
