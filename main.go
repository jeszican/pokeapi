package main

import (
	"context"
	"fmt"

	"github.com/jeszican/pokeapi/client"
)

const (
	SERVER_URL = "https://pokeapi.co/api/v2"
)

var (
	TESTCASE1_POKEMON = []string{
		"abra",
		"bulbasaur",
		"charmeleon",
	}
	TESTCASE2_POKEMON = []int{
		7,
		150,
		80,
	}
)

func main() {
    fmt.Println("Starting tests..")
	ctx := context.Background()
	pokeAPIClient, err := client.NewClientWithResponses(SERVER_URL)
	if err != nil {
		panic(fmt.Errorf("error instantiating client: %w", err))
	}
	testCase1(ctx, pokeAPIClient)
}

func testCase1(ctx context.Context, c *client.ClientWithResponses) {
	fmt.Println("Requesting pokemon details..")
	for _, pokemon := range(TESTCASE1_POKEMON) {
		c.GetPokemonDetailsWithResponse(ctx, client.IdPathParam())
	}

}