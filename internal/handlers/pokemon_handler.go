package handlers

import (
	"apiexterna/internal/database"
	"apiexterna/internal/domain/errors"
	"apiexterna/internal/repository"
	"apiexterna/internal/services"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PokemonHandler struct {
	service *services.PokemonService
}

func NewPokemonsHandler(db *database.PostgresDB) *PokemonHandler {
	pokemonRepo := repository.NewPokemonsRepository(db)
	return &PokemonHandler{
		service: services.NewPokemonsService(nil, pokemonRepo),
	}
}

// @Summary Integrate a new Pokemon
// @Description Integrate a new Pokemon by name or ID
// @Tags Pokemons
// @Accept json
// @Produce json
// @Param name path string true "Pokemon name or ID"
// @Success 200 {object} models.Pokemon
// @Failure 400 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /pokemons/integrate/{name} [post]
func (h *PokemonHandler) PostIntegratePokemon(c *gin.Context) {
	name := c.Param("name")

	c.Header("Content-Type", "application/json")

	ctx := context.Background()
	pokeIns, err := h.service.InsertPokemonInfo(ctx, name)

	if err != nil {
		if customErr, ok := err.(errors.CustomError); ok {
			switch customErr.Code() {
			case errors.ErrCodeNotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
					"code":  customErr.Code(),
				})
			case errors.ErrCodeInvalidInput:
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"code":  customErr.Code(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
					"code":  customErr.Code(),
				})
			}
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Pokemon successfully inserted",
		"data":    pokeIns,
	})
}

// @Summary Get all integrated Pokemons
// @Description Get all integrated Pokemons
// @Tags Pokemons
// @Accept json
// @Produce json
// @Success 200 {array} models.Pokemon
// @Failure 500 {object} errors.CustomError
// @Router /pokemons [get]
func (h *PokemonHandler) GetIntegratedPokemons(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	ctx := context.Background()

	pokemons, err := h.service.GetIntegratedPokemons(ctx)
	if err != nil {
		if customErr, ok := err.(errors.CustomError); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"code":  customErr.Code(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	message := "No Pokemons found"
	if len(pokemons) > 0 {
		message = "Pokemons retrieved successfully"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data":    pokemons,
	})
}

// @Summary Delete a Pokemon by ID
// @Description Delete a Pokemon by ID
// @Tags Pokemons
// @Accept json
// @Produce json
// @Param id path int true "Pokemon ID"
// @Success 200 {object} models.Pokemon
// @Failure 500 {object} errors.CustomError
// @Router /pokemons/deletePokemonByID/{id} [delete]
func (h *PokemonHandler) DeletePokemonByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	delPokemon, err := h.service.DeletePokemonByID(ctx, id)

	if err != nil {
		if customErr, ok := err.(errors.CustomError); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"code":  customErr.Code(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Pokemon deleted successfully",
		"data":    delPokemon,
	})
}
