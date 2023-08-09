package app

import (
	"context"
	"time"

	"github.com/elastic/go-sysinfo"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
)

const _ticker = time.Millisecond * 100

type Process struct {
	Procs map[string][]time.Duration
}

func NewProcess() *Process {
	return &Process{
		Procs: make(map[string][]time.Duration),
	}
}

func (p *Process) Run(ctx context.Context) {
	ticker := time.NewTicker(_ticker)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.tick()
		}
	}
}

func (p *Process) tick() {
	logs.D.Println("tick", len(p.Procs))

	pro := lo.Must1(sysinfo.Host())
	cpu := lo.Must1(pro.CPUTime())

	// p.add("total", cpu.Total())
	p.add("user", cpu.User)
	p.add("system", cpu.System)
	// p.add("idle", cpu.Idle)
	p.add("iowait", cpu.IOWait)
	p.add("irq", cpu.IRQ)
	// p.add("nice", cpu.Nice)
	p.add("soft irq", cpu.SoftIRQ)
	p.add("steal", cpu.Steal)
}

func (p *Process) add(label string, dur time.Duration) {
	if drus, has := p.Procs[label]; has {
		if len(drus) > _size+1 {
			drus = drus[1:]
		}

		p.Procs[label] = append(drus, dur)
	} else {
		p.Procs[label] = []time.Duration{dur}
	}
}
