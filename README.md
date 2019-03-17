# bullbear

[![Go Lang](http://kavehmz.github.io/static/gopher/gopher-front.svg)](https://golang.org/)
[![Build Status](https://travis-ci.org/kavehmz/bullbear.svg?branch=master)](https://travis-ci.org/kavehmz/bullbear)
[![Coverage Status](https://coveralls.io/repos/github/kavehmz/bullbear/badge.svg?branch=master)](https://coveralls.io/github/kavehmz/bullbear?branch=master)

This is an app for fetching and retrieving market data

# How to run

You dont need anything beside docker to test/lint or run the app. Everything happens in docker.

```bash
make up
```

This will create a network named `influxdb` and then starts `influxdb` itself and also `chronograf` and at last it will start our app.

You can use chronograf at http://localhost:8888/ or `make example-select` to query the influxdb and see what is saved.
