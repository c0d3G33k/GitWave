package main

import (
	"github.com/c0d3G33k/GitWave/cmd"
	"github.com/c0d3G33k/GitWave/pkg/common/ui"
)

var banner string = `
 ____    ___   _____  __        __     _     __     __  _____ 
/ ___|  |_ _| |_   _| \ \      / /    / \    \ \   / / | ____|
| |  _   | |    | |    \ \ /\ / /    / _ \    \ \ / /  |  _|  
| |_| |  | |    | |     \ V  V /    / ___ \    \ V /   | |___ 
\____|  |___|   |_|      \_/\_/    /_/   \_\    \_/    |_____|
`

func main() {
	ui.PrintMsg(banner)
	cmd.Execute()
}
