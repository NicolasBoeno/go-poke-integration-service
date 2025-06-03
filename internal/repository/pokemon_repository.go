package repository

import (
	"apiexterna/internal/database"
	"apiexterna/internal/domain/models"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
)

type PokemonRepository struct {
	db *database.PostgresDB
}

func NewPokemonsRepository(db *database.PostgresDB) *PokemonRepository {
	return &PokemonRepository{db: db}
}

func (r *PokemonRepository) Insert(pokemon *models.Pokemon) error {
	query := `
	Insert Into pokemons (id, name, height, weight, stats)
	values ($1, $2, $3, $4, $5)
	returning id`

	statsJson, err := json.Marshal(pokemon.Stats)
	if err != nil {
		return fmt.Errorf("error converting stats to json: %v", err)
	}

	err = r.db.GetDb().QueryRow(
		query,
		pokemon.Id,
		pokemon.Name,
		pokemon.Height,
		pokemon.Weight,
		statsJson,
	).Scan(&pokemon.Id)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return fmt.Errorf("pokemon with id %d already exists", pokemon.Id)
			}
		}
		return fmt.Errorf("error inserting pokemon: %v", err)
	}

	return nil
}

func (r *PokemonRepository) GetAllPokemons() ([]models.Pokemon, error) {
	query := `SELECT id, name, height, weight, stats FROM pokemons`
	rows, err := r.db.GetDb().Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying pokemons: %v", err)
	}
	defer rows.Close()

	var pokemons []models.Pokemon
	for rows.Next() {
		var pokemon models.Pokemon
		var statsJson []byte

		if err := rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Height, &pokemon.Weight, &statsJson); err != nil {
			return nil, fmt.Errorf("error scanning pokemon: %v", err)
		}

		if err := json.Unmarshal(statsJson, &pokemon.Stats); err != nil {
			return nil, fmt.Errorf("error unmarshalling stats: %v", err)
		}

		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

func (r *PokemonRepository) DeletePokemonByID(id int) (*models.Pokemon, error) {
	query := `DELETE FROM pokemons WHERE id = $1 RETURNING id, name, height, weight, stats`
	pokemon := &models.Pokemon{}
	var statsJson []byte

	err := r.db.GetDb().QueryRow(query, id).Scan(
		&pokemon.Id,
		&pokemon.Name,
		&pokemon.Height,
		&pokemon.Weight,
		&statsJson,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pokemon with id %d not found", id)
		}
		return nil, fmt.Errorf("error deleting pokemon: %v", err)
	}

	if err := json.Unmarshal(statsJson, &pokemon.Stats); err != nil {
		return nil, fmt.Errorf("error unmarshalling stats: %v", err)
	}

	return pokemon, nil
}
