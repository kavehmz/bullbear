/*Package coindesk retires data from CoinDesk API.
It is free as long as app mentions: Powered by [CoinDesk](https://www.coindesk.com/price/bitcoin)

	https://www.coindesk.com/api
	CoinDesk provides a simple API to make its Bitcoin Price Index (BPI) data programmatically available to others. You are free to use this API to include our data in any application or website as you see fit, as long as each page or app that uses it includes the text “Powered by CoinDesk”, linking to our price page.


coindesk package at the moment only supports BTCUSD and no other rate.
*/
package coindesk

import (
	"errors"
	"testing"
	"time"

	"github.com/kavehmz/bullbear/exchange"
)

type client struct{ err bool }

func (s *client) JSON(url string, v interface{}) error {
	if s.err {
		return errors.New("test_error")
	}
	if val, ok := v.(*response); ok {
		val.Bpi.USD.Ratefloat = 100.1
		val.Time.UpdatedISO = time.Now()
	}
	return nil
}

func TestCoinDesk_Pull(t *testing.T) {
	c := &client{}
	cd := CoinDesk{HTTPClient: c}

	_, err := cd.Pull(exchange.SymbolCode{Base: "invalid"})
	if err != ErrUnsupportedSymbol {
		t.Error("expected error", err)
	}

	_, err = cd.Pull(exchange.SymbolCode{Base: "BTC", Target: "invalid"})
	if err != ErrUnsupportedSymbol {
		t.Error("expected error", err)
	}

	ticks, err := cd.Pull(exchange.SymbolCode{Base: "BTC", Target: "USD"})
	if err != nil || len(ticks) == 0 || ticks[0].Value != 100100000000 || time.Since(ticks[0].Timestamp) > time.Second {
		t.Error("expected error", err, ticks)
	}

	c.err = true
	_, err = cd.Pull(exchange.SymbolCode{Base: "BTC", Target: "USD"})
	if err.Error() != "test_error" {
		t.Error("expected error", err)
	}
}
