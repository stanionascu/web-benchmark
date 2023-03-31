use askama::Template;
use axum::{
    extract::{Query, State},
    routing::get,
    Router,
};
use serde::Deserialize;
use sqlx::{sqlite::SqliteRow, Row, SqlitePool};

#[derive(Template)]
#[template(path = "index.html")]
struct IndexTemplate {
    name: String,
    films: Vec<Film>,
}

struct Film {
    title: String,
    release_year: String,
}

#[derive(Deserialize)]
struct Params {
    name: Option<String>,
}

#[derive(Clone)]
struct AppState {
    db: SqlitePool,
}

#[tokio::main]
async fn main() {
    let state = AppState {
        db: SqlitePool::connect("sqlite:../db.sqlite3").await.unwrap(),
    };
    let app = Router::new().route("/", get(handler)).with_state(state);
    let addr = "[::]:3000".parse::<std::net::SocketAddr>().unwrap();

    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn handler(Query(params): Query<Params>, State(state): State<AppState>) -> IndexTemplate {
    let films = sqlx::query("SELECT title,release_year FROM film LIMIT 100")
        .map(|row: SqliteRow| {
            return Film {
                title: row.get(0),
                release_year: row.get(1),
            };
        })
        .fetch_all(&state.db)
        .await.unwrap();

    return IndexTemplate {
        name: params.name.unwrap_or("None".to_string()),
        films,
    };
}
