// Pakckage links содержит реализацию репозитория для работы с ссылками.
package links

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/viktor-titov/bookmarks/internal/adapter/config"
	"github.com/viktor-titov/bookmarks/internal/adapter/jsonstore"
	"github.com/viktor-titov/bookmarks/internal/models"
)

type LinkRepositoryInterface interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
	List() (map[string]string, error)
}

type LinksRepository struct {
	store *jsonstore.Store[models.Links]
}

var _ LinkRepositoryInterface = (*LinksRepository)(nil)

func NewLinksRepository(config config.ConfigRepository) (*LinksRepository, error) {
	path, err := config.Get("pathBookmarks")
	if err != nil {
		return nil, fmt.Errorf("get path from config: %w", err)
	}

	store, err := jsonstore.NewStore[models.Links](expandPath(path))
	if err != nil {
		return nil, fmt.Errorf("create links store: %w", err)
	}

	return &LinksRepository{
		store: store,
	}, nil
}

func (lr *LinksRepository) Get(key string) (string, error) {
	data := lr.store.Get()
	value, ok := data.Links[key]
	if !ok {
		return "", fmt.Errorf("no such link: %s", key)
	}
	return value, nil
}

func (lr *LinksRepository) Set(key, value string) error {
	data := lr.store.Get()
	if data.Links == nil {
		data.Links = make(map[string]string)
	}
	data.Links[key] = value
	lr.store.Set(data)

	return nil
}

func (lr *LinksRepository) Delete(key string) error {
	data := lr.store.Get()
	if data.Links == nil {
		return fmt.Errorf("no links to delete from")
	}
	if _, ok := data.Links[key]; !ok {
		return fmt.Errorf("no such link: %s", key)
	}
	delete(data.Links, key)
	lr.store.Set(data)

	return nil
}

func (lr *LinksRepository) List() (map[string]string, error) {
	data := lr.store.Get()
	if data.Links == nil {
		return make(map[string]string), nil
	}
	return data.Links, nil
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
