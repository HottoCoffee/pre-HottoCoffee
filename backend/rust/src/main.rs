use axum::{Json, Router, routing::{get, post}};
use axum::http::StatusCode;
use serde::Serialize;

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();

    let public_route = Router::new()
        .route("/sign-in", get(sign_in))
        .route("/sign-up", post(sign_up));

    let route = Router::new()
        .nest("/public", public_route);

    axum::Server::bind(&"0.0.0.0:8080".parse().unwrap())
        .serve(route.into_make_service())
        .await
        .unwrap();
}

async fn sign_in() -> (StatusCode, Json<User>) {
    (StatusCode::OK, Json(User { id: 0, username: "".to_string() }))
}

async fn sign_up() {
    todo!()
}

#[derive(Serialize)]
struct User {
    id: u64,
    username: String,
}
