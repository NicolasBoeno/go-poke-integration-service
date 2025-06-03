package services

import (
	"apiexterna/internal/domain/errors"
	"apiexterna/internal/domain/models"
	"apiexterna/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type PokemonService struct {
	httpClient *http.Client
	repository *repository.PokemonRepository
}

type ResPokeStats struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stats    struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type ApiPokemonResponse struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Height int            `json:"height"`
	Weight int            `json:"weight"`
	Stats  []ResPokeStats `json:"stats"`
}

func NewPokemonsService(httpClient *http.Client, repository *repository.PokemonRepository) *PokemonService {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	return &PokemonService{
		httpClient: httpClient,
		repository: repository,
	}
}

// validateInputs checks if the required inputs are valid
func (s *PokemonService) validateInputs(name string) error {
	if os.Getenv("URL_POKEAPI") == "" {
		return errors.NewInternalError("URL_POKEAPI is not defined in .env", nil)
	}

	if name = strings.TrimSpace(name); name == "" {
		return errors.NewInvalidInputError("Pokemon name must not be empty")
	}
	return nil
}

// fetchPokemonFromAPI retrieves Pokemon data from the external API
func (s *PokemonService) fetchPokemonFromAPI(ctx context.Context, name string) (*ApiPokemonResponse, error) {
	url := fmt.Sprintf("%s/%s", os.Getenv("URL_POKEAPI"), strings.ToLower(name))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.NewAPIError("Failed to create request", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewAPIError("Failed to request PokeAPI", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIError(fmt.Sprintf("PokeAPI returned unexpected status code: %d", resp.StatusCode), nil)
	}

	var apiResponse ApiPokemonResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, errors.NewInternalError("Error decoding JSON response", err)
	}

	return &apiResponse, nil
}

// convertToDatabaseModel transforms API response to database model
func (s *PokemonService) convertToDatabaseModel(apiResponse *ApiPokemonResponse) *models.Pokemon {
	pokemonDB := &models.Pokemon{
		Id:     apiResponse.Id,
		Name:   apiResponse.Name,
		Height: apiResponse.Height,
		Weight: apiResponse.Weight,
		Stats:  make([]models.PokemonStats, len(apiResponse.Stats)),
	}

	for i, stat := range apiResponse.Stats {
		pokemonDB.Stats[i] = models.PokemonStats{
			BaseStat: stat.BaseStat,
			Effort:   stat.Effort,
			Name:     stat.Stats.Name,
		}
	}

	return pokemonDB
}

func (s *PokemonService) InsertPokemonInfo(ctx context.Context, name string) (*models.Pokemon, error) {
	if err := s.validateInputs(name); err != nil {
		return nil, errors.NewInvalidInputError(fmt.Sprintf("Input validation failed: %s", err.Error()))
	}

	// Fetch data from PokeAPI
	apiResponse, err := s.fetchPokemonFromAPI(ctx, name)
	if err != nil {
		return nil, err
	}

	// Convert to database model
	pokemonDB := s.convertToDatabaseModel(apiResponse)

	// Save in database
	if err := s.repository.Insert(pokemonDB); err != nil {
		return nil, errors.NewInternalError("Error on save pokemon on database", err)
	}

	return pokemonDB, nil
}

func (s *PokemonService) GetIntegratedPokemons(ctx context.Context) ([]models.Pokemon, error) {
	getPokemons, err := s.repository.GetAllPokemons()

	if err != nil {
		fmt.Println("Error getting pokemons from database:", err)
		return nil, errors.NewInternalError("Error getting pokemons from database", err)
	}

	return getPokemons, nil
}

func (s *PokemonService) DeletePokemonByID(ctx context.Context, id int) (*models.Pokemon, error) {
	delPokemon, err := s.repository.DeletePokemonByID(id)

	if err != nil {
		return nil, errors.NewInternalError("Error deleting pokemon", err)
	}

	return delPokemon, nil
}
