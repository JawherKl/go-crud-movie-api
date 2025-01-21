package router

import (
	"net/http"
	"example.com/go-crud-api/db"
	"example.com/go-crud-api/repositories"
    "example.com/go-crud-api/omdb"
	"example.com/go-crud-api/auth"
    "github.com/gin-gonic/gin"
	"strconv"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Public Routes
	r.POST("/login", login)

	// Protected Routes with rate limiting
	r.GET("/movies", auth.AuthMiddleware(), getMovies)
	r.GET("/omdb_movies", func(c *gin.Context) {
        query := c.Query("query")
        movies, err := omdb.FetchMovies(query)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, movies)
    	})
	r.GET("/movies/:id", auth.AuthMiddleware(), getMovie)
	r.POST("/movies", auth.AuthMiddleware(), postMovie)
	r.PUT("/movies/:id", auth.AuthMiddleware(), updateMovie)
	r.DELETE("/movies/:id", auth.AuthMiddleware(), deleteMovie)

	return r
}

func login(ctx *gin.Context) {
	// Example login function
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Here, authenticate the user (e.g., check in the database)
	if credentials.Username == "admin" && credentials.Password == "password" {
		token, err := auth.GenerateToken("123") // Assuming 123 is the user ID
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func RegisterUserRoutes(router *gin.Engine, userRepository repositories.UserRepository) {
    router.GET("/user/:id", func(c *gin.Context) {
        id := c.Param("id")

        // Convert the id from string to uint
        userID, err := strconv.ParseUint(id, 10, 32) // Parse to uint
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }

        // Convert to uint
        user, err := userRepository.GetUserByID(uint(userID))
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"user": user})
    })

    // Add other routes like register, update, delete...
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
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "10")
	filter := ctx.DefaultQuery("filter", "")

	// Convert page and pageSize to integers
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	// Fetch movies from DB with pagination and filtering
	res, err := db.MovieRepo.FindWithPagination(pageInt, pageSizeInt, filter)
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
