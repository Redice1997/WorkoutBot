package cache

import (
	"strconv"
	"time"
	"workout_bot/pkg/models"
)

type ActionCache interface {
	Get(userID int64) (*models.Action, bool)
	Set(userID int64, action *models.Action)
	Delete(userID int64)
}

type actionCache struct {
	*Cache
}

func NewActionCache(defaultExpiration, cleanupInterval time.Duration) ActionCache {
	return &actionCache{New(defaultExpiration, cleanupInterval)}
}

func (c *actionCache) Get(userID int64) (*models.Action, bool) {
	item, found := c.Cache.Get(strconv.FormatInt(userID, 10))
	if !found {
		return nil, false
	}
	action, ok := item.(*models.Action)
	if !ok {
		return nil, false
	}
	return action, true
}

func (c *actionCache) Set(userID int64, action *models.Action) {
	c.Cache.Set(strconv.FormatInt(userID, 10), action, 0)
}

func (c *actionCache) Delete(userID int64) {
	c.Cache.Delete(strconv.FormatInt(userID, 10))
}
