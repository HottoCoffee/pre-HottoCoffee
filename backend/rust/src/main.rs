use axum::{Extension, Router};
use axum::routing::post;
use sqlx::{MySql, MySqlPool, Pool};

use crate::adopter::user_controller;

mod adopter;
mod infra;
mod entity;

#[tokio::main]
async fn main() {
    let pool = MySqlPool::connect(&"mysql://root:root@0.0.0.0:3306/hottocoffee") // TODO: from env ver
        .await
        .unwrap();

    migrate_db(&pool).await;

    let public_route = Router::new()
        .route("/sign-in", post(user_controller::sign_in));
    // .layer(Extension(Arc::new(pool)));

    let route = Router::new()
        .nest("/api/public", public_route)
        .layer(Extension(pool));

    axum::Server::bind(&"0.0.0.0:8080".parse().unwrap())
        .serve(route.into_make_service())
        .await
        .expect("failed to run app");
}

async fn migrate_db(pool: &Pool<MySql>) {
    sqlx::migrate!("./migrations")
        .run(pool)
        .await
        .expect("failed to migrate db")
}
