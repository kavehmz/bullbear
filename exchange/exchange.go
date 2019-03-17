package exchange

// Exchange represents any source exchange to retrieve data from.
type Exchange struct {
	Source interface {
		Pull(symbol string) (*Tick, error)
	}
	Store interface {
		Insert(tick *Tick) error
	}
}

// Pull get market data for a symbol and saves it in the db.
func (s *Exchange) Pull(symbol string) error {
	tick, err := s.Source.Pull(symbol)
	if err != nil {
		return err
	}

	return s.Store.Insert(tick)
}
