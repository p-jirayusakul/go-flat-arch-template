// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg *CreateAccountParams) (string, error)
	CreateAddresses(ctx context.Context, arg *CreateAddressesParams) (string, error)
	DeleteAccount(ctx context.Context, id string) error
	DeleteAddressesById(ctx context.Context, id string) error
	GetAccountByEmail(ctx context.Context, email string) (*GetAccountByEmailRow, error)
	GetAccountByID(ctx context.Context, id string) (*GetAccountByIDRow, error)
	GetAddressById(ctx context.Context, id string) (*GetAddressByIdRow, error)
	IsAccountAlreadyExists(ctx context.Context, id string) (bool, error)
	IsAddressesAlreadyExists(ctx context.Context, arg *IsAddressesAlreadyExistsParams) (bool, error)
	IsEmailAlreadyExists(ctx context.Context, email string) (bool, error)
	ListAccounts(ctx context.Context) ([]*ListAccountsRow, error)
	ListAddresses(ctx context.Context) ([]*ListAddressesRow, error)
	ListAddressesByAccountId(ctx context.Context, accountsID string) ([]*ListAddressesByAccountIdRow, error)
	UpdateAccountPasswordByEmail(ctx context.Context, arg *UpdateAccountPasswordByEmailParams) error
	UpdateAddressById(ctx context.Context, arg *UpdateAddressByIdParams) error
}

var _ Querier = (*Queries)(nil)
