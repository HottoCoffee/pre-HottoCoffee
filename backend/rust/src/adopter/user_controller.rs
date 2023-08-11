use axum::extract::State;
use axum::http::StatusCode;
use axum::Json;
use serde::{Deserialize, Serialize};
use sqlx::MySqlPool;

use crate::adopter::jwt_claims::Jwt;
use crate::entity::repository::user_repository::UserRepository;
use crate::infra::repository::user_repository_db_client::UserDbClient;
use crate::infra::util::cryptor::Cryptor;
use crate::infra::util::hashed_password::HashedPassword;

use super::error_response::ErrorResponse;

pub async fn sign_in(State(pool): State<MySqlPool>, Json(request): Json<SignInUpRequest>)
                     -> Result<Json<UserResponse>, (StatusCode, Json<ErrorResponse>)> {
    let cryptor = Cryptor::new("pass".to_string());

    let user_repository = UserDbClient::new(pool.clone(), cryptor);
    match user_repository.find_by_email_and_password(&request.email, &request.password).await {
        Some(user) => {
            let jwt = Jwt::issue(user.id);
            Ok(Json(UserResponse { token: jwt }))
        }
        None => {
            let not_found_response = ErrorResponse::new(404, "user not found".to_string());
            return Err((StatusCode::NOT_FOUND, Json(not_found_response)));
        }
    }
}

pub async fn sign_up(State(pool): State<MySqlPool>, Json(request): Json<SignInUpRequest>)
                     -> Result<Json<UserResponse>, (StatusCode, Json<ErrorResponse>)> {
    let cryptor = Cryptor::new("pass".to_string());
    let encrypted_email = cryptor.encrypt(&request.email);
    let hashed_password = HashedPassword::from_plain_password(request.password);

    let tx = pool.begin().await.unwrap();
    sqlx::query!(
        r#"
        insert into user(display_name, email, password)
        value ('user', ?, ?)
        "#,
        encrypted_email, hashed_password.value
    )
        .execute(&pool)
        .await
        .unwrap();

    let user_id = sqlx::query_scalar("select last_insert_id()")
        .fetch_one(&pool)
        .await
        .unwrap();

    tx.commit().await.unwrap();

    let jwt = Jwt::issue(user_id);
    Ok(Json(UserResponse { token: jwt }))
}

#[derive(Deserialize)]
pub struct SignInUpRequest {
    email: String,
    password: String,
}

#[derive(Serialize)]
pub struct UserResponse {
    token: String,
}
