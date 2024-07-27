package cache

import (
	"time"
	"workout_bot/pkg/models"
)

type UserCache interface {
	Get(externalId string) (*models.User, bool)
	Set(externalId string, user *models.User)
	Delete(externalId string)
}

type userCache struct {
	*Cache
}

func NewUserCache(defaultExpiration, cleanupInterval time.Duration) UserCache {
	return &userCache{New(defaultExpiration, cleanupInterval)}
}

func (c *userCache) Get(externalId string) (*models.User, bool) {
	item, found := c.Cache.Get(externalId)
	if !found {
		return nil, false
	}
	user, ok := item.(*models.User)
	if !ok {
		return nil, false
	}
	return user, true
}

func (c *userCache) Set(externalId string, user *models.User) {
	c.Cache.Set(externalId, user, 0)
}

func (c *userCache) Delete(externalId string) {
	c.Cache.Delete(externalId)
}
