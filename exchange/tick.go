package exchange

import "time"

// SymbolCode identifies a symbol
type SymbolCode struct {
	Base, Target string
}

// Tick represents market data
type Tick struct {
	Timestamp time.Time
	// e.x BTCUSD
	Symbol SymbolCode
	// Value with nano precision
	Value int64
	// All with nano precision. They dont exists for all sources
	Open, High, Low, Close *int64
	Volume                 *int64
}
