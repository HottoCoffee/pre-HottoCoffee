use axum::Router;
use axum::routing::post;
use sqlx::MySqlPool;

use crate::adopter::user_controller;

mod adopter;
mod infra;
mod entity;

#[tokio::main]
async fn main() {
    let pool = MySqlPool::connect("mysql://root:root@0.0.0.0:3306/hottocoffee") // TODO: from env ver
        .await
        .expect("failed to connect DB");

    migrate_db(&pool).await;

    let public_route = Router::new()
        .route("/sign-in", post(user_controller::sign_in))
        .route("/sign-up", post(user_controller::sign_up));

    let route = Router::new()
        .nest("/api/public", public_route)
        .with_state(pool);

    axum::Server::bind(&"0.0.0.0:8080".parse().unwrap())
        .serve(route.into_make_service())
        .await
        .expect("failed to run app");
}

async fn migrate_db(pool: &MySqlPool) {
    sqlx::migrate!("./migrations")
        .run(pool)
        .await
        .expect("failed to migrate db")
}
