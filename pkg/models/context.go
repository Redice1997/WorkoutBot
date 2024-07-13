package models

import "context"

type ContextKey int

type UserContext interface {
	context.Context
	ExternalID() string
	Language() Language
	Username() string
}

type userContext struct {
	context.Context
	externalID string
	language   Language
	username   string
}

func NewUserContext(ctx context.Context, externalID, username string, language Language) UserContext {
	return &userContext{
		Context:    ctx,
		externalID: externalID,
		username:   username,
		language:   language,
	}
}

func (c *userContext) ExternalID() string {
	return c.externalID
}

func (c *userContext) Language() Language {
	return c.language
}

func (c *userContext) Username() string {
	return c.username
}
