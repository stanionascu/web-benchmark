# Results

This isn't a really fair comparison to the other implementations, but is kept
for visibility. It does quite a bit more than the other versions runs as a single worker
and thus is significantly slower.

```
$ node --version
v19.8.1

$ npm install
$ npm run build
$ PORT=3000 node build/index.js
```

```
Summary:
  Success rate: 1.0000
  Total:        30.0015 secs
  Slowest:      25.7916 secs
  Fastest:      0.6611 secs
  Average:      5.9215 secs
  Requests/sec: 603.4693

  Total data:   270.68 MiB
  Size/request: 15.31 KiB
  Size/sec:     9.02 MiB

Response time histogram:
   0.661 [1]    |
   3.174 [3524] |■■■■■■■■■■■■■■
   5.687 [4672] |■■■■■■■■■■■■■■■■■■
   8.200 [7869] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  10.713 [32]   |
  13.226 [1183] |■■■■
  15.739 [312]  |■
  18.252 [0]    |
  20.765 [0]    |
  23.279 [0]    |
  25.792 [512]  |■■

Latency distribution:
  10% in 1.7963 secs
  25% in 3.6378 secs
  50% in 5.8332 secs
  75% in 6.0431 secs
  90% in 11.6903 secs
  95% in 13.1420 secs
  99% in 25.5627 secs

Details (average, fastest, slowest):
  DNS+dialup:   5.0850 secs, 0.0004 secs, 16.4281 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200] 18105 response
```