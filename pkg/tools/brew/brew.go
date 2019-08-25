package brew

import (
	"fmt"

	"github.com/dimastark/configs/pkg/tool"
)

// Brew works with `brew` files
type Brew struct {
	*tool.Tool
}

// New initialises Brew struct
func New() *Brew {
	return &Brew{tool.New("brew")}
}

// Brewfile returns path to Brewfile
func (b *Brew) Brewfile() string {
	return b.GetPathInToolDir("Brewfile")
}

// Installed implements Installed interface
func (b *Brew) Installed() bool {
	return b.ExecSilent("brew --version") == nil
}

// Install implements Installed interface
func (b *Brew) Install() error {
	return b.Exec(`/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`)
}

// Backup implements Backuper interface
func (b *Brew) Backup() error {
	return b.Exec(fmt.Sprintf(
		`brew bundle dump --file="%s" --force`,
		b.Brewfile(),
	))
}

// Restore implements Restorer interface
func (b *Brew) Restore() error {
	return b.Exec(fmt.Sprintf(
		`brew bundle --file="%s"`,
		b.Brewfile(),
	))
}

// Cleanup implements Cleanuper interface
func (b *Brew) Cleanup() error {
	return b.Exec(
		"brew update " +
			"&& brew upgrade " +
			"&& brew cask upgrade " +
			"&& brew cleanup -s",
	)
}
