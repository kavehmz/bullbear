// Package exchange handles the operation related to pulling and inserting data.
// It can use different sources and storages as long as they implement the correct interface.
package exchange

// Exchange represents any source exchange to retrieve data from.
type Exchange struct {
	Source interface {
		Pull(symbol SymbolCode) (*Tick, error)
	}
	Store interface {
		Insert(tick *Tick) error
	}
}

// Pull get market data for a symbol and saves it in the db.
func (s *Exchange) Pull(symbol SymbolCode) error {
	tick, err := s.Source.Pull(symbol)
	if err != nil {
		return err
	}

	return s.Store.Insert(tick)
}
