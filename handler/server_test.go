package handler

import (
	"net/http"
	"testing"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/fabiantowang/sawitpro/api"
	repo "github.com/fabiantowang/sawitpro/repository"
	"github.com/fabiantowang/sawitpro/utils/jwt"
	middleware "github.com/oapi-codegen/echo-middleware"
	"github.com/oapi-codegen/testutil"
)

func TestServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	swagger, _ := api.GetSwagger()
	swagger.Servers = nil

	e := echo.New()
	e.Use(middleware.OapiRequestValidatorWithOptions(swagger,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: jwt.NewAuthenticator(),
			},
		}))
	mockRepo := repo.NewMockRepositoryInterface(ctrl)
	var server api.StrictServerInterface = NewServer(NewServerOptions{Repository: mockRepo})
	strictHandler := api.NewStrictHandler(server, nil)
	api.RegisterHandlers(e, strictHandler)

	// GET /profile should return 403 forbidden without credentials
	response := testutil.NewRequest().Get("/profile").GoWithHTTPHandler(t, e)
	assert.Equal(t, http.StatusForbidden, response.Code())

	// POST /register failed due to validation
	response = testutil.NewRequest().Post("/register").
		WithAcceptJson().
		WithJsonBody(api.RegisterRequest{
			Phone:    "123",
			Password: "123456Ac!",
			Fullname: "John Smith",
		}).GoWithHTTPHandler(t, e)
	assert.Equal(t, http.StatusBadRequest, response.Code())

	// GET /login failed due to no data
	response = testutil.NewRequest().Post("/login").
		WithAcceptJson().
		WithJsonBody(api.LoginRequest{Phone: "123", Password: "123"}).GoWithHTTPHandler(t, e)
	require.Equal(t, http.StatusBadRequest, response.Code())
}
