package scheduler

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/kavehmz/bullbear/exchange"
)

type exchangeMock struct {
	exit int64
}

func (s *exchangeMock) Pull(sym exchange.SymbolCode) error {
	exit := atomic.LoadInt64(&s.exit)

	if exit == 1 {
		return errors.New("test_error")
	}
	return nil
}
func Test_schedule(t *testing.T) {
	errors := make(chan error)
	ran := make(chan bool)
	j := (&Job{}).Define(&exchangeMock{exit: 1}, exchange.SymbolCode{}, time.Nanosecond).Errors(errors).Ran(ran)
	(&Scheduler{}).Add(j).Run()

	select {
	case e := <-errors:
		if e.Error() != "test_error" {
			t.Error("expecting error", e)
		}
	case <-time.After(time.Millisecond):
		t.Error("expecting error but timed out")
	}

	j = (&Job{}).Define(&exchangeMock{}, exchange.SymbolCode{}, time.Nanosecond).Errors(errors).Ran(ran)
	(&Scheduler{}).Add(j).Run()

	select {
	case <-ran:
	case <-time.After(time.Millisecond):
		t.Error("expecting ran but timed out")
	}

	j = (&Job{}).Define(&exchangeMock{}, exchange.SymbolCode{}, time.Minute).Errors(errors).Ran(ran)
	(&Scheduler{}).Add(j).Run()

	go func() { j.Cancel() }()

	select {
	case <-ran:
		t.Error("Did expected to get cancelled")
	case <-time.After(time.Millisecond):
	}
}
