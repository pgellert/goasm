# Exploring SIMD Assembly in Golang

This is a small example for how to use SIMD and other assembly code in Golang
projects with the help of the [avo](github.com/mmcloughlin/avo) package.

I've implemented a SIMD-based uint64 slice addition method using AVX256 SIMD 
instructions that are available on my laptop. Then, I tested the implementation 
agains a simple pure-Go reference implementation using fuzzing. Finally, I 
benchmarked the SIMD implementation against the pure-Go one.

The results for adding random 10/100/1000/10000-length arrays show over 2x 
improvement for 100+ element arrays:

```
% go test -bench=. -benchtime=20s
goos: darwin
goarch: amd64
pkg: goasm
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkAdd10-16                       1000000000               4.911 ns/op
BenchmarkAdd100-16                      1000000000              20.03 ns/op
BenchmarkAdd1000-16                     127176942              194.6 ns/op
BenchmarkAdd10000-16                    10326230              2231 ns/op
BenchmarkBaselineAdd10-16               1000000000               5.143 ns/op
BenchmarkBaselineAdd100-16              445382661               55.27 ns/op
BenchmarkBaselineAdd1000-16             47817183               520.1 ns/op
BenchmarkBaselineAdd10000-16             4324136              5472 ns/op
PASS
ok      goasm   187.251s
```
