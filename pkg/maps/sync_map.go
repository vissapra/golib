package maps

import "sync"

type ThreadSafeMap struct {
	lock   *sync.RWMutex
	config map[string]string
}

func (c ThreadSafeMap) Get(key string) string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	value, _ := c.config[key]
	return value
}

func (c *ThreadSafeMap) Put(key string, val string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.config[key] = val
}

func (c *ThreadSafeMap) Remove(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.config, key)
}

func NewThreadSafeMap() *ThreadSafeMap {
	return &ThreadSafeMap{config: make(map[string]string), lock: new(sync.RWMutex)}

}
