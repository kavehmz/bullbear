package exchange

import "time"

// Tick represents market data
type Tick struct {
	Timestamp time.Time
	// Value with nano precision
	Value int64
	// e.x BTCUSD
	Symbol string
}
