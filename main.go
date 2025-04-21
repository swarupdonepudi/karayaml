package main

import (
	planton "github.com/swarupdonepudi/karayaml/cmd/karayaml"
	clipanic "github.com/swarupdonepudi/karayaml/internal/panic"
)

func main() {
	finished := new(bool)
	defer clipanic.Handle(finished)
	planton.Execute()
	*finished = true
}
