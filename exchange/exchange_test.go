package exchange

import (
	"errors"
	"testing"
)

type source struct{ err bool }

func (s *source) Pull(symbol string) (*Tick, error) {
	if s.err {
		return nil, errors.New("test_source_error")
	}
	return &Tick{Symbol: symbol, Value: 4010500000000}, nil
}

type store struct {
	err  bool
	data map[string]int64
}

func (s *store) Insert(tick *Tick) error {
	if s.err {
		return errors.New("test_store_error")
	}
	s.data[tick.Symbol] = tick.Value
	return nil
}

func TestExchange_Pull(t *testing.T) {
	so := &source{}
	st := &store{data: make(map[string]int64)}
	sym := "BTCUSD"

	e := Exchange{Source: so, Store: st}

	err := e.Pull(sym)
	if err != nil || st.data[sym] != 4010500000000 {
		t.Error("Wrong sotrage", err, st)
	}

	// with source error
	so.err = true
	err = e.Pull(sym)
	if err == nil || err.Error() != "test_source_error" {
		t.Error("expecting source error", err, so)
	}
	so.err = false

	// with store error
	st.err = true
	err = e.Pull(sym)
	if err == nil || err.Error() != "test_store_error" {
		t.Error("expecting store error", err, st)
	}
}
