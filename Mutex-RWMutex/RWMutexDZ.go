package main

import "sync"

type Storage struct {
	data map[string]string
	mtx  sync.RWMutex
}

func NewStorage() Storage {
	return Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Set(key string, value string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.data[key] = value
}

func (s *Storage) Get(key string) (string, bool) {

	s.mtx.RLock()
	defer s.mtx.RUnlock()

	value, ok := s.data[key]

	return value, ok

}

func main() {
	wg := &sync.WaitGroup{}

	storage := NewStorage()

	storage.Set("age", "21")
	storage.Set("name", "Alex")
	storage.Set("city", "Moscow")
	storage.Set("hobby", "tennis")
	wg.Add(1)
	go func() {
		defer wg.Done()
		storage.Get("name")
		storage.Get("city")
		storage.Get("age")
		storage.Get("hobby")
	}()

	wg.Wait()

}
