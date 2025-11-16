package api

import (
	"database/sql"
	"net/http"

	db "github.com/KothariMansi/hospitalOPD/db/sqlc"
	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" biding:"required"`
	State    string `json:"state"`
	City     string `json:"city"`
	Gender   string `json:"gender"`
	Age      int64  `json:"age"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
		State:          req.State,
		City:           req.City,
		Gender:         db.UserGender(req.Gender),
		Age:            sql.NullInt32{Int32: int32(req.Age)},
	}
	result, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, id)

}
