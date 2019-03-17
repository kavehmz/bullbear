package scheduler

import (
	"errors"
	"time"

	"github.com/kavehmz/bullbear/exchange"
)

// ErrInitPulls defines the init error
var ErrInitPulls = errors.New("Init pulls")

// Scheduler will schedule runs for each symbol in its exchange
type Scheduler struct {
	Jobs map[exchange.SymbolCode]*Job
}

// Run starts the pulling process
func (s *Scheduler) Run() *Scheduler {
	for _, job := range s.Jobs {
		go schedule(job)
	}

	return s
}

// Add adds a new exchange to the list
func (s *Scheduler) Add(job *Job) *Scheduler {
	if s.Jobs == nil {
		s.Jobs = make(map[exchange.SymbolCode]*Job)
	}
	s.Jobs[job.symb] = job
	return s
}

func schedule(job *Job) {
	for {
		select {
		case <-job.ctx.Done():
			return
		case <-time.After(job.freq):
			err := job.exch.Pull(job.symb)
			if err != nil && job.errors != nil {
				job.errors <- err
				break
			}
			if job.ran != nil {
				job.ran <- true
			}
		}
	}
}
