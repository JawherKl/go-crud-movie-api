package omdb

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
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

func FetchMovies(query string) ([]Movie, error) {
    apiKey := os.Getenv("OMDB_API_KEY")
    url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", apiKey, query)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Search []Movie `json:"Search"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return result.Search, nil
}