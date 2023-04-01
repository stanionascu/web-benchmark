[![Rust](https://github.com/stanionascu/web-benchmark/actions/workflows/rust.yml/badge.svg)](https://github.com/stanionascu/web-benchmark/actions/workflows/rust.yml)

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

| Implementation                 | Result                  | Throughput              |
| ------------------------------ | ----------------------- | ----------------------- |
| go-gin-sqlite                  | 2820.5934 rps           | 4.78 MiB                |
| rust-axum-askama-sqlx          | 3227.3343 rps           | 5.52 MiB/s              |
| python-uvicorn-fastapi-sqlite  | 3810.3847 rps           | 6.48 MiB/s              |
| rust-axum-askama-rusqlite      | 10031.8862 rps          | 17.17 MiB/s             |

## Current implementations:
### Rust: axum + sqlite via sqlx + askama
```
cd rust-axum-askama-sqlite/
cargo run --release
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
- [ ] - Golang
- [ ] - C++
- [ ] - Bun
- [ ] - Deno
- [ ] - Node
- [X] - Python
- [ ] - Common chart plotting all versions
