# Results

```
$ node --version
v19.8.1

$ npm install
$ node app.js
```

```
$ oha -z30s -c5000 http://machine:3000/?name=Oha
Summary:
  Success rate: 1.0000
  Total:        30.0018 secs
  Slowest:      29.7715 secs
  Fastest:      0.0808 secs
  Average:      0.6190 secs
  Requests/sec: 3842.0419

  Total data:   197.65 MiB
  Size/request: 1.76 KiB
  Size/sec:     6.59 MiB

Response time histogram:
   0.081 [1]      |
   3.050 [113719] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
   6.019 [263]    |
   8.988 [348]    |
  11.957 [157]    |
  14.926 [162]    |
  17.895 [181]    |
  20.864 [120]    |
  23.833 [108]    |
  26.802 [108]    |
  29.772 [101]    |

Latency distribution:
  10% in 0.2108 secs
  25% in 0.3219 secs
  50% in 0.4579 secs
  75% in 0.5745 secs
  90% in 0.6618 secs
  95% in 0.7177 secs
  99% in 6.0911 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.3559 secs, 0.0005 secs, 1.0872 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200] 115268 responses
```