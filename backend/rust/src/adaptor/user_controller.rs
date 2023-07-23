use axum::http::StatusCode;
use axum::Json;
use chrono::{Duration, Local};
use jsonwebtoken::{Algorithm, EncodingKey, Header};
use ring::digest::{Context, SHA256};
use sqlx::MySqlPool;

use crate::{service::user_service, JwtClaims, SignInUpRequest, UserResponse, JWT_SECRET};

pub async fn sign_in(Json(request): Json<SignInUpRequest>) -> (StatusCode, Json<UserResponse>) {
    let db_conn = MySqlPool::connect(&"mysql://root:root@localhost:3306/hottocoffee")
        .await
        .unwrap();
    let user_service = user_service::UserService::new(db_conn);
    let record = user_service
        .sign_in(request.email, request.password)
        .await
        .unwrap();

    let jwt = make_jwt(record.id);

    (StatusCode::OK, Json(UserResponse::new(jwt)))
}

fn hash(value: String) -> String {
    let salted_email = format!("{}{}", value, "SALT");
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
