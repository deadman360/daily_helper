package main

import (
	"github.com/deadman360/daily_helper/internal/infrastructure"
	"github.com/deadman360/daily_helper/internal/interface/cli"
)

func main() {
	cli.Execute(infrastructure.NewRepository("repositorio"))
}
