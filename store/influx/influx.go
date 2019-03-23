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
func (s *Influx) Insert(ticks []*exchange.Tick) error {
	points := []client.Point{}
	for _, t := range ticks {
		point := client.Point{
			Measurement: "tick",
			Tags: map[string]string{
				"symbol": t.Symbol.Base + t.Symbol.Target,
				"base":   t.Symbol.Base,
				"target": t.Symbol.Target,
			},
			Fields: map[string]interface{}{
				"value": t.Value,
			},
			Time: t.Timestamp,

			Precision: "ns",
		}
		points = append(points, point)
	}

	bps := client.BatchPoints{
		Points:   points,
		Database: s.Database,
	}
	_, err := s.Client.Write(bps)
	return err
}
