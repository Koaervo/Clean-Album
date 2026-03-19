package main

import (
	adapter "CLEAN/Application/InterfaceAdapter"
	usecase "CLEAN/Application/Use-Case"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	app := fiber.New()
	godotenv.Load(".env.local")
	dsn := os.Getenv("DB_DSN")

	db, err := NewPostgresDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	albumRepo := adapter.NewGormAlbumRepository(db)
	albumService := usecase.NewAlbumService(albumRepo)
	albumHandler := adapter.NewHttpAlbumHandler(albumService)

	albumEndpoint := app.Group("album")
	{
		albumEndpoint.Post("", albumHandler.CreateAlbum)
		albumEndpoint.Put("", albumHandler.UpdateAlbum)
		albumEndpoint.Delete("/:id", albumHandler.DeleteAlbum)
		albumEndpoint.Get("", albumHandler.GetAllAlbum)
		albumEndpoint.Get("/:id", albumHandler.GetAlbumById)
	}

	log.Fatal(app.Listen(":8080"))
}

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
