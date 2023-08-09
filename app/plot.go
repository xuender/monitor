package app

import (
	"sort"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/kfont/tsanger/ym"
	"github.com/xuender/kit/logs"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg/draw"
)

type Plot struct {
	size int
}

func NewPlot() *Plot {
	font.DefaultCache.Add([]font.Face{
		{
			Font: font.Font{Typeface: "ym"},
			Face: ym.Font(),
		},
	})

	plot.DefaultFont = font.Font{
		Typeface: "ym",
	}

	return &Plot{size: _size}
}

func (p *Plot) CPU(canvas draw.Canvas, procs map[string][]time.Duration) {
	cpu := plot.New()
	cpu.Title.Text = "CPU 消耗图"

	color := 0

	titles := make([]string, 0, len(procs))
	for k := range procs {
		titles = append(titles, k)
	}

	sort.Strings(titles)
	lo.Reverse(titles)

	for _, title := range titles {
		xys := make([]plotter.XY, p.size)
		for i := 0; i < p.size; i++ {
			xys[i] = plotter.XY{X: float64(i), Y: 0}
		}

		var old time.Duration

		durs := procs[title]
		if len(durs) > p.size+1 {
			durs = durs[len(durs)-p.size-1:]
		}

		for index, dur := range durs {
			if index > 0 {
				xys[p.size-len(durs)+index].Y = float64((dur - old) / 1_000_000) // nolint
			}

			old = dur
		}

		line := lo.Must1(plotter.NewLine(plotter.XYs(xys)))
		line.Color = plotutil.Color(color)
		cpu.Legend.Add(title, line)
		cpu.Add(line)
		logs.D.Println("line:", color, title, len(durs))
		color++
	}

	cpu.Draw(canvas)
}
