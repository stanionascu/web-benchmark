# Results:

```
$ rustup show
stable-x86_64-unknown-linux-gnu (default)
rustc 1.68.2 (9eb3afe9e 2023-03-27)
```

```
$ oha -z30s -c5000 http://machine:3000/?name=Oha
Summary:
  Success rate: 1.0000
  Total:        30.0016 secs
  Slowest:      12.3926 secs
  Fastest:      0.0012 secs
  Average:      0.4848 secs
  Requests/sec: 10031.8862

  Total data:   515.22 MiB
  Size/request: 1.75 KiB
  Size/sec:     17.17 MiB

Response time histogram:
   0.001 [1]      |
   1.240 [244294] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
   2.479 [51952]  |■■■■■■
   3.719 [1829]   |
   4.958 [552]    |
   6.197 [786]    |
   7.436 [123]    |
   8.675 [0]      |
   9.914 [485]    |
  11.153 [664]    |
  12.393 [287]    |

Latency distribution:
  10% in 0.0609 secs
  25% in 0.0792 secs
  50% in 0.0983 secs
  75% in 0.1390 secs
  90% in 1.7508 secs
  95% in 1.8406 secs
  99% in 3.6347 secs

Details (average, fastest, slowest):
  DNS+dialup:   4.2477 secs, 0.0057 secs, 10.4698 secs
  DNS-lookup:   4.2402 secs, 0.0052 secs, 9.4302 secs

Status code distribution:
  [200] 300973 responses
```