# Go Prod Info Log Bench

tl;dr - **Avoid info logging WithFields in the hot path of performance sensitive applications.**

This repository contains tests to illustrate the performance costs of INFO and DEBUG logs in environments
where the log level is set to WARN.

---

I was reluctant to adopt Zap for a long time because... Log performance shouldn't matter.

If your app is writing so many logs in production that logging is a bottleneck, you've got bigger problems.

What is of concern to me is the performance cost of **ignored** logs in production.

What happens when production code encounters a statement like `log.Debug(err)`?

Production environments use log levels to ignore info and debug statements, but that log statement must still
be evaluated.

Results:

```
$ go test -bench=. -benchmem

BenchmarkZapInfo-12                     41377027             28.3 ns/op             0 B/op          0 allocs/op
BenchmarkZapSugarInfo-12                100000000            11.5 ns/op             0 B/op          0 allocs/op
BenchmarkLogrusInfo-12                  249465934             4.81 ns/op            0 B/op          0 allocs/op

BenchmarkZapInfoWithFields-12           11537806            104 ns/op             128 B/op          1 allocs/op
BenchmarkZapSugarInfoWithFields-12      24488396             47.0 ns/op            16 B/op          1 allocs/op
BenchmarkLogrusInfoWithFields-12         2154274            558 ns/op             512 B/op          5 allocs/op

BenchmarkZapWarn-12                      1308540            919 ns/op             216 B/op          2 allocs/op
BenchmarkZapSugarWarn-12                  999940           1210 ns/op             248 B/op          3 allocs/op
BenchmarkLogrusWarn-12                    499970           2406 ns/op             480 B/op         15 allocs/op

BenchmarkZapSugarWarnWithFields-12        857092           1439 ns/op             280 B/op          4 allocs/op
BenchmarkZapWarnWithFields-12             857092           1393 ns/op             344 B/op          3 allocs/op
BenchmarkLogrusWarnWithFields-12          333314           3492 ns/op            1009 B/op         21 allocs/op

PASS
ok      github.com/main 16.594s
```

Some interesting things:

1) Logrus is the most performant for ignored static info logging
2) Zap sugar is more performant than regular zap for ignored WithFields info logs in prod ¯\\\_(?)\_/¯
3) **Ignored WithFields info logging can cost 0.1 - 0.5 microseconds in production**

## Conclusion

- Info and Debug logging is not free in production.
- Avoid info logging WithFields in the hot path of performance sensitive applications.
- Solutions?

Have a nice day.
