package api

import (
	"net/http"
	"time"

	db "github.com/KothariMansi/hospitalOPD/db/sqlc"
	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/gin-gonic/gin"
)

type createCheckUpTimeRequest struct {
	Morning *time.Time `json:"morning"`
	Evening *time.Time `json:"evening"`
	Night   *time.Time `json:"night"`
}

func (server *Server) createCheckUpTime(ctx *gin.Context) {
	var req createCheckUpTimeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateCheckUpTimeParams{
		Morning: util.ToNullTime(req.Morning),
		Evening: util.ToNullTime(req.Evening),
		Night:   util.ToNullTime(req.Night),
	}
	result, err := server.store.CreateCheckUpTime(ctx, arg)
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

type getCheckUpTimeRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) getCheckUpTime(ctx *gin.Context) {
	var req getCheckUpTimeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	checkUpTime, err := server.store.GetCheckUpTime(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, checkUpTime)
}

type listCheckUpTimeRequest struct {
	PageId   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (server *Server) listCheckUpTime(ctx *gin.Context) {
	var req listCheckUpTimeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.PageId <= 0 {
		req.PageId = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	arg := db.ListCheckUpTimesParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	checkUpTimes, err := server.store.ListCheckUpTimes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, checkUpTimes)
}

type updateCheckUpTimeRequest struct {
	Id      int64      `json:"id" binding:"required"`
	Morning *time.Time `json:"morning"`
	Evening *time.Time `json:"evening"`
	Night   *time.Time `json:"night"`
}

func (server *Server) updateCheckUpTime(ctx *gin.Context) {
	var req updateCheckUpTimeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateCheckUpTimeParams{
		ID:      req.Id,
		Morning: util.ToNullTime(req.Morning),
		Evening: util.ToNullTime(req.Evening),
		Night:   util.ToNullTime(req.Night),
	}
	err := server.store.UpdateCheckUpTime(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}

type deleteCheckUpTimeRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteCheckUpTime(ctx *gin.Context) {
	var req deleteCheckUpTimeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteCheckUpTime(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}
