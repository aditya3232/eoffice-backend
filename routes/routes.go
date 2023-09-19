package routes

import (
	"eoffice-backend/database"
	"eoffice-backend/handler"
	"eoffice-backend/middleware"
	"eoffice-backend/models/auth"
	"eoffice-backend/models/division"
	"eoffice-backend/models/employee"
	"eoffice-backend/models/location"
	"eoffice-backend/models/permission"
	"eoffice-backend/models/position"
	"eoffice-backend/models/profile"
	"eoffice-backend/models/role"
	"eoffice-backend/models/rolepermission"
	"eoffice-backend/models/user"

	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	userRepo := user.NewRepository(database.DB(), database.Redis())
	employeeRepo := employee.NewRepository(database.DB(), database.Redis())
	roleRepo := role.NewRepository(database.DB(), database.Redis())
	divisionRepo := division.NewRepository(database.DB(), database.Redis())
	positionRepo := position.NewRepository(database.DB(), database.Redis())
	locationRepo := location.NewRepository(database.DB(), database.Redis())
	permissionRepo := permission.NewRepository(database.DB(), database.Redis())
	rolePermissionRepo := rolepermission.NewRepository(database.DB(), database.Redis())

	// Initialize services
	authService := auth.NewService(userRepo, employeeRepo)
	userService := user.NewService(userRepo, employeeRepo)
	employeeService := employee.NewService(employeeRepo)
	roleService := role.NewService(roleRepo)
	divisionService := division.NewService(divisionRepo)
	positionService := position.NewService(positionRepo)
	locationService := location.NewService(locationRepo)
	permissionService := permission.NewService(permissionRepo)
	profileService := profile.NewService(userRepo, employeeRepo, roleRepo, permissionRepo, rolePermissionRepo)
	rolePermissionService := rolepermission.NewService(rolePermissionRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService, employeeService)
	userHandler := handler.NewUserHandler(userService)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	roleHandler := handler.NewRoleHandler(roleService)
	divisionHandler := handler.NewDivisionHandler(divisionService)
	positionHandler := handler.NewPositionHandler(positionService)
	locationHandler := handler.NewLocationHandler(locationService)
	permissionHandler := handler.NewPermissionHandler(permissionService)
	profileHandler := handler.NewProfileHandler(profileService)
	rolePermissionHandler := handler.NewRolePermissionHandler(rolePermissionService)

	// Configure routes
	api := router.Group("/api/v1")

	authRoutes := api.Group("/auth")
	userRoutes := api.Group("/user", middleware.AuthMiddleware(userService))
	employeeRoutes := api.Group("/employee", middleware.AuthMiddleware(userService))
	roleRoutes := api.Group("/role", middleware.AuthMiddleware(userService))
	divisionRoutes := api.Group("/division", middleware.AuthMiddleware(userService))
	positionRoutes := api.Group("/position", middleware.AuthMiddleware(userService))
	locationRoutes := api.Group("/location", middleware.AuthMiddleware(userService))
	permissionRoutes := api.Group("/permission", middleware.AuthMiddleware(userService))
	profileRoutes := api.Group("/profile", middleware.AuthMiddleware(userService))
	rolePermissionRoutes := api.Group("/role-permission", middleware.AuthMiddleware(userService))

	configureAuthRoutes(authRoutes, authHandler)
	configureUserRoutes(userRoutes, userHandler)
	configureEmployeeRoutes(employeeRoutes, employeeHandler)
	configureRoleRoutes(roleRoutes, roleHandler)
	configureDivisionRoutes(divisionRoutes, divisionHandler)
	configurePositionRoutes(positionRoutes, positionHandler)
	configureLocationRoutes(locationRoutes, locationHandler)
	configurePermissionRoutes(permissionRoutes, permissionHandler)
	configureProfileRoutes(profileRoutes, profileHandler)
	configureRolePermissionRoutes(rolePermissionRoutes, rolePermissionHandler)
}

func configureAuthRoutes(group *gin.RouterGroup, handler *handler.AuthHandler) {
	group.POST("/login", handler.Login)
	group.POST("/logout", handler.Logout)
}

func configureUserRoutes(group *gin.RouterGroup, handler *handler.UserHandler) {
	group.GET("/", handler.GetUser)
	group.GET("/:id", handler.GetUserByID)
	group.POST("/", handler.CreateUser)
	group.PUT("/:id", handler.UpdateUser)
	group.DELETE("/:id", handler.DeleteUser)
}

func configureEmployeeRoutes(group *gin.RouterGroup, handler *handler.EmployeeHandler) {
	group.GET("/", handler.GetEmployee)
	group.GET("/:id", handler.GetEmployeeByID)
	group.POST("/", handler.CreateEmployee)
	group.PUT("/:id", handler.UpdateEmployee)
	group.DELETE("/:id", handler.DeleteEmployee)
}

func configureRoleRoutes(group *gin.RouterGroup, handler *handler.RoleHandler) {
	group.GET("/", handler.GetRole)
	group.GET("/:id", handler.GetRoleByID)
	group.POST("/", handler.CreateRole)
	group.PUT("/:id", handler.UpdateRole)
	group.DELETE("/:id", handler.DeleteRole)
}

func configureDivisionRoutes(group *gin.RouterGroup, handler *handler.DivisionHandler) {
	group.GET("/", handler.GetDivision)
	group.GET("/:id", handler.GetDivisionByID)
	group.POST("/", handler.CreateDivision)
	group.PUT("/:id", handler.UpdateDivision)
	group.DELETE("/:id", handler.DeleteDivision)
}

func configurePositionRoutes(group *gin.RouterGroup, handler *handler.PositionHandler) {
	group.GET("/", handler.GetPosition)
	group.GET("/:id", handler.GetPositionByID)
	group.POST("/", handler.CreatePosition)
	group.PUT("/:id", handler.UpdatePosition)
	group.DELETE("/:id", handler.DeletePosition)
}

func configureLocationRoutes(group *gin.RouterGroup, handler *handler.LocationHandler) {
	group.GET("/", handler.GetLocation)
	group.GET("/:id", handler.GetLocationByID)
	group.POST("/", handler.CreateLocation)
	group.PUT("/:id", handler.UpdateLocation)
	group.DELETE("/:id", handler.DeleteLocation)
}

func configurePermissionRoutes(group *gin.RouterGroup, handler *handler.PermissionHandler) {
	group.GET("/", handler.GetPermission)
	group.GET("/:id", handler.GetPermissionByID)
	group.POST("/", handler.CreatePermission)
	group.PUT("/:id", handler.UpdatePermission)
	group.DELETE("/:id", handler.DeletePermission)
}

func configureRolePermissionRoutes(group *gin.RouterGroup, handler *handler.RolePermissionHandler) {
	group.GET("/", handler.GetRolePermission)
	group.GET("/:id", handler.GetRolePermissionByID)
	group.POST("/", handler.CreateRolePermission)
	group.PUT("/:id", handler.UpdateRolePermission)
	group.DELETE("/:id", handler.DeleteRolePermission)
}

func configureProfileRoutes(group *gin.RouterGroup, handler *handler.ProfileHandler) {
	group.GET("/", handler.GetProfile)
	group.GET("/permission", handler.GetPermission)
}
