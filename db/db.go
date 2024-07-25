package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type Movie struct {
	ID          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func InitPostgresDB() {
    // Load environment variables from .env file
    err = godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file", err)
    }

    // Get the DATABASE_URL environment variable
    databaseURL := os.Getenv("DATABASE_URL")

    var dsn string
    if databaseURL != "" {
        // Use DATABASE_URL if set
        dsn = databaseURL
    } else {
        // Construct the DSN from individual environment variables
        var (
            host     = os.Getenv("DB_HOST")
            port     = os.Getenv("DB_PORT")
            dbUser   = os.Getenv("DB_USER")
            dbName   = os.Getenv("DB_NAME")
            password = os.Getenv("DB_PASSWORD")
        )
        dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
            host,
            port,
            dbUser,
            dbName,
            password,
        )
    }

    // Connect to the database using GORM
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // AutoMigrate your models here
    db.AutoMigrate(&Movie{})
}

func CreateMovie(movie *Movie) (*Movie, error) {
	movie.ID = uuid.New().String()
	res := db.Create(&movie)
	if res.Error != nil {
		return nil, res.Error
	}
	return movie, nil
}

func GetMovie(id string) (*Movie, error) {
	var movie Movie
	res := db.First(&movie, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("movie of id %s not found", id))
	}
	return &movie, nil
}

func GetMovies() ([]*Movie, error) {
	var movies []*Movie
	res := db.Find(&movies)
	if res.Error != nil {
		return nil, errors.New("no movies found")
	}
	return movies, nil
}

func UpdateMovie(movie *Movie) (*Movie, error) {
	var movieToUpdate Movie
	result := db.Model(&movieToUpdate).Where("id = ?", movie.ID).Updates(movie)
	if result.RowsAffected == 0 {
		return &movieToUpdate, errors.New("movie not updated")
	}
	return movie, nil
}

func DeleteMovie(id string) error {
	var deletedMovie Movie
	result := db.Where("id = ?", id).Delete(&deletedMovie)
	if result.RowsAffected == 0 {
		return errors.New("movie not deleted")
	}
	return nil
}