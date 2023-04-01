# Running

```
$ python --version
Python 3.10.10
$ pip install -r requirements.txt
```

# Results

```
$ uvicorn serve:app --port 3000 --host '::' --workers 4 --log-level critical
```

```
oha -z30s -c5000 http://machine:3000/?name=Oha
Summary:
  Success rate: 1.0000
  Total:        30.0020 secs
  Slowest:      3.4290 secs
  Fastest:      0.0448 secs
  Average:      1.2880 secs
  Requests/sec: 3810.3847

  Total data:   194.50 MiB
  Size/request: 1.74 KiB
  Size/sec:     6.48 MiB

Response time histogram:
  0.045 [1]     |
  0.383 [407]   |
  0.722 [19167] |■■■■■■■■■■■■■■■■■■■■■
  1.060 [18316] |■■■■■■■■■■■■■■■■■■■■
  1.398 [28189] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  1.737 [25497] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  2.075 [11441] |■■■■■■■■■■■■
  2.414 [10262] |■■■■■■■■■■■
  2.752 [1028]  |■
  3.091 [7]     |
  3.429 [4]     |

Latency distribution:
  10% in 0.5824 secs
  25% in 0.8284 secs
  50% in 1.2176 secs
  75% in 1.6500 secs
  90% in 2.0683 secs
  95% in 2.1161 secs
  99% in 2.3595 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.1849 secs, 0.0004 secs, 1.0548 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0025 secs

Status code distribution:
  [200] 114319 responses
```