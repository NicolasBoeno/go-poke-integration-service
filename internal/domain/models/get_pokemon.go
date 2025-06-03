package models

type ApiPokeStats struct {
	BaseStat int    `json:"base_stat"`
	Effort   int    `json:"effort"`
	StatName string `json:"statName"`
}
type ApiPokemonResponse struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Height int            `json:"height"`
	Weight int            `json:"weight"`
	Stats  []ApiPokeStats `json:"stats"`
}
