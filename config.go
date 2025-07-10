package main

import "github.com/manonmission88/PokeGoCli/pokeapi"

// this config holds the runtime state
type config struct {
	PokeClient       pokeapi.Client
	NextLocation     *string
	PreviousLocation *string
	PokeDox          map[string]CaughtPokeMon
}

type CaughtPokeMon struct {
	Name string
}
