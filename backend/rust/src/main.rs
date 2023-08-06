use axum::Router;
use axum::routing::post;
use sqlx::MySqlPool;

use crate::adopter::user_controller;

mod adopter;
mod infra;
mod entity;

#[tokio::main]
async fn main() {
    migrate_db().await;

    let public_route = Router::new()
        .route("/sign-in", post(user_controller::sign_in));

    let route = Router::new()
        .nest("/api/public", public_route);

    axum::Server::bind(&"0.0.0.0:8080".parse().unwrap())
        .serve(route.into_make_service())
        .await
        .expect("failed to run app");
}

async fn migrate_db() {
    let pool = MySqlPool::connect("mysql://root:root@0.0.0.0:3306/hottocoffee") // TODO: from env ver
        .await
        .unwrap();

    sqlx::migrate!("./migrations")
        .run(&pool)
        .await
        .expect("failed to migrate db")
}
