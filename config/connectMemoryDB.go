package config

import (
	"github.com/hashicorp/golang-lru"
	"log"
)

var StoreMemory = &LRUCache{}

type LRUCache struct {
	Cache *lru.Cache
}

func NewLRUCache() *LRUCache {
	cache, err := lru.New(10)
	if err != nil {
		log.Fatalf("Failed to create cache - Error: %v", err)
	}
	StoreMemory = &LRUCache{Cache: cache}
	return StoreMemory
}
