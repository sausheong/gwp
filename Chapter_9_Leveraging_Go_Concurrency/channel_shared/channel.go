package main

import (
	"fmt"
	// "math/rand"
	"time"
	"runtime"
)

var DB Store

type Store struct {
	hash map[string]string
	in chan [2]string
	out chan [2]string
}

func StoreInit() {
	DB = Store{
		hash: make(map[string]string),
		in: make(chan [2]string),
	}
	go func() {
		for {
			a := <-DB.in
			DB.hash[a[0]] = a[1]
		}
	}()
}

func (store *Store) Get(key string) (value string, err error) {
	value = store.hash[key]
	return
}

func (store *Store) Add(key string, value string) (err error) {
	a := [2]string{key, value}
	store.in <- a
	// store.hash[key] = value
	return
}

func (store *Store) Set(key string, value string) (err error) {
	
	return
}

func (store *Store) Del(key string) (err error) {
	
	return
}

func (store *Store) Pop(key string) (value string, err error) {
	
	return
}


func main() {
	runtime.GOMAXPROCS(4)
	StoreInit()
	for i := 0; i < 10; i++ {
		go DB.Add("a", "A")
		go DB.Add("a", "B")
		go DB.Add("a", "C")
	
		time.Sleep(1 * time.Microsecond)
	
		s, _ := DB.Get("a")
		fmt.Printf("%s ", s)
		
	}
}
