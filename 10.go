package main

import (
	"fmt"
)

type KVStore struct {
	data map[string]string
}

func NewKVStore() *KVStore {
	return &KVStore{data: make(map[string]string)}
}

func (kv *KVStore) Set(key, value string) {
	kv.data[key] = value
}

func (kv *KVStore) Get(key string) (string, bool) {
	val, ok := kv.data[key]
	return val, ok
}

func (kv *KVStore) Delete(key string) bool {
	_, exists := kv.data[key]
	if exists {
		delete(kv.data, key)
	}
	return exists
}

func (kv *KVStore) Exists(key string) bool {
	_, ok := kv.data[key]
	return ok
}

func main() {
	store := NewKVStore()

	store.Set("name", "Sam")
	fmt.Println(store.Get("name"))

	fmt.Println(store.Exists("age"))

	store.Delete("name")
	fmt.Println(store.Get("name"))
}
