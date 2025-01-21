package omdb

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type Movie struct {
    Title       string `json:"Title"`
    Year        string `json:"Year"`
    Rated       string `json:"Rated"`
    Released    string `json:"Released"`
    Runtime     string `json:"Runtime"`
    Genre       string `json:"Genre"`
    Director    string `json:"Director"`
    Writer      string `json:"Writer"`
    Actors      string `json:"Actors"`
    Plot        string `json:"Plot"`
    Language    string `json:"Language"`
    Country     string `json:"Country"`
    Awards      string `json:"Awards"`
    Poster      string `json:"Poster"`
    Metascore   string `json:"Metascore"`
    ImdbRating  string `json:"imdbRating"`
    ImdbVotes   string `json:"imdbVotes"`
    ImdbID      string `json:"imdbID"`
    Type        string `json:"Type"`
    DVD         string `json:"DVD"`
    BoxOffice   string `json:"BoxOffice"`
    Production  string `json:"Production"`
    Website     string `json:"Website"`
    Response    string `json:"Response"`
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