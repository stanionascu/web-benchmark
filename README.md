[![Rust](https://github.com/stanionascu/web-benchmark/actions/workflows/rust.yml/badge.svg)](https://github.com/stanionascu/web-benchmark/actions/workflows/rust.yml)

# About

This web benchmark suite focuses on a realistic SSR use-case, when the data is read
from a database, and rendered as an html page. Then this page is served to the client.

Only first 100 records are fetched from the Sakila Sample Dataset - film table.

Frontend html uses PicoCSS, which is served from CDN. There might be some minor differences inside
the templates, due to different templating languages. But the rendered output must be identical.

# Results

All tests are running with Oha for 30 seconds:
```
$ oha -z30s http://machine:3000/?name=Oha
```

| Implementation         | Result                  |
| ---------------------- | ----------------------- |
| rust-axum-askama-sqlx  | 4321.3984 rps           |

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
- [ ] - Python
- [ ] - Common chart plotting all versions
