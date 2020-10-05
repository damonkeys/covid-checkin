package database

type (
	// CRUD is the typical CREATE READ UPDATE DELETE interface for an entity
	CRUD interface {
		Create() error
		Update() error
		Delete() error
		Undelete() error
	}
)
