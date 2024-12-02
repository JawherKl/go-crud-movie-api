package router

import (
	"net/http"
	"example.com/go-crud-api/db"
	"example.com/go-crud-api/repositories" // Updated import
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMovie)
	r.POST("/movies", postMovie)
	r.PUT("/movies/:id", updateMovie)
	r.DELETE("/movies/:id", deleteMovie)
	return r
}

func postMovie(ctx *gin.Context) {
	var movie repositories.Movie // Updated reference
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := db.MovieRepo.Create(&movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func getMovies(ctx *gin.Context) {
	res, err := db.MovieRepo.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func getMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.MovieRepo.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func updateMovie(ctx *gin.Context) {
	var movie repositories.Movie // Updated reference
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("id")
	movie.ID = id
	res, err := db.MovieRepo.Update(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func deleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.MovieRepo.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "movie deleted successfully"})
}
