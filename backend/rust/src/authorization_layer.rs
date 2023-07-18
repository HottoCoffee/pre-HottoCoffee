use derive;
use std::ops::Add;

use axum::{Extension, http};
use axum::http::{Request, StatusCode};
use axum::middleware::Next;
use axum::response::Response;
use axum::routing::delete;
use chrono::{DateTime, Duration, FixedOffset, Local};
use derive_new::new;
use hmac::digest::KeyInit;
use hmac::digest::typenum::op;
use hmac::Hmac;
use jsonwebtoken::{Algorithm, DecodingKey, EncodingKey, Header, Validation};
use serde::{Deserialize, Serialize};
use tracing_subscriber::fmt::time;

#[derive(Clone, new, Deserialize, Serialize)]
struct JwtClaims {
    user_id: u32,
    iss: String,
    exp: i64,
}

pub async fn auth<B>(mut req: Request<B>, next: Next<B>) -> Result<Response, StatusCode> {
    let auth_header = req.headers()
        .get(http::header::AUTHORIZATION)
        .and_then(|header| header.to_str().ok());

    let jwt = match auth_header.filter(|it| it.starts_with("Bearer")) // もしSomeでBearerじゃなければNone
        .map(|it| it.replace("Bearer ", "")) {
        Some(value) => value,
        None => return Err(StatusCode::UNAUTHORIZED)
    };

    match authorize_current_user(jwt).await {
        Some(value) => {
            req.extensions_mut().insert(value);
            Ok(next.run(req).await)
        }
        None => Err(StatusCode::UNAUTHORIZED)
    }
}

async fn authorize_current_user(jwt: String) -> Option<JwtClaims> {
    let key = DecodingKey::from_secret("hogefugapiyo".as_ref());
    let mut my_validation = Validation::new(Algorithm::HS512);
    my_validation.set_issuer(&vec!["HottoCoffee"]);

    let debug = jsonwebtoken::decode::<JwtClaims>(&jwt, &key, &my_validation);
    if debug.is_err() {
        println!("{}", debug.clone().err().unwrap());
        let header = Header::new(Algorithm::HS512);
        let claims = JwtClaims::new(30, "HottoCoffee".to_string(), (Local::now() + Duration::minutes(30)).timestamp_nanos());
        let encoding_key = EncodingKey::from_secret("hogefugapiyo".as_ref());
        println!("{}", jsonwebtoken::encode(&header, &claims, &encoding_key).unwrap());
    }
    return debug.map(|it| it.claims).ok();
}
