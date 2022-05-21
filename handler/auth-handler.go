package handler

import (
	"lab-go-gin-jwt/dto"
	"lab-go-gin-jwt/entity"
	"lab-go-gin-jwt/helper"
	"lab-go-gin-jwt/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthHandler(authService service.AuthService, jwtService service.JWTService) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (h *authHandler) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	err := ctx.ShouldBindJSON(&loginDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process login request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := h.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := h.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credentials", "Invalid credentials", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (h *authHandler) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	err := ctx.ShouldBindJSON(&registerDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !h.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)

	} else {

		createdUser := h.authService.CreateUser(registerDTO)
		token := h.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
