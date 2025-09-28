package command

import (
	"fmt"
	"os"
	"os/exec"
)

type Run struct {
	repo CommandRepository
}

func NewRun(repoCommand CommandRepository) *Run {
	return &Run{repo: repoCommand}
}

func (g *Run) Do(key string) error {
	command, err := g.repo.Get(key)
	if err != nil {
		return fmt.Errorf("get command: %w", err)
	}

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run command: %w", err)
	}
	return nil
}
