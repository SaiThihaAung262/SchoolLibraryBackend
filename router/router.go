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

	//Student
	studentRepository repository.SutudentRepository = repository.NewStudentRepository(db)
	studentService    service.StudentService        = service.NewStudentService(studentRepository)
	studentController controller.StudentController  = controller.NewStudentController(studentService, jwtService)

	//Teacher
	teacherRepository repository.TeacherRepository = repository.NeweTeacherRepository(db)
	teacherService    service.TeacherService       = service.NewTeacherService(teacherRepository)
	teacherController controller.TeacherController = controller.NewTeacherController(teacherService, jwtService)

	//Staff
	staffRepository repository.StaffRepository = repository.NeweStaffRepository(db)
	staffService    service.StaffService       = service.NewStaffService(staffRepository)
	staffController controller.StaffController = controller.NewStaffController(staffService, jwtService)

	//BookCategory
	bookCategoryRepository repository.BookCategoryRepository = repository.NewBookCategoryRepository(db)
	bookCategoryService    service.BookCategoryService       = service.NewBookCategoryService(bookCategoryRepository)
	bookCategoryController controller.BookCategoryController = controller.NewBookCategoryControlle(bookCategoryService)

	//BookSubCategory
	bookSubCategoryRepository repository.BookSubCategoryRepository = repository.NewBookSubCategoryRepository(db)
	bookSubCategoryService    service.BookSubCategoryService       = service.NewBookSubCategoryService(bookSubCategoryRepository)
	bookSubCategoryController controller.BookSubCategoryController = controller.NewBookSubCategoryControlle(bookSubCategoryService)

	//Book
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	bookController controller.BookController = controller.NewBookController(bookService)

	//Media
	mediaRepository repository.MediaRepository = repository.NewMediaRepository(db)
	mediaService    service.MedeiaService      = service.NewMediaService(mediaRepository)
	mediaController controller.MediaController = controller.NewMediaController(mediaService)

	//Borrow Log
	borrowLogRepo    repository.BorrowLogRepository = repository.NewBorrowlogRepository(db)
	borrowLogService service.BorrowLogService       = service.NewBorrowLogService(borrowLogRepo)

	//Punishment
	punishmentRepo       repository.PunishmentRepository = repository.NewPunishmentRepository(db)
	punishmentService    service.PunishmentService       = service.NewPunishmentService(punishmentRepo)
	punishmentController controller.PunishmentController = controller.NewPunishmentController(punishmentService)

	//SystemConfig
	systemConfigRepo       repository.SystemConfigRepository = repository.NewSystemConfigRepo(db)
	systemConfigService    service.SystemConfigService       = service.NewSystemConfigService(systemConfigRepo)
	systemConfigController controller.SystemConfigController = controller.NewSystemConfigController(systemConfigService)

	//Borrow
	borrowRepo       repository.BorrowRepository = repository.NewBorrowRepository(db)
	borrowService    service.Borrowservice       = service.NewBorrowService(borrowRepo)
	borrowController controller.BorrowController = controller.NewBorrowController(borrowService, bookService, teacherService, studentService, staffService, borrowLogService, systemConfigService, punishmentService)

	//For clients
	//Clent login
	clientAuthController controller.ClientAuthController = controller.NewClientAuthController(studentService, teacherService, staffService, jwtService)
)

