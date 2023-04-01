# Results:

```
$ go version
go version go1.20.2 linux/amd64
```

```
oha -z30s -c5000 http://machine:3000/?name=Oha
Summary:
  Success rate: 1.0000
  Total:        30.0018 secs
  Slowest:      29.1500 secs
  Fastest:      0.0012 secs
  Average:      1.3796 secs
  Requests/sec: 2820.5934

  Total data:   143.33 MiB
  Size/request: 1.73 KiB
  Size/sec:     4.78 MiB

Response time histogram:
   0.001 [1]     |
   2.916 [74034] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
   5.831 [7225]  |■■■
   8.746 [1792]  |
  11.661 [678]   |
  14.576 [377]   |
  17.490 [254]   |
  20.405 [145]   |
  23.320 [69]    |
  26.235 [36]    |
  29.150 [12]    |

Latency distribution:
  10% in 0.0637 secs
  25% in 0.1993 secs
  50% in 0.5962 secs
  75% in 1.6089 secs
  90% in 3.4012 secs
  95% in 5.1560 secs
  99% in 11.9624 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.2558 secs, 0.0003 secs, 1.0566 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200] 84623 responses
```