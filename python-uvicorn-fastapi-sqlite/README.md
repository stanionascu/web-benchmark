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
oha -z30s http://icebox:3000/?name=Oha
Summary:
  Success rate: 1.0000
  Total:        30.0006 secs
  Slowest:      1.0759 secs
  Fastest:      0.0015 secs
  Average:      0.0088 secs
  Requests/sec: 5688.1239

  Total data:   2.73 GiB
  Size/request: 16.80 KiB
  Size/sec:     93.35 MiB

Response time histogram:
  0.001 [1]      |
  0.109 [170645] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.216 [0]      |
  0.324 [0]      |
  0.431 [0]      |
  0.539 [0]      |
  0.646 [0]      |
  0.754 [0]      |
  0.861 [0]      |
  0.968 [0]      |
  1.076 [1]      |

Latency distribution:
  10% in 0.0065 secs
  25% in 0.0072 secs
  50% in 0.0082 secs
  75% in 0.0102 secs
  90% in 0.0112 secs
  95% in 0.0117 secs
  99% in 0.0191 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0360 secs, 0.0024 secs, 1.0672 secs
  DNS-lookup:   0.0149 secs, 0.0019 secs, 0.0277 secs

Status code distribution:
  [200] 170647 responses
```