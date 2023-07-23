use adaptor::user_controller::{self};
use axum::http::StatusCode;
use axum::{routing::post, Json, Router};
use chrono::{Duration, Local, NaiveDateTime};
use derive_new::new;
use jsonwebtoken::{Algorithm, EncodingKey, Header};
use ring::digest::{Context, SHA256};
use serde::{Deserialize, Serialize};
use sqlx::MySqlPool;

mod adaptor;
mod service;

const SALT: &str = "salt";
const JWT_SECRET: &str = "jwt";

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();

    let public_route = Router::new()
        .route("/sign-in", post(user_controller::sign_in))
        .route("/sign-up", post(sign_up));

    let route = Router::new().nest("/public", public_route);

    axum::Server::bind(&"0.0.0.0:8080".parse().unwrap())
        .serve(route.into_make_service())
        .await
        .unwrap();
}

async fn sign_in(Json(request): Json<SignInUpRequest>) -> (StatusCode, Json<UserResponse>) {
    let hashed_email = hash(request.email);
    let hashed_password = hash(request.password);

    let pool = MySqlPool::connect(&"mysql://root:root@0.0.0.0:3306/hottocoffee")
        .await
        .unwrap();

    let record =
        sqlx::query_as::<_, UserRecord>("select * from user where email = ? and password = ?")
            .bind(hashed_email)
            .bind(hashed_password)
            .fetch_one(&pool)
            .await
            .expect("");

    let jwt = make_jwt(record.id);

    (StatusCode::OK, Json(UserResponse::new(jwt)))
}

async fn sign_up(Json(request): Json<SignInUpRequest>) -> (StatusCode, Json<UserResponse>) {
    let hashed_email = hash(request.email);
    let hashed_password = hash(request.password);

    let pool = MySqlPool::connect(&"mysql://root:root@0.0.0.0:3306/hottocoffee")
        .await
        .unwrap();

    sqlx::query!(
        r#"
        insert into user(display_name, email, password)
        value ('user', ?, ?)
        "#,
        hashed_email,
        hashed_password
    )
    .execute(&pool)
    .await
    .unwrap();

    let user_id = sqlx::query_scalar("select last_insert_id()")
        .fetch_one(&pool)
        .await
        .unwrap();

    let jwt = make_jwt(user_id);

    return (StatusCode::OK, Json(UserResponse::new(jwt)));
}

fn hash(value: String) -> String {
    let salted_email = format!("{}{}", value, SALT);
    let mut context = Context::new(&SHA256);
    context.update(salted_email.as_bytes());
    let digest = context.finish();

    return hex::encode(digest.as_ref());
}

fn make_jwt(user_id: u32) -> String {
    let header = Header::new(Algorithm::HS512);
    let claims = JwtClaims::new(
        user_id,
        "HottoCoffee".to_string(),
        (Local::now() + Duration::minutes(30)).timestamp_nanos(),
    );
    let key = EncodingKey::from_secret(JWT_SECRET.as_ref());
    jsonwebtoken::encode(&header, &claims, &key).unwrap()
}

#[derive(Clone, new, Deserialize, Serialize)]
pub struct JwtClaims {
    user_id: u32,
    iss: String,
    exp: i64,
}

#[derive(Serialize, new)]
pub struct UserResponse {
    token: String,
}

#[derive(Deserialize)]
pub struct SignInUpRequest {
    email: String,
    password: String,
}

#[derive(sqlx::FromRow)]
pub struct UserRecord {
    id: u32,
    display_name: String,
    email: String,
    password: String,
}
