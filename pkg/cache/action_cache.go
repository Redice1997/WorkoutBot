package cache

import (
	"time"
	"workout_bot/pkg/models"
)

type ActionCache interface {
	Get(externalID string) (models.UpdatableAction, bool)
	Set(externalID string, action models.UpdatableAction)
	Delete(externalID string)
}

type actionCache struct {
	*Cache
}

func NewActionCache(defaultExpiration, cleanupInterval time.Duration) ActionCache {
	return &actionCache{New(defaultExpiration, cleanupInterval)}
}

func (c *actionCache) Get(externalID string) (models.UpdatableAction, bool) {
	item, found := c.Cache.Get(externalID)
	if !found {
		return nil, false
	}
	action, ok := item.(models.UpdatableAction)
	if !ok {
		return nil, false
	}
	return action, true
}

func (c *actionCache) Set(externalID string, action models.UpdatableAction) {
	c.Cache.Set(externalID, action, 0)
}

func (c *actionCache) Delete(externalID string) {
	c.Cache.Delete(externalID)
}
