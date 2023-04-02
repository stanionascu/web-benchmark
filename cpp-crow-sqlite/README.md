# Build
```
mkdir build
cd build/
cmake -DCMAKE_BUILD_TYPE=Release ../
cmake --build .
```

# Result

```
$ gcc --version
gcc (GCC) 12.2.1 20230201
```

```
$ oha -z30s -c5000 http://machine:3000/?name=Oha
Summary:
  Success rate: 1.0000
  Total:        30.0030 secs
  Slowest:      4.0593 secs
  Fastest:      0.0066 secs
  Average:      0.6706 secs
  Requests/sec: 6770.7954

  Total data:   348.33 MiB
  Size/request: 1.76 KiB
  Size/sec:     11.61 MiB

Response time histogram:
  0.007 [1]      |
  0.412 [6439]   |■
  0.817 [194454] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  1.222 [647]    |
  1.628 [1492]   |
  2.033 [0]      |
  2.438 [27]     |
  2.844 [0]      |
  3.249 [0]      |
  3.654 [0]      |
  4.059 [84]     |

Latency distribution:
  10% in 0.6336 secs
  25% in 0.6529 secs
  50% in 0.6666 secs
  75% in 0.7073 secs
  90% in 0.7253 secs
  95% in 0.7533 secs
  99% in 0.8524 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.4193 secs, 0.0003 secs, 2.0575 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [200] 203144 responses
```