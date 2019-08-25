package node

import (
	"fmt"

	"github.com/dimastark/configs/pkg/tool"
)

// Node works with `node` files
type Node struct {
	*tool.Tool
}

// New initialises Node struct
func New() *Node {
	return &Node{tool.New("node")}
}

// Bundle returns path to deps list
func (n *Node) Bundle() string {
	return n.GetPathInToolDir("npm-ls")
}

// Installed implements Installed interface
func (n *Node) Installed() bool {
	return n.ExecSilent("node -v") == nil
}

// Install implements Installed interface
func (n *Node) Install() error {
	return n.Exec("brew install node")
}

// Backup implements Backuper interface
func (n *Node) Backup() error {
	return n.Exec(fmt.Sprintf(
		`npm ls -gp --depth=0 `+
			`| awk '{gsub(/\/.*\//,"",$1); print}' `+
			`| grep -v lib | grep -v npm | cat > "%s"`,
		n.Bundle(),
	))
}

// Restore implements Restorer interface
func (n *Node) Restore() error {
	return n.Exec(fmt.Sprintf(
		`cat "%s" | xargs npm install -g`,
		n.Bundle(),
	))
}

// Cleanup implements Cleanuper interface
func (n *Node) Cleanup() error {
	return n.Exec(
		`npm ls -gp --depth=0 ` +
			`| awk '{gsub(/\/.*\//,"",$1); print}' ` +
			`| grep -v lib | grep -v npm ` +
			`| xargs npm -g rm`,
	)
}
