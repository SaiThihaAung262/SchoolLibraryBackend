package router

import (
	"net/http"

	"MyGO.com/m/config"
	"MyGO.com/m/controller"
	"MyGO.com/m/middleware"
	"MyGO.com/m/repository"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDBConnection()
	//JWT
	jwtService service.JwtService = service.NewJwtService()

	//User
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewUserService(userRepository)
	authController controller.AuthController = controller.NewAuthContrller(userService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func InitRoute() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	r.Use(Cors())

	apiRoutes := r.Group("/api")

	//User routes
	userRoutes := apiRoutes.Group("auth")
	{
		userRoutes.POST("/register", authController.Register)
		userRoutes.POST("/login", authController.Login)
	}

	userAdminRoutes := apiRoutes.Group("users")
	userAdminRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		userAdminRoutes.GET("/get-all-users", userController.GetAllUsers)
		userAdminRoutes.POST("/update-user", userController.UpdateUser)
		userAdminRoutes.POST("/delete-use", userController.DeleteUser)

	}

	panic(r.Run(":8090"))
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,Origin,X-Requested-With,Content-Type,Accept")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		c.Next()
	}
}
