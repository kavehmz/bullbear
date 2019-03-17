.PHONY: $(shell ls -d *)

default:
	@echo "Usage: make [command]"

build:
	docker build -t bullbear .

influxdb:
	mkdir -p $$PWD/data/influxdb
	docker run -d --rm --name=influxdb -p 8086:8086 \
		--net=influxdb \
		-v $$PWD/data/influxdb:/var/lib/influxdb \
		-e INFLUXDB_DB=market \
		influxdb -config /etc/influxdb/influxdb.conf

chronograf:
	mkdir -p $$PWD/data/chronograf
	docker run -d --rm --name=chronograf -p 8888:8888 \
		--net=influxdb \
		-v $$PWD/data/chronograf:/var/lib/chronograf \
		chronograf --influxdb-url=http://influxdb:8086

# shared network between the app, influxdb and chronograf
network:
	echo "creating a network for the app"
	(docker network ls|grep -q influxdb) || docker network create influxdb

run: build
	docker run --rm --name=bullbear-gather \
		--net=influxdb \
		bullbear --influxdb="http://influxdb:8086"

up: network influxdb chronograf run

down:
	docker stop chronograf influxdb || true

influx:
	docker exec -ti influxdb influx --database market -precision rfc3339

example-insert:
	docker exec -ti influxdb influx --database market -precision rfc3339  -execute 'INSERT tick,symbol=BTCUSD value=4010.05'

example-select:
	docker exec -ti influxdb influx --database market -precision rfc3339  -execute 'select * from tick' -format 'column'  -pretty

test:
	docker run --rm -u`id -u`:`id -g` -v $$PWD:/go/src/github.com/kavehmz/bullbear golang:1 /bin/bash -c \
	cd /go/src/github.com/kavehmz/bullbear;\
	go test -v --race -cover -coverprofile=cover.out ./...; \
	go tool cover -func=cover.out | \
		awk 'END {sub("[.].*","",$$NF); printf "Coverage: %d%%\n", $$NF; \
			if ($$NF+0 < 100) {print "Coverage is not sufficient"; exit 1}}'

lint:
	docker run --rm -v $$PWD:/go/src/github.com/kavehmz/bullbear -w /go/src/github.com/kavehmz/bullbear \
		golangci/golangci-lint:latest golangci-lint run ./...
