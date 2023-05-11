package routes

import (
	"Mini_Project_Toko-Online/constants"
	"Mini_Project_Toko-Online/controllers"
	m "Mini_Project_Toko-Online/middlewares"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	eUser := e.Group("/users")
	eUser.POST("/register", controllers.RegisterUserController)
	eUser.POST("/login", controllers.LoginUserController)
	// 	Authenticated with JWT

	eAdmin := e.Group("/admin")
	eAdmin.POST("/register", controllers.RegisterUserController)
	eAdmin.POST("/login", controllers.LoginUserController)

	eUserJwt := eUser.Group("")
	eUserJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUserJwt.GET("/:id", controllers.GetUserProfile)
	eUserJwt.PUT("/:id", controllers.UpdateUserController)
	eUserJwt.DELETE("/:id", controllers.DeleteUserController)

	eAdminJwt := eAdmin.Group("/admin")
	eAdminJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eAdminJwt.GET("", controllers.GetAdminControllerAll)
	eAdminJwt.GET("/:id", controllers.GetAdminProfile)
	eAdminJwt.PUT("/:id", controllers.UpdateAdminController)
	eAdminJwt.DELETE("/:id", controllers.DeleteAdminController)
	eUserJwt.GET("/user/:id", controllers.GetUsersId)
	eUserJwt.GET("/user", controllers.GetUsersControllerAll)

	eCategory := e.Group("/categories")
	eCategory.GET("", controllers.GetCategoryControllerAll)

	eCategoryJwt := eCategory.Group("/admin/category")
	eCategoryJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eCategoryJwt.GET("", controllers.GetCategoryControllerAll)
	eCategoryJwt.POST("", controllers.CreateCategoryController)
	eCategoryJwt.PUT("/:id", controllers.UpdateCategoryController)
	eCategoryJwt.DELETE("/:id", controllers.DeleteCategoryController)

	eProductJwt := e.Group("/product")
	eProductJwt.GET("", controllers.GetProductControllerAll)
	eProductJwt.GET("/:id", controllers.GetProductsController)

	//eProductJwt := eProduct.Group("/admin/product")
	//eProductJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	eProductJwt.GET("", controllers.GetProductControllerAll)
	eProductJwt.GET("/:id", controllers.GetProductsController)
	eProductJwt.POST("", controllers.CreateProductController)
	eProductJwt.PUT("/:id", controllers.UpdateProductController)
	eProductJwt.DELETE("/:id", controllers.DeleteProductController)

	eOrderJwt := e.Group("/user/order")
	eOrderJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eOrderJwt.GET("", controllers.GetOrdersController)
	eOrderJwt.GET("/:id", controllers.GetOrderController)
	eOrderJwt.DELETE("/:id", controllers.DeleteOrderController)

	//eOrderJwt := e.Group("/admin/order")
	//eOrderJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eOrderJwt.GET("", controllers.GetOrdersController)
	eOrderJwt.GET("/:id", controllers.GetOrderController)
	eOrderJwt.DELETE("/:id", controllers.DeleteOrderController)
	eOrderJwt.POST("/:id", controllers.CreateOrderController)

	return e
}
