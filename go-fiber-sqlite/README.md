# Results

```
$ oha -z30s -c5000 http://machine:3000/?name=Oha
Summary:
  Success rate: 1.0000
  Total:        30.0010 secs
  Slowest:      27.8769 secs
  Fastest:      0.0011 secs
  Average:      1.6181 secs
  Requests/sec: 2445.2853

  Total data:   169.66 MiB
  Size/request: 2.37 KiB
  Size/sec:     5.66 MiB

Response time histogram:
   0.001 [1]     |
   2.789 [61052] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
   5.576 [8149]  |■■■■
   8.364 [2351]  |■
  11.151 [823]   |
  13.939 [372]   |
  16.727 [253]   |
  19.514 [144]   |
  22.302 [122]   |
  25.089 [65]    |
  27.877 [29]    |

Latency distribution:
  10% in 0.0838 secs
  25% in 0.2523 secs
  50% in 0.7404 secs
  75% in 1.9529 secs
  90% in 4.0038 secs
  95% in 5.9690 secs
  99% in 12.8376 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.2486 secs, 0.0009 secs, 1.0422 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200] 73361 responses
```
