package monita

import (
	"github.com/dimastark/configs/pkg/tool"
)

// Monita installs monita CLI
type Monita struct {
	*tool.Tool
}

// New initialises Monita
func New() *Monita {
	return &Monita{tool.New("monita")}
}

// Installed implements Installed interface
func (m *Monita) Installed() bool {
	return m.ExecSilent("monita --help") == nil
}

// Install implements Installed interface
func (m *Monita) Install() error {
	return m.Exec(`curl -sL https://github.com/dimastark/monita/releases/download/0.0.4/darwin_monita -o /usr/local/bin/monita && chmod +x /usr/local/bin/monita`)
}
