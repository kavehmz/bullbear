package exchange

import (
	"errors"
	"testing"
)

type source struct{ err bool }

func (s *source) Pull(symbol SymbolCode) ([]*Tick, error) {
	if s.err {
		return nil, errors.New("test_source_error")
	}
	return []*Tick{&Tick{Symbol: symbol, Value: 4010500000000}}, nil
}

type store struct {
	err  bool
	data map[SymbolCode]int64
}

func (s *store) Insert(ticks []*Tick) error {
	if s.err {
		return errors.New("test_store_error")
	}
	s.data[ticks[0].Symbol] = ticks[0].Value
	return nil
}

func TestExchange_Pull(t *testing.T) {
	so := &source{}
	st := &store{data: make(map[SymbolCode]int64)}
	sym := SymbolCode{Base: "BTC", Target: "USD"}

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
