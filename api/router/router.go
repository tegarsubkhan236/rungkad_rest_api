package router

import (
	"example/api/controller"
	"example/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	authRoute := api.Group("/auth")
	authRoute.Post("/login", controller.Login)
	authRoute.Post("/register", controller.CreateUser)

	userRoute := api.Group("/user", middleware.Protected())
	userRoute.Get("/", controller.GetUsers)
	userRoute.Get("/list/:id", controller.GetUser)
	userRoute.Post("/", controller.CreateUser)
	userRoute.Put("/:id", controller.UpdateUser)
	userRoute.Delete("/:id", controller.DeleteUser)

	supplierRoute := api.Group("/supplier", middleware.Protected())
	supplierRoute.Post("/search/", controller.GetSupplierByColumn)
	supplierRoute.Get("/", controller.GetSuppliers)
	supplierRoute.Get("/:id", controller.GetSupplier)
	supplierRoute.Post("/", controller.CreateSupplier)
	supplierRoute.Put("/:id", controller.UpdateSupplier)
	supplierRoute.Delete("/:id", controller.DeleteSupplier)

	productCategoryRoute := api.Group("/product_category", middleware.Protected())
	productCategoryRoute.Get("/", controller.GetProductCategories)
	productCategoryRoute.Get("/:id", controller.GetProductCategory)
	productCategoryRoute.Post("/", controller.CreateProductCategory)
	productCategoryRoute.Put("/:id", controller.UpdateProductCategory)
	productCategoryRoute.Delete("/:id", controller.DeleteProductCategory)

	productRoute := api.Group("/product", middleware.Protected())
	productRoute.Get("/", controller.GetProducts)
	productRoute.Get("/search", controller.GetProductsByFilter)
	productRoute.Get("/:id", controller.GetProduct)
	productRoute.Post("/", controller.CreateProduct)
	productRoute.Put("/:id", controller.UpdateProduct)
	productRoute.Delete("/:id", controller.DeleteProduct)

	digiflazzRoute := api.Group("/digiflazz", middleware.Protected())
	digiflazzRoute.Get("/cek-saldo", controller.BalanceCheck)
	digiflazzRoute.Get("/price-list", controller.PriceList)
	digiflazzRoute.Get("/deposit", controller.Deposit)
}
