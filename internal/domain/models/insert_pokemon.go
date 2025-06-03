package models

type PokemonStats struct {
	Name     string `json:"name"`
	BaseStat int    `json:"base_stat"`
	Effort   int    `json:"effort"`
}

type Pokemon struct {
	Status string
	Id     int            `json:"id" db:"id"`
	Name   string         `json:"name" db:"name"`
	Height int            `json:"height" db:"height"`
	Weight int            `json:"weight" db:"weight"`
	Stats  []PokemonStats `json:"stats" db:"stats"`
}
