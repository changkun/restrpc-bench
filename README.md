# restrpc-bench

```go
→go test -bench=. -benchmem
INFO[0000] http is running                              
INFO[0000] grpc is running                              
goos: darwin
goarch: amd64
pkg: github.com/changkun/restrpc-bench
BenchmarkAPIRestful-4               2000            539627 ns/op           31610 B/op        206 allocs/op
BenchmarkAPIgRPC-4                 10000            198214 ns/op           10121 B/op        182 allocs/op
PASS
ok      github.com/changkun/restrpc-bench       3.232s
```