/*Gather fetch market data
Source of this application is Coindesk.

	https://www.coindesk.com/api
	CoinDesk provides a simple API to make its Bitcoin Price Index (BPI) data programmatically available to others. You are free to use this API to include our data in any application or website as you see fit, as long as each page or app that uses it includes the text “Powered by CoinDesk”, linking to our price page.

This app is just an exercise for playing with influxdb and has no other purposes at this moment.
*/
package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kavehmz/bullbear/exchange"

	"github.com/kavehmz/bullbear/scheduler"

	"github.com/kavehmz/bullbear/httpclient"
	"github.com/kavehmz/bullbear/store/influx"

	"github.com/kavehmz/bullbear/source/coindesk"

	client "github.com/influxdata/influxdb1-client"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	freq         = kingpin.Flag("frequency", "how often fetch new rates in ms").Envar("FREQ").Default("10000").Int64()
	influxdbHost = kingpin.Flag("influxdb", "influxdb host").Envar("INFLUXDB").Default("http://localhost:8086").String()
)

func main() {
	// This notice is required by CoinDesk. They mentioned use of their api is free as long as this notice is displayed.
	// Otherwise this app is just an exercise for testing use influxdb in Go.
	log.Println(`“Powered by [CoinDesk](https://www.coindesk.com/price/bitcoin)”`)
	kingpin.Parse()

	host, err := url.Parse(*influxdbHost)
	if err != nil {
		log.Panic(err)
	}
	fluxClient, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		log.Panic(err)
	}
	flux := &influx.Influx{Client: fluxClient, Database: "market"}

	exch := &exchange.Exchange{
		Source: &coindesk.CoinDesk{HTTPClient: &httpclient.HTTPClient{}},
		Store:  flux,
	}

	job := (&scheduler.Job{}).Define(exch, exchange.SymbolCode{Base: "BTC", Target: "USD"}, time.Millisecond*time.Duration(*freq)).Errors(logErrors("CoinDesk")).Ran(countTicks("CoinDesk"))

	(&scheduler.Scheduler{}).Add(job).Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT)
	sig := <-sigChan
	log.Printf("Received signal '%v', shutting down\n", sig)
}

func logErrors(title string) chan error {
	ch := make(chan error)
	go func() {
		for e := range ch {
			log.Printf("%s: %v", title, e)
		}
	}()
	return ch
}

func countTicks(title string) chan bool {
	ch := make(chan bool)
	count := 0
	go func() {
		for range ch {
			count++
			log.Printf("%s: received %d", title, count)
		}
	}()
	return ch
}
