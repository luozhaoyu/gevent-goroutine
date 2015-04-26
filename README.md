# gevent-goroutine
Comparison of Two Lightweight Processing Techniques: gevent vs goroutine

## install
### install apachebench
* follow [apachebench-standalone] (https://code.google.com/p/apachebench-standalone/wiki/HowToBuild)
* install [apr and apr-util] (http://archive.apache.org/dist/apr/)
* compile ab

## run
### go
1. `go run go_server.go`
- `w3m http://localhost:8080/contention?lock=0&print=1`

### gevent

## test
### ab
* normal echo: `ab -n 1000 -c 10 'http://localhost:8080/'`
* normal echo with command line output (default is disabled): `ab -n 1000 -c 10 'http://localhost:8080/?print=1'`
* contention with lock (default is enabled): `ab -n 1000 -c 10 'http://localhost:8080/contention'`
* contention without lock (default is enabled): `ab -n 1000 -c 10 'http://localhost:8080/contention?lock=0'`

## profile
### go
* visit [http://localhost:6060/debug/pprof/]

## reference
* [profiling go programs] (http://blog.golang.org/profiling-go-programs)
