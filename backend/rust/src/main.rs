use axum::{Json, Router, routing::{get, post}};
use axum::http::StatusCode;
use chrono::{DateTime, Local, NaiveDateTime};
use serde::Serialize;
use sqlx::{MySqlPool, Row};

#[tokio::main]
async fn main() {
    let pool = MySqlPool::connect(&"mysql://root:root@0.0.0.0:3306/hottocoffee")
        .await
        .unwrap();

    let x = sqlx::query_as!(UserRecord, "select * from user")
        .fetch_one(&pool)
        .await.unwrap();
    println!("{:?}", x);

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

#[derive(Debug)]
struct UserRecord {
    id: u32,
    display_name: String,
    email: Box<[u8]>,
    password: Box<[u8]>,
    created_at: NaiveDateTime,
    updated_at: NaiveDateTime,
    deleted_at: Option<NaiveDateTime>,
}
