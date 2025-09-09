package events

import "github.com/google/uuid"

func getUniqueID() string {
	return uuid.New().String()
}
