package handler

import (
	"fmt"
	"lab-go-gin-jwt/dto"
	"lab-go-gin-jwt/helper"
	"lab-go-gin-jwt/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserHandler(userServ service.UserService, jwtServ service.JWTService) UserHandler {
	return &userHandler{
		userService: userServ,
		jwtService:  jwtServ,
	}
}

// Update implements UserHandler
func (h *userHandler) Update(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	err := ctx.ShouldBindJSON(&userUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = uint(id)
	u := h.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	ctx.JSON(http.StatusOK, res)

}

// Profile implements UserHandler
func (h *userHandler) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := h.userService.Profile(id)
	res := helper.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, res)
}
