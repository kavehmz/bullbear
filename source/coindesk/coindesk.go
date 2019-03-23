/*Package coindesk retires data from CoinDesk API.
It is free as long as app mentions: Powered by [CoinDesk](https://www.coindesk.com/price/bitcoin)

	https://www.coindesk.com/api
	CoinDesk provides a simple API to make its Bitcoin Price Index (BPI) data programmatically available to others. You are free to use this API to include our data in any application or website as you see fit, as long as each page or app that uses it includes the text “Powered by CoinDesk”, linking to our price page.


coindesk package at the moment only supports BTCUSD and no other rate.
*/
package coindesk

import (
	"errors"

	"github.com/kavehmz/bullbear/exchange"
)

// ErrUnsupportedSymbol defines the Unsupported symbol error
var ErrUnsupportedSymbol = errors.New("Unsupported symbol")

// CoinDesk implement a source for CoinDesk API
type CoinDesk struct {
	HTTPClient interface {
		JSON(url string, v interface{}) error
	}
}

// Pull retrieves a tick
// https://api.coindesk.com/v1/bpi/currentprice.json
func (s *CoinDesk) Pull(symbol exchange.SymbolCode) ([]*exchange.Tick, error) {
	if symbol.Base != "BTC" || symbol.Target != "USD" {
		return nil, ErrUnsupportedSymbol
	}
	resp := response{}
	err := s.HTTPClient.JSON("https://api.coindesk.com/v1/bpi/currentprice.json", &resp)
	if err != nil {
		return nil, err
	}

	return []*exchange.Tick{&exchange.Tick{
		Symbol:    symbol,
		Value:     int64(resp.Bpi.USD.Ratefloat * 1000000000),
		Timestamp: resp.Time.UpdatedISO,
	}}, nil
}
