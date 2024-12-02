package repositories

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	ID          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MovieRepository interface {
	Create(movie *Movie) (*Movie, error)
	FindByID(id string) (*Movie, error)
	FindAll() ([]*Movie, error)
	Update(movie *Movie) (*Movie, error)
	Delete(id string) error
}

type GormMovieRepository struct {
	db *gorm.DB
}

func NewGormMovieRepository(db *gorm.DB) *GormMovieRepository {
	return &GormMovieRepository{db: db}
}

func (r *GormMovieRepository) Create(movie *Movie) (*Movie, error) {
	movie.ID = uuid.New().String()
	res := r.db.Create(movie)
	if res.Error != nil {
		return nil, res.Error
	}
	return movie, nil
}

func (r *GormMovieRepository) FindByID(id string) (*Movie, error) {
	var movie Movie
	res := r.db.First(&movie, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &movie, nil
}

func (r *GormMovieRepository) FindAll() ([]*Movie, error) {
	var movies []*Movie
	res := r.db.Find(&movies)
	if res.Error != nil {
		return nil, errors.New("no movies found")
	}
	return movies, nil
}

func (r *GormMovieRepository) Update(movie *Movie) (*Movie, error) {
	res := r.db.Model(movie).Where("id = ?", movie.ID).Updates(movie)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not updated")
	}
	return movie, nil
}

func (r *GormMovieRepository) Delete(id string) error {
	res := r.db.Where("id = ?", id).Delete(&Movie{})
	if res.RowsAffected == 0 {
		return errors.New("movie not deleted")
	}
	return nil
}
