package learngo

import "container/list"

type entry struct {
	key   string
	value string
}

type LRUCache struct {
	capacity int
	items    map[string]*list.Element
	order    *list.List
}

func newLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	element, exist := c.items[key]
	if !exist {
		return "", false
	}

	c.order.MoveToFront(element)

	return element.Value.(string), true
}

func (c *LRUCache) Put(key, value string) {
	element, exist := c.items[key]
	if exist {
		element.Value = value
		c.order.MoveToFront(element)
		return
	}

	element = c.order.PushFront(&entry{key: key, value: value})
	c.items[key] = element

	if c.order.Len() == c.capacity {
		back := c.order.Back()
		c.order.Remove(back)
		delete(c.items, back.Value.(string))
	}
}
