package fish

import (
	"fmt"

	"github.com/dimastark/configs/pkg/tool"
)

// Fish works with `fish` files
type Fish struct {
	*tool.Tool
}

// New initialises Fish struct
func New() *Fish {
	return &Fish{tool.New("fish")}
}

// Config returns path to fish config
func (f *Fish) Config() string {
	return f.GetPathInToolDir("config.fish")
}

// Functions returns path to fish functions
func (f *Fish) Functions() string {
	return f.GetPathInToolDir("functions")
}

// Installed implements Installed interface
func (f *Fish) Installed() bool {
	return f.ExecSilent("fish -v") == nil
}

// Install implements Installed interface
func (f *Fish) Install() error {
	return f.Exec("brew install fish")
}

// Backup implements Backuper interface
func (f *Fish) Backup() error {
	return f.Exec(fmt.Sprintf(
		`cp ~/.config/fish/config.fish "%s" `+
			`&& cp -r ~/.config/fish/functions "%s"`,
		f.Config(),
		f.Functions(),
	))
}

// Restore implements Restorer interface
func (f *Fish) Restore() error {
	return f.Exec(fmt.Sprintf(
		`cp "%s" ~/.config/fish/config.fish `+
			`&& cp -r "%s" ~/.config/fish/functions`,
		f.Config(),
		f.Functions(),
	))
}

// Cleanup implements Cleanuper interface
func (f *Fish) Cleanup() error {
	return nil
}
