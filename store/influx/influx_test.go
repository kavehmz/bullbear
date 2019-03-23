package influx

import (
	"testing"
	"time"

	"github.com/kavehmz/bullbear/exchange"

	client "github.com/influxdata/influxdb1-client"
)

type influxMock struct {
	bp client.BatchPoints
}

func (s *influxMock) Write(bp client.BatchPoints) (*client.Response, error) {
	s.bp = bp
	return nil, nil
}

func TestInflux_Insert(t *testing.T) {
	imock := influxMock{}
	i := Influx{Client: &imock}
	err := i.Insert([]*exchange.Tick{&exchange.Tick{Value: 1000000000, Timestamp: time.Now(), Symbol: exchange.SymbolCode{Base: "BTC", Target: "USD"}}})
	if err != nil || imock.bp.Points == nil {
		t.Error("expected to save data")
	}
}
