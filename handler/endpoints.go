package handler

import (
	"context"
	"encoding/base64"

	"github.com/fabiantowang/sawitpro/api"
	"github.com/fabiantowang/sawitpro/repository"
	"github.com/fabiantowang/sawitpro/utils/db"
	"github.com/fabiantowang/sawitpro/utils/hash"
	"github.com/fabiantowang/sawitpro/utils/jwt"
	"github.com/fabiantowang/sawitpro/utils/validator"
	"github.com/google/uuid"
)

// This is endpoint for user login.
// (POST /login)
func (s *Server) Login(ctx context.Context, request api.LoginRequestObject) (api.LoginResponseObject, error) {
	userInfo, err := s.Repository.GetUserByPhone(ctx, request.Body.Phone)
	if err != nil {
		return nil, err
	}

	// get salt bytes
	salt, err := base64.RawStdEncoding.Strict().DecodeString(userInfo.Salt)
	if err != nil {
		return nil, err
	}

	// compare password
	if !hash.PasswordsMatch(userInfo.Password, request.Body.Password, salt) {
		return api.Login400Response(api.Login400Response{}), nil
	}

	// get jwt
	jwtString, err := jwt.GenerateToken(userInfo.Id.String())
	if err != nil {
		return nil, err
	}

	// update successful login counter
	err = s.Repository.IncrementSuccessfulLogin(ctx, userInfo.Id.String())
	if err != nil {
		return nil, err
	}

	var result api.LoginResponse
	result.Jwt = jwtString
	result.Userid = userInfo.Id

	return api.Login200JSONResponse(result), nil
}

// This is endpoint for getting user profile.
// (GET /profile)
func (s *Server) ProfileGet(ctx context.Context, request api.ProfileGetRequestObject) (api.ProfileGetResponseObject, error) {
	// get claims from token
	claims, err := jwt.GetClaimsCtx(ctx)

	if err != nil {
		return nil, err
	}

	userInfo, err := s.Repository.GetUserById(ctx, claims.Userid)
	if err != nil {
		return nil, err
	}

	var result api.ProfileResponse
	result.Fullname = userInfo.Fullname
	result.Phone = userInfo.Phone

	return api.ProfileGet200JSONResponse(result), nil
}

// This is endpoint for updating user profile.
// (PUT /profile)
func (s *Server) ProfileUpdate(ctx context.Context, request api.ProfileUpdateRequestObject) (api.ProfileUpdateResponseObject, error) {
	var valErrs api.ErrorResponse

	errs := validator.ValidateProfileUpdate(request.Body.Phone, request.Body.Fullname)
	for _, err := range errs {
		var structErr struct {
			Message string `json:"message"`
		}

		structErr.Message = err.Error()
		valErrs = append(valErrs, structErr)
	}
	if len(valErrs) > 0 {
		return api.ProfileUpdate400JSONResponse(valErrs), nil
	}

	// get claims from token
	claims, err := jwt.GetClaimsCtx(ctx)

	if err != nil {
		return nil, err
	}

	userInfo, err := s.Repository.GetUserById(ctx, claims.Userid)
	if err != nil {
		return nil, err
	}

	var updateUserInfo repository.UpdateUserInput
	updateUserInfo.Id = uuid.MustParse(claims.Userid)
	updateUserInfo.Phone = userInfo.Phone
	if request.Body.Phone != nil {
		updateUserInfo.Phone = *request.Body.Phone
	}
	updateUserInfo.Fullname = userInfo.Fullname
	if request.Body.Fullname != nil {
		updateUserInfo.Fullname = *request.Body.Fullname
	}

	err = s.Repository.UpdateUser(ctx, updateUserInfo)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return api.ProfileUpdate409Response(api.ProfileUpdate409Response{}), nil
		}
		return nil, err
	}

	return api.ProfileUpdate200Response(api.ProfileUpdate200Response{}), nil
}

// This is endpoint for user registration.
// (POST /register)
func (s *Server) Register(ctx context.Context, request api.RegisterRequestObject) (api.RegisterResponseObject, error) {
	var valErrs api.ErrorResponse

	errs := validator.ValidateNewUser(request.Body.Phone, request.Body.Fullname, request.Body.Password)
	for _, err := range errs {
		var structErr struct {
			Message string `json:"message"`
		}

		structErr.Message = err.Error()
		valErrs = append(valErrs, structErr)
	}
	if len(valErrs) > 0 {
		return api.Register400JSONResponse(valErrs), nil
	}

	salt, err := hash.GenerateRandomSalt()
	if err != nil {
		return nil, err
	}

	var userInfo repository.AddUserInput
	userInfo.Phone = request.Body.Phone
	userInfo.Fullname = request.Body.Fullname
	userInfo.Salt = base64.RawStdEncoding.EncodeToString(salt)
	userInfo.Password = hash.HashPassword(request.Body.Password, salt)

	userOutput, err := s.Repository.AddUser(ctx, userInfo)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			var structErr struct {
				Message string `json:"message"`
			}

			structErr.Message = "Phone number already exists"
			valErrs = append(valErrs, structErr)
			return api.Register400JSONResponse(valErrs), nil
		}
		return nil, err
	}

	var result api.RegisterResponse
	result.Userid = userOutput.Id
	return api.Register200JSONResponse(result), nil
}