func InitRoute() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	r.Use(Cors())

	apiRoutes := r.Group("/api")

	//User end points
	userRoutes := apiRoutes.Group("admin")
	{
		// userRoutes.POST("/register", authController.Register)
		userRoutes.POST("/login", authController.Login)
	}

	userAdminRoutes := apiRoutes.Group("admin-users")
	// userAdminRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		userAdminRoutes.POST("/create", userController.CreateUser)
		userAdminRoutes.GET("/get-users", userController.GetAllUsers)
		userAdminRoutes.POST("/update", userController.UpdateUser)
		userAdminRoutes.POST("/delete", userController.DeleteUser)
		userAdminRoutes.GET("/dashboard", userController.GetDashbordData)
		userAdminRoutes.GET("/popular-books", userController.GetMostBorrowLog)

	}

	//Students end points
	studentRoutes := apiRoutes.Group("student")
	// studentRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		studentRoutes.POST("/create", studentController.CreateStudent)
		studentRoutes.GET("/get-students", studentController.GetAllStudents)
		studentRoutes.POST("/update", studentController.UpdateStudent)
		studentRoutes.POST("/delete", studentController.DeleteStudent)

	}

	//Teachers end points
	teacherRoutes := apiRoutes.Group("teacher")
	// teacherRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		teacherRoutes.POST("/create", teacherController.CreateTeacher)
		teacherRoutes.GET("/get-teachers", teacherController.GetAllTeachers)
		teacherRoutes.POST("/update", teacherController.UpdateTeacher)
		teacherRoutes.POST("/delete", teacherController.DeleteTeacher)

	}

	//Teachers end points
	staffRoutes := apiRoutes.Group("staff")
	// staffRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		staffRoutes.POST("/create", staffController.CreateStaff)
		staffRoutes.GET("/get-staffs", staffController.GetAllStaff)
		staffRoutes.POST("/update", staffController.UpdateStaff)
		staffRoutes.POST("/delete", staffController.DeleteStaff)

	}

	//Book Category end points
	bookCategoryRoutes := apiRoutes.Group("book-category")
	// bookCategoryRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		bookCategoryRoutes.POST("/create", bookCategoryController.CreateBookCategory)
		bookCategoryRoutes.GET("/get-categories", bookCategoryController.GetAllBookCategory)
		bookCategoryRoutes.POST("/update", bookCategoryController.UpdateBookCategory)
		bookCategoryRoutes.POST("/delete", bookCategoryController.DeleteBookCategory)
	}

	//Book sub Category end points
	bookSubCategoryRoutes := apiRoutes.Group("book-sub-category")
	// bookCategoryRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		bookSubCategoryRoutes.POST("/create", bookSubCategoryController.CreateBookSubCategory)
		bookSubCategoryRoutes.GET("/get-categories", bookSubCategoryController.GetAllBookSubCategory)
		bookSubCategoryRoutes.POST("/update", bookSubCategoryController.UpdateBookSubCategory)
		bookSubCategoryRoutes.POST("/delete", bookSubCategoryController.DeleteBookSubCategory)
	}

	//Book end points
	bookRoutes := apiRoutes.Group("book")
	// bookRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.POST("/create", bookController.CreateBook)
		bookRoutes.GET("/get-books", bookController.GetAllBooks)
		bookRoutes.POST("/update", bookController.UpdateBook)
		bookRoutes.POST("/delete", bookController.DeleteBook)
	}

	//Borrow end points
	borrowRoutes := apiRoutes.Group("borrow")
	// borrowRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		borrowRoutes.POST("/create", borrowController.CreateBorrow)
		borrowRoutes.GET("/get-history", borrowController.GetBorrowHistory)
		borrowRoutes.POST("/update-status", borrowController.UpdateBorrowStatus)
		borrowRoutes.GET("/get-summary", borrowController.GetBookSummaryData)
		borrowRoutes.POST("/re-borrow", borrowController.ReBorrow)

	}

	//Punishmet end points
	punishRoutes := apiRoutes.Group("punishment")
	// borrowRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		punishRoutes.POST("/create", punishmentController.CreatePunishment)
		punishRoutes.GET("/get", punishmentController.GetPunishmentData)
		punishRoutes.POST("/update", punishmentController.UpdatePunishment)
		punishRoutes.POST("/delete", punishmentController.DeletePunishment)

	}

	//SystemConfig end points
	systemConfigRoutes := apiRoutes.Group("system-config")
	// borrowRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		systemConfigRoutes.POST("/create", systemConfigController.CreateSystemConfg)
		systemConfigRoutes.GET("/get", systemConfigController.GetSystemConfig)
		systemConfigRoutes.POST("/update", systemConfigController.UpdateSystemConfig)
		systemConfigRoutes.POST("/delete", systemConfigController.DeleteSystemConfig)

	}

	//Client Login end points
	clientAuthRoutes := apiRoutes.Group("client")
	{
		clientAuthRoutes.POST("/login", clientAuthController.ClientLogin)

	}

	//Client user end points
	clientUserRoutes := apiRoutes.Group("user")
	// clientUserRoutes.Use(middleware.AuthorizeJWT(jwtService))
	{
		clientUserRoutes.GET("/get-books", bookController.GetAllBooks)
		clientUserRoutes.GET("/get-book-detail", borrowController.GetBookByUUID)
		clientUserRoutes.GET("/get-categories", bookCategoryController.GetAllBookCategory)
		clientUserRoutes.GET("/get-sub-categories", bookSubCategoryController.GetAllBookSubCategory)
		clientUserRoutes.GET("/get-user", clientAuthController.GetClientByUUID)
		clientUserRoutes.GET("/get-borrow-history", borrowController.GetBorrowHistory)
		clientUserRoutes.POST("/borrow-book", borrowController.CreateBorrow)
		clientUserRoutes.POST("/change-password", clientAuthController.ChangePassword)

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
