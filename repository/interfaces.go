// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
)

type RepositoryInterface interface {
	AddUser(ctx context.Context, input AddUserInput) (output AddUserOutput, err error)
	GetUserById(ctx context.Context, id string) (output GetUserByIdOutput, err error)
	GetUserByPhone(ctx context.Context, phone string) (output GetUserByPhoneOutput, err error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (err error)
	IncrementSuccessfulLogin(ctx context.Context, id string) (err error)
}
