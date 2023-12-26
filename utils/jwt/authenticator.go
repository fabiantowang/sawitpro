package jwt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	middleware "github.com/oapi-codegen/echo-middleware"
)

const JWTClaimsContextKey = "jwt_claims"

var (
	ErrNoAuthHeader      = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is malformed")
)

func NewAuthenticator() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticate(ctx, input)
	}
}

func Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	// Get the JWS from the request, to match the request expectations against request contents.
	jws, err := GetJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("getting jws: %w", err)
	}

	// if the JWS is valid, we have a JWT, which will contain a bunch of claims.
	claims, err := GetClaims(jws)
	if err != nil {
		return err
	}

	// Set the property on the echo context so the handler is able to
	// access the claims data we generate in here.
	echoCtx := middleware.GetEchoContext(ctx)
	authdCtx := context.WithValue(echoCtx.Request().Context(), JWTClaimsContextKey, claims)
	requestWithAuth := echoCtx.Request().WithContext(authdCtx)
	echoCtx.SetRequest(requestWithAuth)

	return nil
}

// GetJWSFromRequest extracts a JWS string from an Authorization: Bearer <jws> header
func GetJWSFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")
	// Check for the Authorization header.
	if authHdr == "" {
		return "", ErrNoAuthHeader
	}
	// We expect a header value of the form "Bearer <token>", with 1 space after
	// Bearer, per spec.
	prefix := "Bearer "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}
