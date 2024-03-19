// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (string, error)
	DeleteAccount(ctx context.Context, id string) error
	GetAccountByEmail(ctx context.Context, email string) (GetAccountByEmailRow, error)
	GetAccountByID(ctx context.Context, id string) (GetAccountByIDRow, error)
	IsEmailAlreadyExists(ctx context.Context, email string) (bool, error)
	ListAccounts(ctx context.Context) ([]ListAccountsRow, error)
	UpdateAccountPasswordByEmail(ctx context.Context, arg UpdateAccountPasswordByEmailParams) error
}

var _ Querier = (*Queries)(nil)
