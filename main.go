package main

import (
	"lab-go-gin-jwt/config"
	"lab-go-gin-jwt/handler"
	"lab-go-gin-jwt/middleware"
	"lab-go-gin-jwt/repository"
	"lab-go-gin-jwt/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authHandler    handler.AuthHandler       = handler.NewAuthHandler(authService, jwtService)
	userHandler    handler.UserHandler       = handler.NewUserHandler(userService, jwtService)
	bookHandler    handler.BookHandler       = handler.NewBookHandler(bookService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookHandler.FindAll)
		bookRoutes.POST("/", bookHandler.Insert)
		bookRoutes.GET("/:id", bookHandler.FindByID)
		bookRoutes.PUT("/:id", bookHandler.Update)
		bookRoutes.DELETE("/:id", bookHandler.Delete)
	}

	r.Run()
}
