package api

import (
	"database/sql"
	"net/http"
	"strings"

	db "github.com/KothariMansi/hospitalOPD/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createClientRequest struct {
	Name   string `json:"name" binding:"required"`
	State  string `json:"state" binding:"required"`
	City   string `json:"city" binding:"required"`
	Number int64  `json:"number" binding:"required"`
	Age    int32  `json:"age" binding:"required"`
}

func (server *Server) createClient(ctx *gin.Context) {
	var req createClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateClientParams{
		Name:   req.Name,
		State:  req.State,
		City:   req.City,
		Number: req.Number,
		Age:    req.Age,
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

type getClientRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getClient(ctx *gin.Context) {
	var req getClientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := req.Id
	client, err := server.store.GetClient(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusAccepted, client)
}

type listClientRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=20"`
}

func (server *Server) listClients(ctx *gin.Context) {
	var req listClientRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListClientsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	clients, err := server.store.ListClients(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clients)
}

type deleteClientRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteClient(ctx *gin.Context) {
	var req deleteClientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := req.Id
	err := server.store.DeleteClient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}

type updateClientRequest struct {
	Id     int64  `json:"id" binding:"required"`
	Name   string `json:"name"`
	State  string `json:"state"`
	City   string `json:"city"`
	Number int64  `json:"number"`
	Age    int32  `json:"age"`
}

func (server *Server) updateClient(ctx *gin.Context) {
	var req updateClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateClientParams{
		ID:     req.Id,
		Name:   req.Name,
		State:  req.State,
		City:   req.City,
		Number: req.Number,
		Age:    req.Age,
	}
	err := server.store.UpdateClient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}

func (server *Server) countClients(ctx *gin.Context) {
	count, err := server.store.CountClients(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, count)
}

type SearchClientsByNameRequest struct {
	Name     string `form:"name" binding:"required"`
	PageId   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=1,max=20"`
}

func (server *Server) searchClientsByName(ctx *gin.Context) {
	var req SearchClientsByNameRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.SearchClientsByNameParams{
		Name:   "%" + strings.TrimSpace(req.Name) + "%",
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	clients, err := server.store.SearchClientsByName(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clients)
}
