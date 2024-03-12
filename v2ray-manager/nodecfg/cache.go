package nodecfg

import "sync"

var nodeCache *Cache            // key为唯一ID，value为*NodeInfo
var userCache map[string]*Cache // cache中，key为email，value为user

func init() {
	nodeCache = &Cache{}
	userCache = make(map[string]*Cache)
}

type Cache struct {
	cache sync.Map
}

func (c *Cache) Store(key string, value interface{}) {
	c.cache.Store(key, value)
}

func (c *Cache) Delete(key string) {
	c.cache.Delete(key)
}

func (c *Cache) Clear() {
	c.cache = sync.Map{}
}

func (c *Cache) Get(key any) (any, bool) {
	return c.cache.Load(key)
}

func (c *Cache) GetAll() map[any]any {
	result := make(map[any]any)
	c.cache.Range(func(key, value any) bool {
		result[key] = value
		return true
	})
	return result
}
