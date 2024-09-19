package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/turk2003/workflow/controllers"
	"github.com/turk2003/workflow/middlewares"
	"github.com/turk2003/workflow/repositories"
	"github.com/turk2003/workflow/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// ที่ main.go
	db, err := gorm.Open(
		postgres.Open(
			os.Getenv("DATABASE_URL"),
		),
	)

	if err != nil {
		log.Panic(err)
	}

	repo := repositories.NewItemRepository(db)
	service := services.NewItemService(repo)
	itemController := controllers.NewItemController(service)
	authController := controllers.NewAuthController()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // ระบุ domain ที่อนุญาตให้เข้าถึง
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.GET("/items", itemController.GetAllItems)
	// r.GET("/items/:id", itemController.GetItemByID)
	r.POST("/login", authController.Login)
	// r.PUT("/items/:id", itemController.UpdateItem)
	authRoutes := r.Group("/")
	authRoutes.Use(middlewares.AuthMiddleware())
	{
		authRoutes.POST("/items", itemController.CreateItem)
		// authRoutes.GET("/items", itemController.GetAllItems)
		authRoutes.GET("/items/:id", itemController.GetItemByID)
		authRoutes.PUT("/items/:id", itemController.UpdateItem)
		authRoutes.PATCH("/items/:id", itemController.PatchItemStatus)
		authRoutes.DELETE("/items/:id", itemController.DeleteItem)
	}

	// // Routes for Items
	// r.POST("/items", itemController.CreateItem)
	// r.GET("/items", itemController.GetAllItems)
	// r.GET("/items/:id", itemController.GetItemByID)
	// r.PUT("/items/:id", itemController.UpdateItem)
	// r.PATCH("/items/:id", itemController.PatchItemStatus)
	// r.DELETE("/items/:id", itemController.DeleteItem)

	// // Route for Login
	// r.POST("/login", authController.Login)

	r.Run() // เริ่ม API server
}
