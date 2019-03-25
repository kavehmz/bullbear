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
	j := (&Job{}).Define(&exchangeMock{exit: 1}, exchange.SymbolCode{}, time.Nanosecond).Errors(errors)
	(&Scheduler{}).Add(j).Run()

	testSleepTime := time.Millisecond * 100
	select {
	case e := <-errors:
		if e.Error() != "test_error" {
			t.Error("expecting error", e)
		}
	case <-time.After(testSleepTime):
		t.Error("expecting error but timed out")
	}

	ran := make(chan bool)
	j = (&Job{}).Define(&exchangeMock{}, exchange.SymbolCode{}, time.Nanosecond).Ran(ran)
	(&Scheduler{}).Add(j).Run()

	select {
	case <-ran:
	case <-time.After(testSleepTime):
		t.Error("expecting ran but timed out")
	}

	ran = make(chan bool)
	j = (&Job{}).Define(&exchangeMock{}, exchange.SymbolCode{}, time.Minute).Ran(ran)
	j.Cancel()
	(&Scheduler{}).Add(j).Run()

	select {
	case <-ran:
		t.Error("Did expected to get cancelled")
	case <-time.After(testSleepTime):
	}
}
