package main

import (
	"context"
	"image/png"
	"os"

	"github.com/samber/lo"
	"github.com/xuender/kit/oss"
	"github.com/xuender/monitor/app"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

// nolint: gomnd
func main() {
	proc := app.NewProcess()
	plot := app.NewPlot()
	ctx, cancel := context.WithCancel(context.Background())

	go oss.Cancel(cancel)

	proc.Run(ctx)

	// canvas := vgimg.NewWith(vgimg.UseImage(img))
	canvas := vgimg.New(640, 480)
	canvas.Image()
	// canvas := vgsvg.New(3*vg.Inch, 3*vg.Inch)
	plot.CPU(draw.New(canvas), proc.Procs)

	// file := lo.Must1(os.Create("cpu.svg"))
	// defer file.Close()

	// canvas.WriteTo(file)
	file := lo.Must1(os.Create("cpu.png"))
	defer file.Close()

	lo.Must0(png.Encode(file, canvas.Image()))
}
