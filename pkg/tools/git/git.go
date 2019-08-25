package git

import (
	"fmt"

	"github.com/dimastark/configs/pkg/tool"
)

// Git works with `git` files
type Git struct {
	*tool.Tool
}

// New initialises Git struct
func New() *Git {
	return &Git{tool.New("git")}
}

// GitConfig returns path to .gitconfig
func (g *Git) GitConfig() string {
	return g.GetPathInToolDir(".gitconfig")
}

// GitIgnore returns path to .gitignore
func (g *Git) GitIgnore() string {
	return g.GetPathInToolDir(".gitignore")
}

// Backup implements Backuper interface
func (g *Git) Backup() error {
	return g.Exec(fmt.Sprintf(
		`cp ~/.gitconfig "%s" && cp ~/.gitignore "%s"`,
		g.GitConfig(),
		g.GitIgnore(),
	))
}

// Restore implements Restorer interface
func (g *Git) Restore() error {
	return g.Exec(fmt.Sprintf(
		`cp "%s" ~/.gitconfig && cp "%s" ~/.gitignore`,
		g.GitConfig(),
		g.GitIgnore(),
	))
}
