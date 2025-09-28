package jsonstore_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/viktor-titov/bookmarks/internal/adapter/jsonstore"
)

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestNewStore_CreatesFileIfNotExists(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.json")

	store, err := jsonstore.NewStore[TestData](path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// файл не должен создаться
	if _, err := os.Stat(path); os.IsExist(err) {
		t.Fatalf("unexpected file %s to exist", path)
	}

	// данные должны быть пустыми
	data := store.Get()
	if data != (TestData{}) {
		t.Errorf("expected empty data, got %+v", data)
	}
}

func TestNewStore_LoadsExistingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.json")

	expected := TestData{Name: "Alice", Age: 30}
	bytes, _ := json.Marshal(expected)
	if err := os.WriteFile(path, bytes, 0644); err != nil {
		t.Fatalf("failed to prepare file: %v", err)
	}

	store, err := jsonstore.NewStore[TestData](path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data := store.Get()
	if data != expected {
		t.Errorf("expected %+v, got %+v", expected, data)
	}
}

func TestStore_SetAndGet(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.json")

	store, err := jsonstore.NewStore[TestData](path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	newData := TestData{Name: "Bob", Age: 42}
	if err := store.Set(newData); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := store.Get()
	if got != newData {
		t.Errorf("expected %+v, got %+v", newData, got)
	}
}

func TestStore_Persistence(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.json")

	store, _ := jsonstore.NewStore[TestData](path)
	data := TestData{Name: "Charlie", Age: 25}
	if err := store.Set(data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// пересоздаём Store
	store2, err := jsonstore.NewStore[TestData](path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := store2.Get()
	if got != data {
		t.Errorf("expected %+v, got %+v", data, got)
	}
}

func TestStore_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.json")

	// пишем некорректный JSON
	if err := os.WriteFile(path, []byte("{invalid json"), 0644); err != nil {
		t.Fatalf("failed to prepare file: %v", err)
	}

	_, err := jsonstore.NewStore[TestData](path)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}
