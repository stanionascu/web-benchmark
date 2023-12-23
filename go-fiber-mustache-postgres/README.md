# Results

```
$ oha -z30s -c5000 http://machine:3000/?name=Oha
Summary:
  Success rate: 100.00%
  Total:        30.0018 secs
  Slowest:      28.2461 secs
  Fastest:      0.0011 secs
  Average:      0.5412 secs
  Requests/sec: 8374.6098

  Total data:   583.22 MiB
  Size/request: 2.38 KiB
  Size/sec:     19.44 MiB

Response time histogram:
   0.001 [1]      |
   2.826 [243339] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
   5.650 [4888]   |
   8.475 [1625]   |
  11.299 [792]    |
  14.124 [354]    |
  16.948 [155]    |
  19.773 [55]     |
  22.597 [22]     |
  25.422 [14]     |
  28.246 [8]      |

Response time distribution:
  10.00% in 0.0242 secs
  25.00% in 0.0644 secs
  50.00% in 0.1907 secs
  75.00% in 0.5217 secs
  90.00% in 1.1896 secs
  95.00% in 1.9955 secs
  99.00% in 6.2757 secs
  99.90% in 14.1568 secs
  99.99% in 22.3402 secs


Details (average, fastest, slowest):
  DNS+dialup:   0.1063 secs, 0.0005 secs, 1.0653 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0117 secs

Status code distribution:
  [200] 251253 responses
```
