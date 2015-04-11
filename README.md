# gevent-goroutine
Comparison of Two Lightweight Processing Techniques: gevent vs goroutine

## install
### install apachebench
* follow [apachebench-standalone] (https://code.google.com/p/apachebench-standalone/wiki/HowToBuild)
* install [apr and apr-util] (http://archive.apache.org/dist/apr/)
* compile ab

## run
### go
1. `go run go_echo.go`
- `w3m http://localhost:8080/customized_path?with=get_string&digt=1`

### gevent

## test
### ab
`ab -n 1000 -c 10 http://localhost:8080/abtest`

## reference
* [profiling go programs] (http://blog.golang.org/profiling-go-programs)
