package link

import (
	"fmt"
	"os/exec"
)

type Go struct {
	repoLink LinkRepository
	cfg      ConfigRepository
}

func NewGo(repoLink LinkRepository, cfg ConfigRepository) *Go {
	return &Go{repoLink: repoLink, cfg: cfg}
}

func (g *Go) Do(key string) error {
	app, err := g.cfg.Get("browserApp")
	if err != nil {
		return fmt.Errorf("get browser app from config: %w", err)
	}

	link, err := g.repoLink.Get(key)
	if err != nil {
		return fmt.Errorf("get link: %w", err)
	}

	exec.Command(app, link).Start()
	return nil
}
