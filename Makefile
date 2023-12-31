
oapi-codegen-install:
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1
	@oapi-codegen -version

oapi-codegen-generate:
	@oapi-codegen -package client docs/openapi.yaml > client/client.gen.go

integration-tests:
	go run main.go