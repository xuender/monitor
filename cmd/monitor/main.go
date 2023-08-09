package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/xuender/monitor/app"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	lo.Must0(ebiten.RunGame(app.InitApp()))
}

func usage() {
	fmt.Fprintf(os.Stderr, "monitor\n\n")
	fmt.Fprintf(os.Stderr, "TODO: monitor.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
