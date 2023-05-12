package routes

import (
	"Mini_Project_Toko-Online/controllers"
	m "Mini_Project_Toko-Online/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	eUser := e.Group("/user")
	eUser.POST("/register", controllers.RegisterUserController)
	eUser.POST("/login", controllers.LoginUserController)
	eUser.GET("/:id", controllers.GetUserProfile)
	eUser.PUT("/:id", controllers.UpdateUserController)
	eUser.DELETE("/:id", controllers.DeleteUserController)
	eUser.GET("/product", controllers.GetProductControllerAll)
	eUser.GET("/product/:id", controllers.GetProductController)
	eUser.GET("/category", controllers.GetCategoryControllerAll)
	eUser.GET("/category/:id", controllers.GetCategoryController)
	eUser.GET("/order", controllers.GetOrdersControllerAll)
	eUser.GET("/order/:id", controllers.GetOrdersControllerAll)
	eUser.POST("/order", controllers.CreateOrderController)
	eUser.DELETE("/order/:id", controllers.DeleteOrderController)
	//Authenticated with JWT

	eAdmin := e.Group("/admin")
	eAdmin.POST("/register", controllers.RegisterAdminController)
	eAdmin.POST("/login", controllers.LoginAdminController)
	eAdmin.GET("", controllers.GetAdminControllerAll)
	eAdmin.GET("/:id", controllers.GetAdminById)
	eAdmin.PUT("/:id", controllers.UpdateAdminController)
	eAdmin.DELETE("/:id", controllers.DeleteAdminController)
	eAdmin.GET("/user/:id", controllers.GetUsersId)
	eAdmin.GET("/user", controllers.GetUsersControllerAll)
	eAdmin.GET("/product", controllers.GetProductControllerAll)
	eAdmin.GET("/product/:id", controllers.GetProductController)
	eAdmin.POST("/product", controllers.CreateProductController)
	eAdmin.PUT("/product/:id", controllers.UpdateProductController)
	eAdmin.DELETE("/product/:id", controllers.DeleteProductController)
	eAdmin.GET("/category", controllers.GetCategoryControllerAll)
	eAdmin.GET("/category/:id", controllers.GetCategoryController)
	eAdmin.POST("/category", controllers.CreateCategoryController)
	eAdmin.PUT("/category/:id", controllers.UpdateCategoryController)
	eAdmin.DELETE("/category/:id", controllers.DeleteCategoryController)
	eAdmin.GET("/order", controllers.GetOrdersControllerAll)
	eAdmin.GET("/order/:id", controllers.GetOrdersControllerAll)

	return e
}
