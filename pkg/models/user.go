package models

type User struct {
	ID         int64
	Username   string
	ExternalID string
	Role       int
	Language   Language
}
