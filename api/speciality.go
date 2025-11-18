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
