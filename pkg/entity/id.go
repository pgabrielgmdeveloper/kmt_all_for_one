package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func ParseID(id string) (ID, error) {
	if id == "" {
		return uuid.Nil, nil
	}
	return uuid.Parse(id)
}

func IsValidID(id ID) bool {
	return id == uuid.Nil
}
