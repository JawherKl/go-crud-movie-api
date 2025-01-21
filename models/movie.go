package models

import "gorm.io/gorm"

type Movie struct {
    gorm.Model
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