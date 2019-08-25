package tools

import (
	"github.com/dimastark/configs/pkg/tool"
	"github.com/dimastark/configs/pkg/tools/brew"
	"github.com/dimastark/configs/pkg/tools/fish"
	"github.com/dimastark/configs/pkg/tools/git"
	"github.com/dimastark/configs/pkg/tools/node"
	"github.com/dimastark/configs/pkg/tools/python"
)

// Registry contains all tools
var Registry = []tool.Namer{
	brew.New(),
	fish.New(),
	git.New(),
	node.New(),
	python.New(),
}
