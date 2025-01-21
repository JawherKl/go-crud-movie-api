package repositories

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	ID          string `json:"id" gorm:"primarykey"`
	Title       string `json:"title"`
    Year        string `json:"year"`
    Rated       string `json:"rated"`
    Released    string `json:"released"`
    Runtime     string `json:"runtime"`
    Genre       string `json:"genre"`
    Director    string `json:"director"`
    Writer      string `json:"writer"`
    Actors      string `json:"actors"`
    Plot        string `json:"plot"`
    Language    string `json:"language"`
    Country     string `json:"country"`
    Awards      string `json:"awards"`
    Poster      string `json:"poster"`
    Metascore   string `json:"metascore"`
    ImdbRating  string `json:"imdbrating"`
    ImdbVotes   string `json:"imdvotes"`
    ImdbID      string `json:"imdbid"`
    Type        string `json:"type"`
    DVD         string `json:"dvd"`
    BoxOffice   string `json:"boxoffice"`
    Production  string `json:"production"`
    Website     string `json:"website"`
}

type MovieRepository interface {
	Create(movie *Movie) (*Movie, error)
	FindByID(id string) (*Movie, error)
	FindAll() ([]*Movie, error)
	FindWithPagination(page, pageSize int, filter string) ([]*Movie, error)
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

func (r *GormMovieRepository) FindWithPagination(page, pageSize int, filter string) ([]*Movie, error) {
	var movies []*Movie
	offset := (page - 1) * pageSize
	query := r.db.Offset(offset).Limit(pageSize)

	// Filter by movie name or description
	if filter != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+filter+"%", "%"+filter+"%")
	}

	res := query.Find(&movies)
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
