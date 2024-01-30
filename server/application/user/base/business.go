package base

import (
	"context"

	"github.com/google/uuid"
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
	out = &CreateResponse{
		Name:     in.Name,
		Age:      in.Age,
		Document: in.Document,
	}

	return out, nil
}

// Update - update an user
func Update(ctx context.Context, in *UpdateRequest) error {
	return nil
}

// Delete - delete an user
func Delete(ctx context.Context, in *uuid.UUID) error {
	return nil
}
