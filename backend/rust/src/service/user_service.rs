use crate::{JwtClaims, SignInUpRequest, UserRecord, UserResponse, JWT_SECRET};
use axum::body::HttpBody;
use chrono::NaiveDateTime;
use ring::digest::{Context, SHA256};
use sqlx::{query_as, MySqlPool};

pub(crate) struct UserService {
    db_pool: MySqlPool,
}

pub struct QueryableUser {
    pub id: u32,
    pub display_name: String,
    pub email: String,
    pub password: String,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
    pub deleted_at: Option<NaiveDateTime>,
}

impl UserService {
    pub fn new(db_pool: MySqlPool) -> Self {
        UserService { db_pool }
    }

    pub async fn sign_in(&self, email: String, password: String) -> Option<QueryableUser> {
        let hashed_email = hash(email);
        let hashed_password = hash(password);

        // let record = query_as!(
        //     QueryableUser,
        //     r#"
        //         select * from user where email = ? and password = ?
        //     "#,
        //     hashed_email,
        //     hashed_password
        // )
        // .fetch_one(&self.db_pool)
        // .await;

        // match record {
        //     Ok(record) => Some(record),
        //     Err(_) => None,
        // }

        None
    }
}

fn hash(value: String) -> String {
    let salted_email = format!("{}{}", value, "SALT");
    let mut context = Context::new(&SHA256);
    context.update(salted_email.as_bytes());
    let digest = context.finish();

    return hex::encode(digest.as_ref());
}
