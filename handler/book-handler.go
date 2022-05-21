package handler

import (
	"fmt"
	"lab-go-gin-jwt/dto"
	"lab-go-gin-jwt/entity"
	"lab-go-gin-jwt/helper"
	"lab-go-gin-jwt/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type BookHandler interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookHandler struct {
	bookService service.BookService
	jwtService  service.JWTService
}

func NewBookHandler(bookServ service.BookService, jwtServ service.JWTService) BookHandler {
	return &bookHandler{
		bookService: bookServ,
		jwtService:  jwtServ,
	}
}

func (h *bookHandler) FindAll(ctx *gin.Context) {
	var books []entity.Book = h.bookService.FindAll()
	res := helper.BuildResponse(true, "OK", books)
	ctx.JSON(http.StatusOK, res)
}

func (h *bookHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusOK, res)
		return
	}

	var book entity.Book = h.bookService.FindByID(uint(id))
	if (book == entity.Book{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", book)
		ctx.JSON(http.StatusOK, res)
	}
}

func (h *bookHandler) Insert(ctx *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	err := ctx.ShouldBindJSON(&bookCreateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := h.getUserIDByToken(authHeader)
		id, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			bookCreateDTO.UserID = uint(id)
		}
		result := h.bookService.Insert(bookCreateDTO)
		res := helper.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusCreated, res)
	}
}

func (h *bookHandler) Update(ctx *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	err := ctx.ShouldBindJSON(&bookUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if h.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		id, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			bookUpdateDTO.UserID = uint(id)
		}
		result := h.bookService.Update(bookUpdateDTO)
		res := helper.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("You don't have permission", "Your not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
	}

}

func (h *bookHandler) Delete(ctx *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	book.ID = uint(id)
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if h.bookService.IsAllowedToEdit(userID, book.ID) {
		h.bookService.Delete(book)
		res := helper.BuildResponse(true, "Delete", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("You don't have permission", "Your not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
	}
}

func (h *bookHandler) getUserIDByToken(token string) string {
	aToken, err := h.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
