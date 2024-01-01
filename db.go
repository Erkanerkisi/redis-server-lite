package main

import (
	"sync"
	"time"
)

var storage = Storage{db: make(map[string]interface{}), mu: sync.Mutex{}}

type Storage struct {
	mu sync.Mutex
	db map[string]interface{}
}

type RedisObj struct {
	Data interface{}
	Typ  string
	TTL  time.Time
}

func (storage *Storage) set(key string, value string) {
	time.Now()
	storage.mu.Lock()
	storage.db[key] = value
	defer storage.mu.Unlock()
}

func (storage *Storage) setArray(key string, value []interface{}) {
	storage.mu.Lock()
	storage.db[key] = value
	defer storage.mu.Unlock()
}

func (storage *Storage) get(key string) interface{} {
	if val, ok := storage.db[key]; ok {
		if IsArray(val) {
			return val.([]interface{})
		}
		return &val
	}
	return nil
}

func (storage *Storage) delete(key string) bool {
	if storage.exists(key) {
		delete(storage.db, key)
		return true
	} else {
		return false
	}
}

func (storage *Storage) exists(key string) bool {
	if _, ok := storage.db[key]; ok {
		return true
	}
	return false
}

func GetStorage() *Storage {
	return &storage
}
