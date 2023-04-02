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

| Implementation                 | Result                  | Throughput              | Relative Performance |
| ------------------------------ | ----------------------- | ----------------------- | -------------------- |
| (node-sveltekit-sqlite)        | (603.4693 rps)          | (9.02 MiB)              | (0.06x)              |
| go-gin-sqlite                  | 2820.5934 rps           | 4.78 MiB                | 0.28x                |
| rust-axum-askama-sqlx          | 3227.3343 rps           | 5.52 MiB/s              | 0.32x                |
| python-uvicorn-fastapi-sqlite  | 3810.3847 rps           | 6.48 MiB/s              | 0.37x                |
| rust-axum-askama-rusqlite      | 10031.8862 rps          | 17.17 MiB/s             | 1.0x                 |

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
npm run buil
node build/index.js
```

# Creating the test SQLite database

To create the test SQLite database in the root folder:

```
sqlite> .open db.sqlite3
sqlite> .read db/sqlite-sakila-db/sqlite-sakila-schema.sql
sqlite> .read db/sqlite-sakila-db/sqlite-sakila-insert-data.sql
sqlite> .exit
```

# Test setup

Client machine: Ryzen 5900X / RAM 32GB

Server machine: Intel i3-8100T / RAM 12GB

Client is connected to the server with a 1GBps connection via switch in a homelab setup.

# WIP

- [ ] - Add CPU/Memory usage plot per implementation
- [X] - Rust
- [X] - Golang
- [ ] - C++
- [ ] - Bun
- [ ] - Deno
- [X] - Node
- [X] - Python
- [ ] - Common chart plotting all versions
