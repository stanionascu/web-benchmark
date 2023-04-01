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
  Total:        30.0019 secs
  Slowest:      4.6723 secs
  Fastest:      0.0351 secs
  Average:      1.5003 secs
  Requests/sec: 3227.3343

  Total data:   165.75 MiB
  Size/request: 1.75 KiB
  Size/sec:     5.52 MiB

Response time histogram:
  0.035 [1]     |
  0.499 [1807]  |
  0.963 [4075]  |■
  1.426 [1197]  |
  1.890 [87631] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  2.354 [1384]  |
  2.817 [463]   |
  3.281 [42]    |
  3.745 [0]     |
  4.209 [0]     |
  4.672 [226]   |

Latency distribution:
  10% in 1.4699 secs
  25% in 1.5099 secs
  50% in 1.5302 secs
  75% in 1.5599 secs
  90% in 1.5805 secs
  95% in 1.6000 secs
  99% in 2.2891 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.6364 secs, 0.0006 secs, 3.1568 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200] 96826 responses
```