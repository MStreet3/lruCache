package main

import (
	"errors"
	"fmt"

	"lruCache/dbllinkedlist"
)

type LRUCache struct {
	capacity int
	held     int
	log      *dbllinkedlist.DblLinkedList
	cache    *map[string]*dbllinkedlist.Node
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		capacity: cap,
		log:      &dbllinkedlist.DblLinkedList{},
		cache:    &map[string]*dbllinkedlist.Node{},
	}
}

func (l *LRUCache) Get(key string) (interface{}, bool) {
	store := *l.cache
	node, ok := store[key]
	if ok {
		/* bump key to top of list */
		l.log.DeleteNode(node)
		l.log.SetListHead(node)

		/* return data held by the node */
		return node.Value, ok
	}
	return nil, false
}

func (l *LRUCache) Set(key string, data interface{}) error {
	store := *l.cache
	_, prs := store[key]
	node := &dbllinkedlist.Node{Key: key, Value: data}
	if l.held < l.capacity {
		store[key] = node
		l.log.SetListHead(node)
		l.held++
		return nil
	} else if l.held == l.capacity && prs {
		store[key] = node
		l.log.DeleteNode(node)
		l.log.SetListHead(node)
		return nil
	} else if l.held == l.capacity && !prs {
		store[key] = node
		delete(store, l.log.Tail.Key)
		l.log.DeleteListTail()
		l.log.SetListHead(node)
		return nil
	}
	return errors.New("improper cache state")
}

func main() {
	cache := NewLRUCache(2)
	cache.Set("first", 1)
	cache.Set("next", 2)
	cache.Set("third", 3)
	item, err := cache.Get("first")
	fmt.Println(item)
	fmt.Println(err)
	item, err = cache.Get("next")
	fmt.Println(item)
	fmt.Println(err)
}
