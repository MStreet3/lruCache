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

func (l *LRUCache) Get(key string) (interface{}, error) {
	dataStore := *l.cache
	data, ok := dataStore[key]
	if ok {
		l.log.DeleteNode(data)
		l.log.SetListHead(data)
		return data.Value, nil
	}
	return nil, errors.New("failed to get item from store")
}

func (l *LRUCache) Set(key string, data interface{}) error {
	dataStore := *l.cache
	_, prs := dataStore[key]
	node := &dbllinkedlist.Node{Key: key, Value: data}
	if l.held < l.capacity {
		dataStore[key] = node
		l.log.SetListHead(node)
		l.held++
		return nil
	} else if l.held == l.capacity && prs {
		dataStore[key] = node
		l.log.DeleteNode(node)
		l.log.SetListHead(node)
		return nil
	} else if l.held == l.capacity && !prs {
		dataStore[key] = node
		keyToRem := l.log.Tail.Key
		delete(dataStore, keyToRem)
		l.log.DeleteListTail()
		l.log.SetListHead(node)
		l.held--
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
