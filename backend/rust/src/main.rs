use std::arch::asm;
use std::net::SocketAddr;

use axum::{http::StatusCode, Json, middleware, response::IntoResponse, Router, routing::{get, post}, ServiceExt};
use axum::handler::Handler;
use axum::routing::Route;
use serde::{Deserialize, Serialize};

mod authorization_layer;

#[tokio::main]
async fn main() {

    tracing_subscriber::fmt::init();

    let require_authorized_route = Router::new()
        .route("/", get(get_workspace))
        .route_layer(middleware::from_fn(authorization_layer::auth));

    let public_route = Router::new()
        .route("/sign-in", get(sign_in))
        .route("/sign-up", post(sign_up));

    let route = Router::new()
        .nest("/api/workspace", require_authorized_route)
        .nest("/public", public_route);

    axum::Server::bind(&"0.0.0.0:3000".parse().unwrap())
        .serve(route.into_make_service())
        .await
        .unwrap();
}

async fn authorize() {}

async fn get_workspace() -> &'static str {
    return "authorized";
}

async fn create_user(Json(payload): Json<CreateUser>) -> (StatusCode, Json<User>) {
    todo!()
}

async fn sign_in() {
    todo!()
}

async fn sign_up() {
    todo!()
}

#[derive(Deserialize)]
struct CreateUser {
    username: String,
}

#[derive(Serialize)]
struct User {
    id: u64,
    username: String,
}
