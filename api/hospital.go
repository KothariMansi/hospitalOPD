package api

import (
	"database/sql"
	"net/http"

	db "github.com/KothariMansi/hospitalOPD/db/sqlc"
	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/gin-gonic/gin"
)

type createHospitalRequest struct {
	Name    string `json:"name" binding:"required"`
	Photo   string `json:"photo"`
	State   string `json:"state" binding:"required"`
	City    string `json:"city" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone"`
}

func (server *Server) createHospital(ctx *gin.Context) {
	var req createHospitalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateHospitalParams{
		Name:    req.Name,
		Photo:   util.ToNullString(req.Photo),
		State:   req.State,
		City:    req.City,
		Address: req.Address,
		Phone:   util.ToNullString(req.Phone),
	}
	result, err := server.store.CreateHospital(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, id)
}

type getHospitalRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getHospital(ctx *gin.Context) {
	var req getHospitalRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hospital, err := server.store.GetHospital(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, hospital)
}

type listHospitalRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1"`
}

func (server *Server) listHospital(ctx *gin.Context) {
	var req listHospitalRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListHospitalsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	hospitals, err := server.store.ListHospitals(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, hospitals)
}

type updateHospitalRequest struct {
	Id      int64          `json:"id" binding:"required"`
	Name    string         `json:"name"`
	Photo   sql.NullString `json:"photo"`
	State   string         `json:"state"`
	City    string         `json:"city"`
	Address string         `json:"address"`
	Phone   sql.NullString `json:"phone"`
}

func (server *Server) updateHospital(ctx *gin.Context) {
	var req updateHospitalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateHospitalParams{
		ID:      req.Id,
		Name:    req.Name,
		Photo:   req.Photo,
		State:   req.State,
		City:    req.State,
		Address: req.Address,
		Phone:   req.Phone,
	}
	err := server.store.UpdateHospital(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}

type deleteHospitalRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteHospital(ctx *gin.Context) {
	var req deleteHospitalRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteHospital(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}
