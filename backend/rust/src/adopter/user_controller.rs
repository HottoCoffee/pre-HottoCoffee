use axum::http::StatusCode;
use axum::Json;
use serde::{Deserialize, Serialize};

use crate::adopter::jwt_claims::Jwt;
use crate::entity::repository::user_repository::UserRepository;
use crate::infra::repository::dummy_user_repository::DummyUserRepository;

use super::error_response::ErrorResponse;

pub async fn sign_in(Json(request): Json<SignInUpRequest>)
                     -> Result<Json<UserResponse>, (StatusCode, Json<ErrorResponse>)> {
    let user_repository = DummyUserRepository {};
    let user = match user_repository.find_by_email_and_password(request.email, request.password).await {
        Some(user) => user,
        None => {
            let not_found_response = ErrorResponse::new(404, "user not found".to_string());
            return Err((StatusCode::NOT_FOUND, Json(not_found_response)));
        }
    };

    let jwt = Jwt::issue(user.id).await;
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
