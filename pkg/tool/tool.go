package tool

import (
	"go/build"
	"os"
	"os/exec"
	"path"
)

const (
	packageName = "github.com/dimastark/configs"
)

// Namer interface with name
type Namer interface {
	Name() string
}

// Installer runs if there is need to install some deps
type Installer interface {
	Install() error
	Installed() bool
}

// Backuper runs on "backup" command
type Backuper interface {
	Backup() error
}

// Restorer runs on "restore" command
type Restorer interface {
	Restore() error
}

// Cleanuper runs on "cleanup" command
type Cleanuper interface {
	Cleanup() error
}

// Tool has utility methods for reading and writing tool configs
type Tool struct {
	name string
}

// New initialises Tool struct
func New(name string) *Tool {
	return &Tool{name}
}

// Name returns tool `name`
func (t *Tool) Name() string {
	return t.name
}

// GetPathInToolDir returns path relative to tool configs
func (t *Tool) GetPathInToolDir(filePath string) string {
	return path.Join(build.Default.GOPATH, "src", packageName, "pkg", "tools", t.name, filePath)
}

// GetBaseInToolDir returns base of filePath relative to tool configs
func (t *Tool) GetBaseInToolDir(filePath string) string {
	return t.GetPathInToolDir(path.Base(filePath))
}

// Exec runs `command` in /bin/sh
func (t *Tool) Exec(command string) error {
	return cmd(command, false)
}

// ExecSilent runs `command` in /bin/sh silently
func (t *Tool) ExecSilent(command string) error {
	return cmd(command, true)
}

func cmd(command string, silent bool) error {
	cmd := exec.Command("/bin/sh", "-c", command)

	if !silent {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}
