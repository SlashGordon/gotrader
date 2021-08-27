package main

import (
	"github.com/SlashGordon/gotrader/cmd"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("$GOTRADER$", "", "green", true)
	myFigure.Print()
	cmd.Execute()
}
