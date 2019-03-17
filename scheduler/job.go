package scheduler

import (
	"context"
	"time"

	"github.com/kavehmz/bullbear/exchange"
)

// Exchange define the interface for what this scheduler accepts for pulling data
type Exchange interface {
	Pull(symbol exchange.SymbolCode) error
}

// Job identifies a preiodical runs and its specs
type Job struct {
	exch   Exchange
	freq   time.Duration
	symb   exchange.SymbolCode
	errors chan error
	ran    chan bool
	ctx    context.Context
	cancel context.CancelFunc
}

// Define sets the main params of the job
func (s *Job) Define(exch Exchange, symbol exchange.SymbolCode, freq time.Duration) *Job {
	s.exch = exch
	s.symb = symbol
	s.freq = freq
	s.ctx, s.cancel = context.WithCancel(context.Background())

	return s
}

// Ran will set the channel for receiving info whenever there is a new tick. It is optional
func (s *Job) Ran(ran chan bool) *Job {
	s.ran = ran
	return s
}

// Errors will set the channel for receiving errors. It is optional
func (s *Job) Errors(errors chan error) *Job {
	s.errors = errors
	return s
}

// Cancel stops the job
func (s *Job) Cancel() {
	s.cancel()
}
