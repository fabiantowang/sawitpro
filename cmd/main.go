package main

import (
	"fmt"
	"os"

	"github.com/fabiantowang/sawitpro/api"
	"github.com/fabiantowang/sawitpro/handler"
	"github.com/fabiantowang/sawitpro/repository"
	"github.com/fabiantowang/sawitpro/utils/jwt"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	middleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	// Clear out the servers array in the swagger spec
	swagger.Servers = nil

	e := echo.New()
	// Use echo validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidatorWithOptions(swagger,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: jwt.NewAuthenticator(),
			},
		}))

	var server api.StrictServerInterface = newServer()
	strictHandler := api.NewStrictHandler(server, nil)

	api.RegisterHandlers(e, strictHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}
