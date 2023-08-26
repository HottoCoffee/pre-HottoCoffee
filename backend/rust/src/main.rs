use axum::Router;
use axum::routing::post;
use sqlx::MySqlPool;

use crate::adopter::user_controller;
use crate::service_locator::ServiceLocator;

mod adopter;
mod infra;
mod entity;
mod service_locator;

#[tokio::main]
async fn main() {
    let service_locator = ServiceLocator::new("mysql://root:root@0.0.0.0:3306/hottocoffee").await;

    migrate_db(&service_locator.pool).await;

    let public_route = Router::new()
        .route("/sign-in", post(user_controller::sign_in))
        .route("/sign-up", post(user_controller::sign_up));

    let route = Router::new()
        .nest("/api/public", public_route)
        .with_state(service_locator);

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
