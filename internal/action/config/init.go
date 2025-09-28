package config

import "github.com/viktor-titov/bookmarks/constants"

type ConfigRepositoryForInit interface {
	Set(key, value string) error
}

type Init struct {
	repo ConfigRepositoryForInit
}

func NewInit(repo ConfigRepositoryForInit) *Init {
	return &Init{
		repo: repo,
	}
}

func (i *Init) Do() error {
	if err := i.repo.Set("pathBookmarks", constants.DefaultPathBookmarks); err != nil {
		return err
	}
	if err := i.repo.Set("pathCommands", constants.DefaultPathCommands); err != nil {
		return err
	}
	if err := i.repo.Set("browserApp", constants.DefaultBrowserApp); err != nil {
		return err
	}
	return nil
}
