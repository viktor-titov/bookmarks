// Package config содержит реализацию репозитория для работы с конфигурацией.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/viktor-titov/bookmarks/constants"
	"github.com/viktor-titov/bookmarks/internal/adapter/jsonstore"
	"github.com/viktor-titov/bookmarks/internal/models"
)

type ConfigRepositoryInterface interface {
	Get(key string) (string, error)
	Set(key, value string) error
	List() map[string]string
}

var _ ConfigRepositoryInterface = (*ConfigRepository)(nil) // проверка реализации интерфейса на этапе компиляции

type ConfigRepository struct {
	store *jsonstore.Store[models.Config]
}

func NewConfigRepository() (*ConfigRepository, error) {
	configPath := expandPath(constants.DefaultPathConfig)
	store, err := jsonstore.NewStore[models.Config](configPath)
	if err != nil {
		return nil, fmt.Errorf("create config store: %w", err)
	}

	return &ConfigRepository{
		store: store,
	}, nil
}

func (cr *ConfigRepository) List() map[string]string {
	data := cr.store.Get()
	return map[string]string{
		"pathBookmarks": data.PathBookmarks,
		"pathCommands":  data.PathCommands,
		"browserApp":    data.BrowserApp,
	}
}

func (cr *ConfigRepository) Get(key string) (string, error) {
	data := cr.store.Get()
	switch key {
	case "pathBookmarks":
		return data.PathBookmarks, nil
	case "pathCommands":
		return data.PathCommands, nil
	case "browserApp":
		return data.BrowserApp, nil
	default:
		return "", fmt.Errorf("no such field: %s", key)
	}
}

func (cr *ConfigRepository) Set(key, value string) error {
	data := cr.store.Get()
	switch key {
	case "pathBookmarks":
		data.PathBookmarks = value
	case "pathCommands":
		data.PathCommands = value
	case "browserApp":
		data.BrowserApp = value
	default:
		return fmt.Errorf("no such field: %s", key)
	}

	err := cr.store.Set(data)
	if err != nil {
		return fmt.Errorf("save config: %w", err)
	}

	return nil
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
