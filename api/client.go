package api

import (
	"net/http"

	db "github.com/KothariMansi/hospitalOPD/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createClientRequest struct {
	Name  string `json:"name" binding:"required"`
	State string `json:"state" binding:"required"`
	City  string `json:"city" binding:"required"`
	Age   int32  `json:"age" binding:"required"`
}

func (server *Server) createClient(ctx *gin.Context) {
	var req createClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateClientParams{
		Name:  req.Name,
		State: req.State,
		City:  req.City,
		Age:   req.Age,
	}
	result, err := server.store.CreateClient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, id)
}
