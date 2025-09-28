// Package command содержит реализацию репозитория для управления командами.
package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/viktor-titov/bookmarks/internal/adapter/config"
	"github.com/viktor-titov/bookmarks/internal/adapter/jsonstore"
	"github.com/viktor-titov/bookmarks/internal/models"
)

type CommandRepositoryInterface interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
	List() (map[string]string, error)
}

type CommandRepository struct {
	store *jsonstore.Store[models.Commands]
}

var _ CommandRepositoryInterface = (*CommandRepository)(nil)

func NewCommandRepository(cfg config.ConfigRepository) (*CommandRepository, error) {
	path, err := cfg.Get("pathCommands")
	if err != nil {
		return nil, fmt.Errorf("get path from config: %w", err)
	}

	store, err := jsonstore.NewStore[models.Commands](expandPath(path))
	if err != nil {
		return nil, fmt.Errorf("create commands store: %w", err)
	}

	return &CommandRepository{
		store: store,
	}, nil
}

func (cr *CommandRepository) Get(key string) (string, error) {
	data := cr.store.Get()
	value, ok := data.Commands[key]
	if !ok {
		return "", fmt.Errorf("no such command: %s", key)
	}
	return value, nil
}

func (cr *CommandRepository) Set(key, value string) error {
	data := cr.store.Get()
	if data.Commands == nil {
		data.Commands = make(map[string]string)
	}
	data.Commands[key] = value
	cr.store.Set(data)

	return nil
}

func (cr *CommandRepository) Delete(key string) error {
	data := cr.store.Get()
	if data.Commands == nil {
		return fmt.Errorf("no such command: %s", key)
	}
	if _, ok := data.Commands[key]; !ok {
		return fmt.Errorf("no such command: %s", key)
	}
	delete(data.Commands, key)
	cr.store.Set(data)

	return nil
}

func (cr *CommandRepository) List() (map[string]string, error) {
	data := cr.store.Get()
	if data.Commands == nil {
		return make(map[string]string), nil
	}
	return data.Commands, nil
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}

		if path == "~" {
			path = home
		} else if strings.HasPrefix(path, "~/") {
			path = filepath.Join(home, path[2:])
		}
	}

	p, err := filepath.Abs(path)
	if err != nil {
		return path
	}

	return p
}
