# Results:

```
$ rustup show
stable-x86_64-unknown-linux-gnu (default)
rustc 1.68.2 (9eb3afe9e 2023-03-27)
```

```
$ oha -z30s http://machine:3000/?name=Oha
Summary:
  Success rate: 1.0000
  Total:        30.0002 secs
  Slowest:      0.0777 secs
  Fastest:      0.0008 secs
  Average:      0.0116 secs
  Requests/sec: 4321.3984

  Total data:   2.12 GiB
  Size/request: 17.15 KiB
  Size/sec:     72.38 MiB

Response time histogram:
  0.001 [1]     |
  0.008 [22462] |■■■■■■■
  0.016 [93023] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.024 [13016] |■■■■
  0.032 [1002]  |
  0.039 [105]   |
  0.047 [16]    |
  0.055 [5]     |
  0.062 [6]     |
  0.070 [4]     |
  0.078 [3]     |

Latency distribution:
  10% in 0.0069 secs
  25% in 0.0095 secs
  50% in 0.0111 secs
  75% in 0.0135 secs
  90% in 0.0165 secs
  95% in 0.0187 secs
  99% in 0.0235 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0312 secs, 0.0026 secs, 0.0611 secs
  DNS-lookup:   0.0308 secs, 0.0020 secs, 0.0610 secs

Status code distribution:
  [200] 129643 responses
```