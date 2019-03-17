package influx

import (
	client "github.com/influxdata/influxdb1-client"
	"github.com/kavehmz/bullbear/exchange"
)

// Influx represents an influxdb store
type Influx struct {
	Client interface {
		Write(bp client.BatchPoints) (*client.Response, error)
	}
	Database string
}

// Insert save the tick
func (s *Influx) Insert(tick *exchange.Tick) error {
	point := client.Point{
		Measurement: "tick",
		Tags: map[string]string{
			"symbol": tick.Symbol.Base + tick.Symbol.Target,
			"base":   tick.Symbol.Base,
			"target": tick.Symbol.Target,
		},
		Fields: map[string]interface{}{
			"value": tick.Value,
		},
		Time:      tick.Timestamp,
		Precision: "ms",
	}
	bps := client.BatchPoints{
		Points:   []client.Point{point},
		Database: s.Database,
	}
	_, err := s.Client.Write(bps)
	return err
}
