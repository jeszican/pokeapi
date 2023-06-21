## How to use this SDK
### Installation

To use this SDK in your go code, go get it.
`go get github.com/jeszican/pokeapi/client `

You can then instantiate the client like this:
```golang
    import (
        "github.com/jeszican/pokeapi/client"
    )
    const (
	    SERVER_URL = "https://pokeapi.co/api/v2"
    )
    cli, err := client.NewClientWithResponses(SERVER_URL)
```

### Documentation

You can view the openapi specification for this API [here](https://jeszican.github.io/pokeapi/).

### Testing

You can use the generated mock in the mocks/ folder for testing this API client. You can use it like:

```golang
package mocks

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	client "github.com/jeszican/pokeapi/client"
	v4 "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMockClientInterface_GetNatureDetails(t *testing.T) {
    var natureId = "x"
    
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockClientInterface(ctrl)
	ctx := context.TODO()
	id := client.IdPathParam(natureId)

	expectedResponse := &http.Response{}
	expectedError := assert.AnError

	mockClient.EXPECT().
		GetNatureDetails(ctx, id).
		Return(expectedResponse, expectedError)

	response, err := mockClient.GetNatureDetails(ctx, id)

	assert.Equal(t, expectedResponse, response)
	assert.Equal(t, expectedError, err)
}
```

---
## The Process
### Introduction

This document is intended to provide an overview of the process i went through to generate the sdk for [pokeapi](https://pokeapi.co/).

### Steps

1) My first step was to understand more about the PokeAPI, I wanted to understand the purpose of the API and the general format of the responses. I also went digging around to see if I could find any documentation on the structure of the API, things like JSON schemas that might describe the responses, or OpenAPI documentation that would describe the API as a whole.

2) When I didn't find anything my first step was to start generating an OpenAPI document myself, (see file `/drafts/openapi.yaml`)

    a. I started with the `/pokemon/{id or name}` endpoint and added a paginated result schema, and a pokemon schema. I found that some fields were missing for bulbasaur, the first pokemon, so I used a couple of other examples, such as Azurill to fill in the blanks, more specifically the `past_types` field.

3) Once I'd done the `/pokemon/{id or name}` endpoint, which took a little bit of time, I decided to try and get ChatGPT to generate the specification for me instead. It took a couple of goes to get the prompt correct, but this is the one that gave me the best output:

    ```
    Act as a programmatic api generator and create an openapi yaml 3.0.0 specification for the https://pokeapi.co/ api, and more specifically the /pokemon/{id}/ endpoint.

    When creating the output, you should make sure to fully expand and document all of the fields, you must provide examples and descriptions for all fields, this is what a JSON response from this end looks like:

    <cut down json response here>
    ```

    a. Once I had a working prompt, I asked for the other two endpoints as specs:
    - `/nature/{id}`
    - `/stat/{id}`

    b. I now had 3 seperate specs from ChatGPT so I combined them together into     one, added a description, url and made some tweaks, (see file `docs/openapi.yaml`):

4) Next I generated the api using the following package and command:

    - https://github.com/deepmap/oapi-codegen
    - `oapi-codegen -package client openapi.yaml > client/client.gen.go`

5) After this I decided to try and create a github action that regenerated my client based on changes and pushed to main. _I don't have experience creating github actions so thought this might be fun._

6) Now I am going to start on the integration tests, my plan is to:

    a) Create a go binary that will import the client, and use it.

    b) Create a github action that will run the go binary and fail/succeed.

    c) If I have time, I will create some unit tests also.

7) Integration tests are done, added some extra information, generated some mocks and tried to refine the regeneration of openapi workflow but failed. :)

8) & at about 4 hours in I've stopped. Things I would like to improve with more time include:

    a) I'd like to combine the functions that run the tests into one by interfacing the response types.

    b) I would like to change the generate-oapi workflow to only generate if the openapi spec has changed.

    c) I'd like to add some unit testing.

    d) I think I could add more integration tests and spend more time on the required fields.
    



