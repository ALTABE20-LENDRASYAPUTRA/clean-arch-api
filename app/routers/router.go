package routers

import (
	"clean-arch-api/app/middlewares"
	pd "clean-arch-api/features/product/data"
	ph "clean-arch-api/features/product/handler"
	ps "clean-arch-api/features/product/service"
	ud "clean-arch-api/features/user/data"
	uh "clean-arch-api/features/user/handler"
	us "clean-arch-api/features/user/service"
	"clean-arch-api/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {

	userData := ud.New(db)
	hash := encrypts.New()
	// userData := _userData.NewRaw(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService)

	productData := pd.New(db)
	productService := ps.New(productData)
	productHandlerAPI := ph.New(productService)
	// define routes/ endpoint
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/users", userHandlerAPI.GetAllUsers)
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	e.POST("/products", productHandlerAPI.CreateProduct, middlewares.JWTMiddleware())
	e.GET("/products", productHandlerAPI.GetAllProduct)
	e.PUT("/products/:product_id", productHandlerAPI.UpdateProduct, middlewares.JWTMiddleware())
	e.DELETE("/products/:product_id", productHandlerAPI.DeleteProduct, middlewares.JWTMiddleware())
	e.GET("/users/:user_id/products", productHandlerAPI.GetProductByIdUser, middlewares.JWTMiddleware())
}
