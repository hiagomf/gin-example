package base

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiagomf/gin-example/database"
	"github.com/hiagomf/gin-example/domain/user/base"
)

// FindAll - find all users
func FindAll(ctx context.Context) error {
	return nil
}

// FindByID - find user by ID
func FindByID(ctx context.Context, in *uuid.UUID) (out *FindByIDResponse, err error) {
	return nil, nil
}

// Create - create an user
func Create(ctx context.Context, in *CreateRequest) (out *CreateResponse, err error) {
	tx, err := database.NewTransaction(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	repository := base.NewRepository(ctx, tx)

	usr := &base.User{
		Name:     in.Name,
		Age:      in.Age,
		Document: in.Document,
	}
	if err := repository.Insert(usr); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &CreateResponse{
		ID:        usr.ID,
		CreatedAt: usr.CreatedAt,
		Name:      usr.Name,
		Age:       usr.Age,
		Document:  usr.Document,
	}, nil
}

// Update - update an user
func Update(ctx context.Context, in *UpdateRequest) error {
	return nil
}

// Delete - delete an user
func Delete(ctx context.Context, in *uuid.UUID) error {
	return nil
}
