package api

import (
	"net/http"

	db "github.com/KothariMansi/hospitalOPD/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createSpecialityRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createSpeciality(ctx *gin.Context) {
	var req createSpecialityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result, err := server.store.CreateSpeciality(ctx, req.Name)
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

type getSpecialityRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) getSpeciality(ctx *gin.Context) {
	var req getSpecialityRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	speciality, err := server.store.GetSpeciality(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, speciality)
}

type listSpecialitiesRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1"`
}

func (server *Server) listSpecialities(ctx *gin.Context) {
	var req listSpecialitiesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListSpecialitiesParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	specialites, err := server.store.ListSpecialities(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, specialites)
}

type deleteSpecialityReqest struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteSpeciality(ctx *gin.Context) {
	var req deleteSpecialityReqest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteSpeciality(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}

type updateSpecialityRequest struct {
	Id   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (server *Server) updateSpeciality(ctx *gin.Context) {
	var req updateSpecialityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateSpecialityParams{
		ID:             req.Id,
		SpecialityName: req.Name,
	}
	err := server.store.UpdateSpeciality(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}

func (server *Server) countSpecialites(ctx *gin.Context) {
	count, err := server.store.CountSpecialities(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, count)
}

type searchSpecialitiesByNameRequest struct {
	Name     string `form:"name" binding:"required"`
	PageId   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=1"`
}

func (server *Server) searchSpecialitiesByName(ctx *gin.Context) {
	var req searchSpecialitiesByNameRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.SearchSpecialitiesByNameParams{
		SpecialityName: req.Name,
		Limit:          req.PageSize,
		Offset:         (req.PageId - 1) * req.PageSize,
	}
	specialities, err := server.store.SearchSpecialitiesByName(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, specialities)
}
