package main

import (
	"apiexterna/internal/database"
	"apiexterna/internal/handlers"
	"apiexterna/internal/middleware"
	"fmt"
	"log"
	"os"

	_ "apiexterna/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Pokemon API Integration
// @version         1.0
// @description     API for integration with PokeAPI and Pokemon management.
// @host            localhost:3000
// @BasePath        /api/v1

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	dbConfig := database.NewConfig()
	db, err := database.NewPostgresDB(dbConfig)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Gin framework
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.SetTrustedProxies([]string{"%s", os.Getenv("TRUSTED_PROXIES")})

	r.Use(middleware.LoggerMiddleware())

	// Initialize handlers
	pokemonsHandler := handlers.NewPokemonsHandler(db)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configure API routes
	api := r.Group("api/v1")
	{
		// Pokemon routes - Integration with PokeAPI
		pokemons := api.Group("/pokemons")
		{
			pokemons.GET("", pokemonsHandler.GetIntegratedPokemons)
			pokemons.POST("/integrate/:name", pokemonsHandler.PostIntegratePokemon)
			pokemons.DELETE("/deletePokemonByID/:id", pokemonsHandler.DeletePokemonByID)
		}
	}

	// Start HTTP server
	fmt.Println("Go server running on http://localhost:3000")
	if err := r.Run(":3000"); err != nil {
		fmt.Printf("Failed to initialize server: %v\n", err)
	}
}
