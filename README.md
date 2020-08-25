# branchless



```shell script
go test -bench=. -benchmem

goos: darwin
goarch: amd64
pkg: github.com/romanitalian/branchless
BenchmarkMinBranchless-4        1000000000               0.585 ns/op           0 B/op          0 allocs/op
BenchmarkMin-4                  1000000000               0.596 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/romanitalian/branchless      1.587s
```


```shell script
go test -bench=. -benchmem

goos: darwin
goarch: amd64
pkg: github.com/romanitalian/branchless
BenchmarkMinBranchless-4        1000000000               0.590 ns/op           0 B/op          0 allocs/op
BenchmarkMin-4                  1000000000               0.571 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/romanitalian/branchless      1.366s
```
