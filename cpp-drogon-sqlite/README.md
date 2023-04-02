# Result

```
$ gcc --version
gcc (GCC) 12.2.1 20230201
```

## Build
```
mkdir build
cd build/
cmake -DCMAKE_BUILD_TYPE=Release ../
cmake --build .
```

## Generating CSP
From the build folder:

```
./drogon/bin/drogon_ctl create view ../templates/IndexHtml.csp -o ../
```

```
$ oha -z30s -c5000 http://machine:3000/?name=Oha
Summary:
  Success rate: 1.0000
  Total:        30.0016 secs
  Slowest:      29.7128 secs
  Fastest:      0.0030 secs
  Average:      0.2564 secs
  Requests/sec: 10745.5820

  Total data:   556.49 MiB
  Size/request: 1.77 KiB
  Size/sec:     18.55 MiB

Response time histogram:
   0.003 [1]      |
   2.974 [321928] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
   5.945 [55]     |
   8.916 [52]     |
  11.887 [53]     |
  14.858 [50]     |
  17.829 [51]     |
  20.800 [49]     |
  23.771 [50]     |
  26.742 [49]     |
  29.713 [47]     |

Latency distribution:
  10% in 0.2040 secs
  25% in 0.2169 secs
  50% in 0.2324 secs
  75% in 0.2495 secs
  90% in 0.2678 secs
  95% in 0.2784 secs
  99% in 0.2983 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0968 secs, 0.0016 secs, 0.2957 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200] 322385 responses
```