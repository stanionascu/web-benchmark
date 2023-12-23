# Results:

```
$ go version
go version go1.20.2 linux/amd64
```

```
$ oha -z30s -c5000 http://machine:3000/?name=Oha
Summary:
  Success rate: 100.00%
  Total:        30.0015 secs
  Slowest:      21.9191 secs
  Fastest:      0.0009 secs
  Average:      0.4008 secs
  Requests/sec: 11648.3316

  Total data:   579.90 MiB
  Size/request: 1.70 KiB
  Size/sec:     19.33 MiB

Response time histogram:
   0.001 [1]      |
   2.193 [339162] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
   4.385 [6017]   |
   6.576 [2300]   |
   8.768 [990]    |
  10.960 [438]    |
  13.152 [244]    |
  15.344 [169]    |
  17.535 [106]    |
  19.727 [28]     |
  21.919 [12]     |

Response time distribution:
  10.00% in 0.0193 secs
  25.00% in 0.0470 secs
  50.00% in 0.1337 secs
  75.00% in 0.3663 secs
  90.00% in 0.8504 secs
  95.00% in 1.4323 secs
  99.00% in 4.9685 secs
  99.90% in 12.8490 secs
  99.99% in 17.8079 secs


Details (average, fastest, slowest):
  DNS+dialup:   0.0765 secs, 0.0005 secs, 0.3107 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0141 secs

Status code distribution:
  [200] 349467 responses
```