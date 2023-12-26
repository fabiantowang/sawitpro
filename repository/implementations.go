package repository

import (
	"context"
)

func (r *Repository) AddUser(ctx context.Context, input AddUserInput) (output AddUserOutput, err error) {
	query := "INSERT INTO users (phone, fullname, salt, password) VALUES($1, $2, $3, $4) RETURNING (id)"
	err = r.Db.QueryRow(ctx, query, input.Phone, input.Fullname, input.Salt, input.Password).Scan(&output.Id)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetUserById(ctx context.Context, id string) (output GetUserByIdOutput, err error) {
	query := "SELECT phone, fullname FROM users WHERE id = $1"
	err = r.Db.QueryRow(ctx, query, id).Scan(&output.Phone, &output.Fullname)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetUserByPhone(ctx context.Context, phone string) (output GetUserByPhoneOutput, err error) {
	query := "SELECT id, salt, password FROM users WHERE phone = $1"
	err = r.Db.QueryRow(ctx, query, phone).Scan(&output.Id, &output.Salt, &output.Password)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateUser(ctx context.Context, input UpdateUserInput) (err error) {
	query := "UPDATE users SET phone = $1, fullname = $2 WHERE id = $3"
	_, err = r.Db.Exec(ctx, query, input.Phone, input.Fullname, input.Id)
	if err != nil {
		return
	}
	return
}

func (r *Repository) IncrementSuccessfulLogin(ctx context.Context, id string) (err error) {
	query := "UPDATE users SET successful_login = successful_login + 1 WHERE id = $1"
	_, err = r.Db.Exec(ctx, query, id)
	if err != nil {
		return
	}
	return
}
