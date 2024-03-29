[![Rust](https://github.com/stanionascu/web-benchmark/actions/workflows/rust.yml/badge.svg)](https://github.com/stanionascu/web-benchmark/actions/workflows/rust.yml)
[![Go](https://github.com/stanionascu/web-benchmark/actions/workflows/go.yml/badge.svg)](https://github.com/stanionascu/web-benchmark/actions/workflows/go.yml)

# About

This web benchmark suite focuses on a realistic SSR use-case, when the data is read
from a database, and rendered as an html page. Then this page is served to the client.

Only first 100 records are fetched from the Sakila Sample Dataset - film table.

Frontend html uses PicoCSS, which is served from CDN. There might be some minor differences inside
the templates, due to different templating languages. But the rendered output must be identical.

### Update:
* Compression is enabled as the payload * rps was exceed the 1Gbps link capacity.
* Number of connections from Oha, increased to 5000

# Results

All tests are running with Oha for 30 seconds, and 5000 connections:
```
$ oha -z30s -c5000 http://machine:3000/?name=Oha
```

| Implementation                 | Result                  | Throughput              | Relative Performance | Latency p10/p50/p99 (sec)  |
| ------------------------------ | ----------------------- | ----------------------- | -------------------- | -------------------------- |
| (node-sveltekit-sqlite)        | (603.4693 rps)          | (9.02 MiB/s)            | (0.06x)              | 1.79/5.83/25.56 s          |
| go-fiber-sqlite                | 2445.2853 rps           | 5.66 MiB/s              | 0.24x                | 0.08/0.74/12.83 s          |
| go-gin-sqlite                  | 2820.5934 rps           | 4.78 MiB/s              | 0.28x                | 0.06/0.59/11.96 s          |
| rust-axum-askama-sqlx          | 3227.3343 rps           | 5.52 MiB/s              | 0.32x                | 1.47/1.53/2.29 s           |
| python-uvicorn-fastapi-sqlite  | 3810.3847 rps           | 6.48 MiB/s              | 0.37x                | 0.58/1.22/2.36 s           |
| node-fastify-handlebars-sqlite | 3842.0419 rps           | 6.59 MiB/s              | 0.38x                | 0.21/0.46/6.09 s           |
| cpp-crow-sqlite                | 6770.7954 rps           | 11.61 MiB/s             | 0.67x                | 0.63/0.66/0.85 s           |
| go-fiber-mustache-postgres     | 8374.6098 rps           | 19.44 MiB/s             | 0.83x                | 0.02/0.19/6.28 s           |
| rust-axum-askama-rusqlite      | 10031.8862 rps          | 17.17 MiB/s             | 1.0x                 | 0.06/0.09/3.63 s           |
| cpp-drogon-sqlite              | 10745.5820 rps          | 18.55 MiB/s             | 1.07x                | 0.20/0.23/0.29 s           |
| go-gin-templ-postgres          | 11648.3316 rps          | 19.33 MiB/s             | 1.16x                | 0.02/0.13/4.97 s           |

## Current implementations:
### Rust: axum + sqlite via sqlx + askama
```
cd rust-axum-askama-sqlite/
cargo run --release
```

### Rust: axum + sqlite via rusqlite + askama
```
cd rust-axum-askama-rusqlite/
cargo run --release
```

### Python: uvicorn + fastapi + sqlite
```
cd python-uvicorn-fastapi-sqlite/
pip -r requirements.txt
uvicorn serve:app --port 3000 --host '::' --workers 4 --log-level critical
```

### Go: go-gin + sqlite3
```
cd go-gin-sqlite/
GIN_MODE=release go run main.go
```

### Node: node + sveltekit + sqlite3
```
cd node-sveltekit-sqlite/
npm install
npm run build
node build/index.js
```

### Node: node + fastify + handlebars + sqlite3
```
cd node-fastify-handlebars-sqlite/
npm install
node app.js
```

### Go: go-fiber + sqlite3
```
cd go-fiber-sqlite/
go run main.go
```

### Go: go-fiber + sqlx (postgres)
```
cd go-fiber-postgres/
go run main.go
```

### Go: go-gin + templ + sqlx (postgres)

**NOTE**: Templ needs a pregeneration step.

```
cd go-gin-templ-postgres/
templ generate -f index.templ
go run main.go index_templ.go
```


# Creating the test SQLite database

To create the test SQLite database in the root folder:

```
sqlite> .open db.sqlite3
sqlite> .read db/sqlite-sakila-db/sqlite-sakila-schema.sql
sqlite> .read db/sqlite-sakila-db/sqlite-sakila-insert-data.sql
sqlite> .exit
```

# Creating the test Postgres database (docker)

```
docker run --name db -p 5432:5432 -e POSTGRES_PASSWORD=<password> -d postgres
psql -h localhost -U postgres
psql> CREATE DATABASE sakila;
psql> \c sakila;
psql#sakila> \i db/postgres-sakila-db/postgres-sakila-schema.sql
psql#sakila> \i db/postgres-sakila-db/postgres-sakila-insert-data.sql
```

# Test setup

Client machine: Ryzen 5900X / RAM 32GB

Server machine: Intel i3-8100T / RAM 12GB

Client is connected to the server with a 1GBps connection via switch in a homelab setup.

# WIP

- [ ] - Add CPU/Memory usage plot per implementation
- [X] - Rust
- [X] - Golang
- [X] - C++
- [ ] - Bun
- [ ] - Deno
- [X] - Node
- [X] - Python
- [ ] - Common chart plotting all versions
