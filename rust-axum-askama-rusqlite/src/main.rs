use std::sync::{Arc, Mutex};

use askama::Template;

use axum::{
    extract::{Query, State},
    routing::get,
    Router,
};
use serde::Deserialize;
use tower::ServiceBuilder;
use tower_http::{compression::CompressionLayer, ServiceBuilderExt};

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
    db: Arc<Mutex<rusqlite::Connection>>,
}

#[tokio::main]
async fn main() {
    let conn = rusqlite::Connection::open("../db.sqlite3").unwrap();
    let state = AppState {
        db: Arc::new(Mutex::new(conn)),
    };

    let middleware = ServiceBuilder::new()
        .layer(CompressionLayer::new())
        .map_response_body(axum::body::boxed)
        .compression();

    let app = Router::new()
        .route("/", get(handler))
        .layer(middleware)
        .with_state(state);
    let addr = "[::]:3000".parse::<std::net::SocketAddr>().unwrap();

    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn handler(
    Query(params): Query<Params>,
    State(state): State<AppState>,
) -> Result<IndexTemplate, ()> {
    let db = state.db.lock().unwrap();
    let mut stmt = db
        .prepare("SELECT title,release_year FROM film LIMIT 100")
        .unwrap();
    let mut rows = stmt.query([]).unwrap();
    let mut films = Vec::new();
    while let Some(row) = rows.next().unwrap() {
        films.push(Film {
            title: row.get(0).unwrap(),
            release_year: row.get(1).unwrap(),
        });
    }

    return Ok(IndexTemplate {
        name: params.name.unwrap_or("None".to_string()),
        films,
    });
}
