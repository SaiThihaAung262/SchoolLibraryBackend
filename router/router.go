package router

import (
	"net/http"

	"MyGO.com/m/config"
	"MyGO.com/m/controller"
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

	//Client User
	clientRepository repository.ClientRepository = repository.NewClientRepository(db)
	clientService    service.ClientService       = service.NewClientService(clientRepository)
	clientController controller.ClientController = controller.NewClientController(clientService, jwtService)

	//BookCategory
	bookCategoryRepository repository.BookCategoryRepository = repository.NewBookCategoryRepository(db)
	bookCategoryService    service.BookCategoryService       = service.NewBookCategoryService(bookCategoryRepository)
	bookCategoryController controller.BookCategoryController = controller.NewBookCategoryControlle(bookCategoryService)

	//Book
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	bookController controller.BookController = controller.NewBookController(bookService)

	//Media
	mediaRepository repository.MediaRepository = repository.NewMediaRepository(db)
	mediaService    service.MedeiaService      = service.NewMediaService(mediaRepository)
	mediaController controller.MediaController = controller.NewMediaController(mediaService)
)

func InitRoute() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	r.Use(Cors())

	apiRoutes := r.Group("/api")

	//User end points
	userRoutes := apiRoutes.Group("admin")
	{
		userRoutes.POST("/register", authController.Register)
		userRoutes.POST("/login", authController.Login)
	}

	userAdminRoutes := apiRoutes.Group("admin-users")
	// userAdminRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		userAdminRoutes.POST("/create", userController.CreateUser)
		userAdminRoutes.GET("/get-users", userController.GetAllUsers)
		userAdminRoutes.POST("/update", userController.UpdateUser)
		userAdminRoutes.POST("/delete", userController.DeleteUser)

	}

	clientUserRoutes := apiRoutes.Group("client-users")
	{
		clientUserRoutes.POST("/create", clientController.CreateClient)
		clientUserRoutes.GET("/get-clients", clientController.GetAllClients)
		clientUserRoutes.POST("/update", clientController.UpdateClient)
		clientUserRoutes.POST("/delete", clientController.DeleteClient)

	}

	//Book Category end points
	bookCategoryRoutes := apiRoutes.Group("book-category")
	{
		bookCategoryRoutes.POST("/create", bookCategoryController.CreateBookCategory)
		bookCategoryRoutes.GET("/get-categories", bookCategoryController.GetAllBookCategory)
		bookCategoryRoutes.POST("/update", bookCategoryController.UpdateBookCategory)
		bookCategoryRoutes.POST("/delete", bookCategoryController.DeleteBookCategory)
	}

	//Book end points
	bookRoutes := apiRoutes.Group("book")
	{
		bookRoutes.POST("/create", bookController.CreateBook)
		bookRoutes.GET("/get-books", bookController.GetAllBooks)
		bookRoutes.POST("/update", bookController.UpdateBook)
		bookRoutes.POST("/delete", bookController.DeleteBook)
	}

	//Media end points
	mediaRoutes := apiRoutes.Group("media")
	{
		mediaRoutes.POST("/upload", mediaController.CreateMedia)
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
