## How to use this SDK
### Installation

To use this SDK in your go code, go get it.
`go get github.com/jeszican/pokeapi/client `

You can then instantiate the client & call the functions like this:
```golang
    import (
        "github.com/jeszican/pokeapi/client"
    )
    const (
	    SERVER_URL = "https://pokeapi.co/api/v2"
    )
    cli, err := client.NewClientWithResponses(SERVER_URL)
    response, err := cli.GetPokemonDetailsWithResponse(ctx, idPathParam)
    // check the err
    // the response.JSON200 is a *client.PokemonDetails
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


