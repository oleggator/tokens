# Token storage
- 8 cores (Intel Xeon Processor (Skylake, IBRS))
- 8gb ram

## Rust
- actix-web
- rusty_tarantool (https://crates.io/crates/rusty_tarantool)

```
wrk http://localhost:8080/check?token=cbc4a2e0-4ead-4574-92d6-a0f681b159fd -t8 -c256 -d 5m --latency
Running 5m test @ http://localhost:8080/check?token=cbc4a2e0-4ead-4574-92d6-a0f681b159fd
  8 threads and 256 connections
^C  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.11ms    9.68ms 380.27ms   98.49%
    Req/Sec    14.09k     3.07k   22.97k    68.26%
  Latency Distribution
     50%    2.09ms
     75%    2.74ms
     90%    3.60ms
     99%   25.90ms
  8674577 requests in 1.29m, 0.90GB read
Requests/sec: 112153.17
Transfer/sec:     11.87MB
```

```
wrk http://localhost:8080/new -t8 -c256 -d 5m --latency
Running 5m test @ http://localhost:8080/new
  8 threads and 256 connections
^C  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.13ms    4.73ms 215.85ms   99.26%
    Req/Sec    11.09k     1.54k   22.86k    70.57%
  Latency Distribution
     50%    2.71ms
     75%    3.40ms
     90%    4.28ms
     99%    7.02ms
  3885380 requests in 44.04s, 641.03MB read
Requests/sec:  88227.58
Transfer/sec:     14.56MB
```

```
wrk http://localhost:8080/stub -t8 -c256 -d 5m --latency
Running 5m test @ http://localhost:8080/stub
  8 threads and 256 connections
^C  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   787.95us  703.71us  22.74ms   90.96%
    Req/Sec    39.17k     8.06k   70.31k    67.96%
  Latency Distribution
     50%  617.00us
     75%  848.00us
     90%    1.37ms
     99%    3.73ms
  11358436 requests in 36.46s, 1.17GB read
Requests/sec: 311557.21
Transfer/sec:     32.98MB
```


## Go (tarantool/go-tarantool)
- tarantool/go-tarantool
- valyala/fasthttp
- fasthttp/router

```
wrk http://localhost:8080/check?token=cbc4a2e0-4ead-4574-92d6-a0f681b159fd -t8 -c256 -d 5m --latency
Running 5m test @ http://localhost:8080/check?token=cbc4a2e0-4ead-4574-92d6-a0f681b159fd
  8 threads and 256 connections
^C  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.88ms    1.44ms  59.76ms   89.66%
    Req/Sec    17.52k     3.27k   29.82k    67.37%
  Latency Distribution
     50%    1.62ms
     75%    2.28ms
     90%    3.17ms
     99%    5.88ms
  4119024 requests in 29.56s, 534.24MB read
Requests/sec: 139349.06
Transfer/sec:     18.07MB
```

```
wrk http://localhost:8080/new -t8 -c256 -d 5m --latency
Running 5m test @ http://localhost:8080/new
  8 threads and 256 connections
^C  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.60ms    1.66ms  68.23ms   85.54%
    Req/Sec    12.62k     2.01k   18.09k    64.90%
  Latency Distribution
     50%    2.28ms
     75%    3.26ms
     90%    4.32ms
     99%    7.07ms
  2874477 requests in 28.65s, 529.07MB read
Requests/sec: 100314.26
Transfer/sec:     18.46MB
```

```
wrk http://localhost:8080/stub -t8 -c256 -d 5m --latency
Running 5m test @ http://localhost:8080/stub
  8 threads and 256 connections
^C  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.99ms    1.14ms  40.22ms   88.89%
    Req/Sec    37.48k    16.13k  103.82k    72.48%
  Latency Distribution
     50%  745.00us
     75%    1.28ms
     90%    2.25ms
     99%    5.21ms
  6984857 requests in 23.48s, 0.88GB read
Requests/sec: 297423.94
Transfer/sec:     38.58MB
```

## Go (FZambia/tarantool)
- FZambia/tarantool
- valyala/fasthttp
- fasthttp/router
```
wrk http://localhost:8080/check?token=cbc4a2e0-4ead-4574-92d6-a0f681b159fd -t8 -c256 -d 5m --latency
Running 5m test @ http://localhost:8080/check?token=cbc4a2e0-4ead-4574-92d6-a0f681b159fd
  8 threads and 256 connections
^C  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.91ms    1.37ms  56.82ms   87.85%
    Req/Sec    17.17k     3.40k   85.93k    70.32%
  Latency Distribution
     50%    1.65ms
     75%    2.32ms
     90%    3.22ms
     99%    6.01ms
  5037244 requests in 36.93s, 653.33MB read
Requests/sec: 136392.64
Transfer/sec:     17.69MB

```

```
wrk http://localhost:8080/stub -t8 -c256 -d 5m --latency
Running 5m test @ http://localhost:8080/stub
  8 threads and 256 connections
    Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.95ms    1.16ms  40.39ms   89.46%
    Req/Sec    39.29k    17.84k  109.12k    71.50%
  Latency Distribution
     50%  702.00us
     75%    1.22ms
     90%    2.18ms
     99%    5.12ms
  10849671 requests in 34.74s, 1.37GB read
Requests/sec: 312350.18
Transfer/sec:     40.51MB
```

```
wrk http://localhost:8080/new -t8 -c256 -d 5m --latency
Running 5m test @ http://localhost:8080/new
  8 threads and 256 connections
    Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.63ms    1.40ms  17.74ms   75.07%
    Req/Sec    12.38k     1.92k   22.51k    67.59%
  Latency Distribution
     50%    2.32ms
     75%    3.32ms
     90%    4.44ms
     99%    7.23ms
  2810113 requests in 28.54s, 517.23MB read
Requests/sec:  98473.69
Transfer/sec:     18.12MB
```
