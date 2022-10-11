package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

    func EnvCloudName() string {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
        return os.Getenv("CLOUDINARY_CLOUD_NAME")
    }

    func EnvCloudAPIKey() string {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
        return os.Getenv("CLOUDINARY_API_KEY")
    }

    func EnvCloudAPISecret() string {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
        return os.Getenv("CLOUDINARY_API_SECRET")
    }

    func EnvCloudUploadFolder() string {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
        return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
    }

    func EnvJWTKey() string {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
        return os.Getenv("JWT_KEY")
    }

    func EnvMongoUrl() string {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
        return os.Getenv("MONGO_URL")
    }