package base

import (
	"time"

	"github.com/google/uuid"
)

type CreateRequest struct {
	Name     *string `json:"name"`
	Age      *int    `json:"age"`
	Document *string `json:"document"`
}

type CreateResponse struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Age       *int       `json:"age,omitempty"`
	Document  *string    `json:"document,omitempty"`
}

type UpdateRequest struct {
	ID       *uuid.UUID `json:"-"`
	Name     *string    `json:"name"`
	Age      *int       `json:"age"`
	Document *string    `json:"document"`
}

type FindAllResponse struct {
}

type FindByIDResponse struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Age       *int       `json:"age,omitempty"`
	Document  *string    `json:"document,omitempty"`
}
