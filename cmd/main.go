package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/timofeyka25/test-jungle/internal/handler"
	"github.com/timofeyka25/test-jungle/internal/repository"
	"github.com/timofeyka25/test-jungle/internal/usecase"
	"github.com/timofeyka25/test-jungle/pkg/cloudstorage"
	database "github.com/timofeyka25/test-jungle/pkg/db"
	"github.com/timofeyka25/test-jungle/pkg/drivestorage"
	"github.com/timofeyka25/test-jungle/pkg/jwt"
	"github.com/timofeyka25/test-jungle/pkg/server"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init db
	db, err := database.NewDBConnect(database.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer closeDB(db)

	// init google cloud storage file uploader
	cloudFileUploader, err := cloudstorage.NewFileUploader(cloudstorage.Config{
		CloudCredentialsPath: os.Getenv("CLOUD_CREDENTIALS_PATH"),
		ImagesBucketName:     os.Getenv("IMAGES_BUCKET_NAME"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// init google drive file uploader
	driveFileUploader, err := drivestorage.NewFileUploader(drivestorage.Config{
		CloudCredentialsPath: os.Getenv("CLOUD_CREDENTIALS_PATH"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// init helpers
	jwtGenerator := jwt.NewTokenGenerator(jwt.Config{SecretKey: os.Getenv("JWT_SECRET_KEY")})
	jwtValidator := jwt.NewTokenValidator(jwt.Config{SecretKey: os.Getenv("JWT_SECRET_KEY")})

	// init repositories
	userRepository := repository.NewUserRepository(db)
	imageRepository := repository.NewImageRepository(db)

	// init use cases
	userUseCase := usecase.NewUserUseCase(userRepository, jwtGenerator)
	imageUseCase := usecase.NewImageUseCase(imageRepository)

	// init handlers
	userHandler := handler.NewUserHandler(userUseCase)
	imageHandler := handler.NewImageHandler(
		imageUseCase,
		cloudFileUploader,
		driveFileUploader,
	)
	handlers := handler.NewHandler(
		userHandler,
		imageHandler,
	)

	// init app
	app := server.NewHTTPServer(server.Config{ReadTimeoutSeconds: os.Getenv("SERVER_READ_TIMEOUT")}, jwtValidator)
	handlers.InitRoutes(app)

	// run app
	go func() {
		err = app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
		if err != nil {
			log.Println("Server unexpectedly stopped")
		}
	}()

	// shutdown app and db
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)
	<-stop

	if err = app.Shutdown(); err != nil {
		log.Printf("Error shutting down server %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}

func closeDB(db io.Closer) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing db connection %s", err)
	} else {
		log.Println("DB connection gracefully closed")
	}
}
