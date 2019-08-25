package python

import (
	"fmt"

	"github.com/dimastark/configs/pkg/tool"
)

// Python works with `python` files
type Python struct {
	*tool.Tool
}

// New initialises Python struct
func New() *Python {
	return &Python{tool.New("python")}
}

// RequirementsTxt returns path to deps list
func (p *Python) RequirementsTxt() string {
	return p.GetPathInToolDir("requirements.txt")
}

// Installed implements Installed interface
func (p *Python) Installed() bool {
	return p.ExecSilent("python3 --version") == nil
}

// Install implements Installed interface
func (p *Python) Install() error {
	return p.Exec("brew install python")
}

// Backup implements Backuper interface
func (p *Python) Backup() error {
	return p.Exec(fmt.Sprintf(
		`pip3 freeze > "%s"`,
		p.RequirementsTxt(),
	))
}

// Restore implements Restorer interface
func (p *Python) Restore() error {
	err := p.Exec(fmt.Sprintf(
		`pip3 install -r "%s"`,
		p.RequirementsTxt(),
	))

	if err != nil {
		return err
	}

	return p.Cleanup()
}

// Cleanup implements Cleanuper interface
func (p *Python) Cleanup() error {
	err := p.Exec(fmt.Sprintf(
		`pip3 freeze `+
			`| grep -v -f "%s" - `+
			`| xargs pip3 uninstall -y`,
		p.RequirementsTxt(),
	))

	if err != nil {
		return err
	}

	return p.Exec(
		`pip3 freeze ` +
			`| grep -v "^\\-e" ` +
			`| cut -d = -f 1 ` +
			`| xargs -n1 pip3 install -U`,
	)
}
