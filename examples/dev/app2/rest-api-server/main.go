package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint      `gorm:"primary_key"`
	Task        string    `gorm:"size:255"`
	IsCompleted bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&Task{})

	r := gin.Default()
	// endpoint: tasks GET
	r.GET("/tasks", func(ctx *gin.Context) {
		var tasks []Task
		db.Find(&tasks)
		ctx.JSON(http.StatusOK, tasks)
	})

	// endpoint: tasks POST
	r.POST("/tasks", func(ctx *gin.Context) {
		var task Task
		if err := ctx.ShouldBindJSON(&task); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&task)
		ctx.JSON(http.StatusOK, task)
	})

	// endpoint: tasks PUT
	r.PUT("/tasks/:id", func(ctx *gin.Context) {
		var task Task
		id := ctx.Param("id")

		if err := db.First(&task, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		if err := ctx.ShouldBindJSON(&task); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&task)
		ctx.JSON(http.StatusOK, task)
	})

	// endpoint tasks DELETE
	r.DELETE("/tasks/:id", func(ctx *gin.Context) {
		var task Task
		id := ctx.Param("id")

		if err := db.First(&task, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		db.Delete(&task)
		ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	})

	// endpoint: home GET
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello gin.",
		})
	})

	r.Run(":8080")
}
