package jsonstore

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Store хранит данные из JSON-файла
type Store[T any] struct {
	filePath string
	data     T
	mu       sync.RWMutex
}

// NewStore создаёт новый Store, загружая данные из файла
func NewStore[T any](filePath string) (*Store[T], error) {
	var data T
	store := &Store[T]{
		filePath: filePath,
		data:     data,
	}

	err := store.load()
	if err != nil {
		return nil, fmt.Errorf("load data from file: %w", err)
	}

	return store, nil
}

// load загружает данные из файла
func (s *Store[T]) load() error {
	_, err := os.Stat(s.filePath)
	if os.IsNotExist(err) {
		return nil
	}

	bytes, err := os.ReadFile(s.filePath)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	err = json.Unmarshal(bytes, &s.data)
	if err != nil {
		return fmt.Errorf("unmarshal json: %w", err)
	}

	return nil
}

// List возвращает все ключи и значения
func (s *Store[T]) Get() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data
}

func (s *Store[T]) Set(data T) error {
	s.data = data
	return s.save()
}

// save сохраняет данные в файл
func (s *Store[T]) save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, err := os.Stat(s.filePath)
	if os.IsNotExist(err) {
		_, err := s.createJson()
		if err != nil {
			return fmt.Errorf("create json file: %w", err)
		}
	}

	bytes, err := json.MarshalIndent(s.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, bytes, 0644)
}

func (s *Store[T]) createJson() (*os.File, error) {
	dir := filepath.Dir(s.filePath)

	// Создаём все подкаталоги (если они уже есть, ошибки не будет)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("create a directory: %w", err)
	}

	// Создаём файл (если он уже есть, он будет перезаписан)
	f, err := os.Create(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("create a file: %w", err)
	}

	return f, nil
}
