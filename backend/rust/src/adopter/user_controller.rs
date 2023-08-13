use axum::extract::State;
use axum::http::StatusCode;
use axum::Json;
use serde::{Deserialize, Serialize};

use crate::adopter::jwt_claims::Jwt;
use crate::entity::repository::user_repository::UserRepository;
use crate::service_locator::ServiceLocator;

use super::error_response::ErrorResponse;

pub async fn sign_in(State(service_locator): State<ServiceLocator>, Json(request): Json<SignInUpRequest>)
                     -> Result<Json<UserResponse>, (StatusCode, Json<ErrorResponse>)> {
    let user_repository = service_locator.user_db_client;
    match user_repository.find_by_email_and_password(&request.email, &request.password).await {
        Some(user) => {
            let jwt = Jwt::issue(user.id);
            Ok(Json(UserResponse { token: jwt }))
        }
        None => {
            let not_found_response = ErrorResponse::new(404, "user not found".to_string());
            Err((StatusCode::NOT_FOUND, Json(not_found_response)))
        }
    }
}

pub async fn sign_up(State(service_locator): State<ServiceLocator>, Json(request): Json<SignInUpRequest>)
                     -> Result<Json<UserResponse>, (StatusCode, Json<ErrorResponse>)> {
    let user_repository = service_locator.user_db_client;
    user_repository.create(&request.email, &request.password)
        .await
        .map(|user| Json(UserResponse { token: Jwt::issue(user.id) }))
        .map_err(|_| (
            StatusCode::BAD_REQUEST,
            Json(ErrorResponse::new(400, String::from("email has been used")))
        ))
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
