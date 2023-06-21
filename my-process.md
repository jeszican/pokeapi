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

    b. I now had 3 seperate specs from ChatGPT so I combined them together into one, added a description, url and made some tweaks, (see file `openapi.yaml`):

4) Next I generated the api using the following package and command:

    - https://github.com/deepmap/oapi-codegen
    - `oapi-codegen -package client openapi.yaml > client/client.gen.go`



