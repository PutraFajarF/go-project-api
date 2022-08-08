package main

import (
	"github.com/PutraFajarF/go-project-api/config"
	"github.com/PutraFajarF/go-project-api/controller"
	"github.com/PutraFajarF/go-project-api/middleware"
	"github.com/PutraFajarF/go-project-api/repository"
	"github.com/PutraFajarF/go-project-api/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.GetAllBooks)
		bookRoutes.GET("/:id", bookController.GetBooksById)
		bookRoutes.POST("/", bookController.CreateBooks)
		bookRoutes.PUT("/:id", bookController.UpdateBooks)
		bookRoutes.DELETE("/:id", bookController.DeleteBooks)
	}

	r.Run()
}
