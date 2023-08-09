package app

import (
	"context"
	"image"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xuender/kit/cache"
	"github.com/xuender/kit/oss"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

const _key = "image"

type App struct {
	cancel  context.CancelFunc
	cache   *cache.Cache[string, *ebiten.Image]
	plot    *Plot
	process *Process
}

func NewApp(
	process *Process,
	plot *Plot,
) *App {
	ebiten.SetWindowTitle("Monitor")
	ebiten.SetWindowClosingHandled(true)

	ctx, cancel := oss.CancelContext(context.Background())
	go process.Run(ctx)
	// nolint: gomnd
	return &App{
		cancel:  cancel,
		cache:   cache.NewStringKey[*ebiten.Image](time.Millisecond*100, time.Second),
		plot:    plot,
		process: process,
	}
}

func (p *App) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "test")
	eimg, has := p.cache.GetNoExtension(_key)
	if !has {
		procs := p.process.Procs
		if len(procs) == 0 {
			return
		}

		img := image.NewRGBA(image.Rect(0, 0, _width, _height))
		canvas := vgimg.NewWith(vgimg.UseImage(img))
		p.plot.CPU(draw.New(canvas), procs)

		eimg = ebiten.NewImageFromImage(canvas.Image())
		p.cache.Set(_key, eimg)
	}

	screen.DrawImage(eimg, &ebiten.DrawImageOptions{})
}

func (p *App) Update() error {
	if ebiten.IsWindowBeingClosed() {
		p.cancel()
		ebiten.SetWindowClosingHandled(false)
		os.Exit(0)
	}

	return nil
}

func (p *App) Layout(_, _ int) (int, int) { return _width, _height }
