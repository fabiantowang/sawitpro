// This file contains types that are used in the repository layer.
package repository

import "github.com/google/uuid"

type AddUserInput struct {
	Phone    string
	Fullname string
	Salt     string
	Password string
}

type AddUserOutput struct {
	Id uuid.UUID
}

type GetUserByIdOutput struct {
	Phone    string
	Fullname string
}

type GetUserByPhoneOutput struct {
	Id       uuid.UUID
	Salt     string
	Password string
}

type UpdateUserInput struct {
	Phone    string
	Fullname string
	Id       uuid.UUID
}
