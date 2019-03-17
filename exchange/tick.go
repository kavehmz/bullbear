package exchange

import "time"

// SymbolCode identifies a symbol
type SymbolCode struct {
	Base, Target string
}

// Tick represents market data
type Tick struct {
	Timestamp time.Time
	// Value with nano precision
	Value int64
	// e.x BTCUSD
	Symbol SymbolCode
}
