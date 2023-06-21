package main

import (
	"context"
	"net/http"
	"os"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jeszican/pokeapi/client"
)

const (
	SERVER_URL = "https://pokeapi.co/api/v2"
)

var (
	failures         = []string{}
	pokemonTestCases = []testCases{
		{
			name: "valid pokemon by name",
			input: []string{
				"abra",
				"bulbasaur",
				"charmeleon",
			},
			wantErr:        false,
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid pokemon by name - 404",
			input: []string{
				"fakepokemon",
				"chappie",
				"mao",
			},
			wantErr:        false,
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "valid pokemon by id",
			input: []string{
				"17",
				"150",
				"221",
			},
			wantErr:        false,
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid pokemon by id - 404",
			input: []string{
				"20000",
				"XXX",
			},
			wantErr:        false,
			wantStatusCode: http.StatusNotFound,
		},
	}
	natureTestCases = []testCases{
		{
			name: "valid nature by name",
			input: []string{
				"bold",
				"sassy",
				"hasty",
			},
			wantErr:        false,
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid nature by name",
			input: []string{
				"jessica",
				"likes",
				"cats",
			},
			wantErr:        false,
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "valid nature by id",
			input: []string{
				"10",
				"1",
				"11",
			},
			wantErr:        false,
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid nature by id",
			input: []string{
				"99999999999",
				"11111111111",
			},
			wantErr:        false,
			wantStatusCode: http.StatusNotFound,
		},
	}
	statTestCases = []testCases{
		{
			name: "valid stat by name",
			input: []string{
				"speed",
				"hp",
				"special-defense",
			},
			wantErr:        false,
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid stat by name",
			input: []string{
				"sad",
				"integration",
				"testing",
			},
			wantErr:        false,
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "valid stat by id",
			input: []string{
				"1",
				"2",
				"5",
			},
			wantErr:        true,
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid stat by id",
			input: []string{
				"0000000000",
			},
			wantErr:        false,
			wantStatusCode: http.StatusNotFound,
		},
	}
)

type testCases struct {
	name           string
	input          []string
	wantErr        bool
	wantStatusCode int
}

func main() {
	fmt.Println("Starting tests..")
	ctx := context.Background()
	validator := validator.New()
	cli, err := client.NewClientWithResponses(SERVER_URL)
	if err != nil {
		panic(fmt.Errorf("error instantiating client: %w", err))
	}

	for _, testCase := range pokemonTestCases {
		performPokemonTest(ctx, cli, validator, testCase.name, testCase.input, testCase.wantErr, testCase.wantStatusCode)
	}
	for _, testCase := range natureTestCases {
		performNatureTest(ctx, cli, validator, testCase.name, testCase.input, testCase.wantErr, testCase.wantStatusCode)
	}
	for _, testCase := range statTestCases {
		performStatTest(ctx, cli, validator, testCase.name, testCase.input, testCase.wantErr, testCase.wantStatusCode)
	}
	// if there has been even 1 failure, panic
	if len(failures) > 0 {
		for _, f := range failures {
			fmt.Println(f)
		}
		os.Exit(2)
	}
}

// TODO: Combine functions that call the test functions
func performPokemonTest(ctx context.Context, c *client.ClientWithResponses, v *validator.Validate, name string, input []string, wantErr bool, wantStatusCode int) {
	for _, item := range input {
		idPathParam := client.IdPathParam(item)
		response, err := c.GetPokemonDetailsWithResponse(ctx, idPathParam)

		if response.StatusCode() != wantStatusCode {
			if wantErr {
				fmt.Printf("PASS: Test '%s' with input '%s' - Expected status code '%d' and got status code '%d'\n", name, item, wantStatusCode, response.StatusCode())
			} else {
				failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected status code '%d' but got status code '%d'", name, item, wantStatusCode, response.StatusCode()))
			}
			continue
		}

		if err != nil {
			if wantErr {
				fmt.Printf("PASS: Test '%s' with input '%s' - Expected error and got '%v'\n", name, item, err)
			} else {
				failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected no error but got '%v'", name, item, err))
			}
			continue
		}

		if response.JSON200 != nil {
			err := v.Struct(response.JSON200)
			if err != nil {
				if wantErr {
					fmt.Printf("PASS: Test '%s' with input '%s' - Expected error and got '%v'\n", name, item, err)
				} else {
					failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected no error but got '%v'", name, item, err))
				}
				continue
			}
		}
		if !wantErr {
			fmt.Printf("PASS: Test '%s' with input '%s'\n", name, item)
		} else {
			failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected error but got none", name, item))
		}

	}
}

func performStatTest(ctx context.Context, c *client.ClientWithResponses, v *validator.Validate, name string, input []string, wantErr bool, wantStatusCode int) {
	for _, item := range input {
		idPathParam := client.IdPathParam(item)
		response, err := c.GetStatDetailsWithResponse(ctx, idPathParam)

		if response.StatusCode() != wantStatusCode {
			if wantErr {
				fmt.Printf("PASS: Test '%s' with input '%s' - Expected status code '%d' and got status code '%d'\n", name, item, wantStatusCode, response.StatusCode())
			} else {
				failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected status code '%d' but got status code '%d'", name, item, wantStatusCode, response.StatusCode()))
			}
			continue
		}

		if err != nil {
			if wantErr {
				fmt.Printf("PASS: Test '%s' with input '%s' - Expected error and got '%v'\n", name, item, err)
			} else {
				failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected no error but got '%v'", name, item, err))
			}
			continue
		}

		if response.JSON200 != nil {
			err := v.Struct(response.JSON200)
			if err != nil {
				if wantErr {
					fmt.Printf("PASS: Test '%s' with input '%s' - Expected error and got '%v'\n", name, item, err)
				} else {
					failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected no error but got '%v'", name, item, err))
				}
				continue
			}
		}
		if !wantErr {
			fmt.Printf("PASS: Test '%s' with input '%s'\n", name, item)
		} else {
			failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected error but got none", name, item))
		}

	}
}

func performNatureTest(ctx context.Context, c *client.ClientWithResponses, v *validator.Validate, name string, input []string, wantErr bool, wantStatusCode int) {
	for _, item := range input {
		idPathParam := client.IdPathParam(item)
		response, err := c.GetNatureDetailsWithResponse(ctx, idPathParam)

		if response.StatusCode() != wantStatusCode {
			if wantErr {
				fmt.Printf("PASS: Test '%s' with input '%s' - Expected status code '%d' and got status code '%d'\n", name, item, wantStatusCode, response.StatusCode())
			} else {
				failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected status code '%d' but got status code '%d'", name, item, wantStatusCode, response.StatusCode()))
			}
			continue
		}

		if err != nil {
			if wantErr {
				fmt.Printf("PASS: Test '%s' with input '%s' - Expected error and got '%v'\n", name, item, err)
			} else {
				failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected no error but got '%v'", name, item, err))
			}
			continue
		}

		if response.JSON200 != nil {
			err := v.Struct(response.JSON200)
			if err != nil {
				if wantErr {
					fmt.Printf("PASS: Test '%s' with input '%s' - Expected error and got '%v'\n", name, item, err)
				} else {
					failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected no error but got '%v'", name, item, err))
				}
				continue
			}
		}
		if !wantErr {
			fmt.Printf("PASS: Test '%s' with input '%s'\n", name, item)
		} else {
			failures = append(failures, fmt.Sprintf("FAIL: Test '%s' with input '%s' - Expected error but got none", name, item))
		}

	}
}
