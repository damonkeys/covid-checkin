package database

import "context"

// Model defines standard database model funcs for all models
type Model interface {
	Create(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context, value interface{}) error
	Undelete(ctx context.Context, value interface{}) error
}
