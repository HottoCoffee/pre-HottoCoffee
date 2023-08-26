use async_trait::async_trait;
use derive_new::new;
use sqlx::MySqlPool;

use crate::entity::domain::user::User;
use crate::entity::repository::user_repository::{UserAlreadyExistError, UserRepository};
use crate::infra::util::cryptor::Cryptor;
use crate::infra::util::hashed_password::HashedPassword;

#[derive(new, Clone)]
pub struct UserDbClient {
    pool: MySqlPool,
    cryptor: Cryptor,
}

#[async_trait]
impl UserRepository for UserDbClient {
    async fn find_by_email_and_password(&self, email: &String, password: &String) -> Option<User> {
        let encrypted_email = self.cryptor.encrypt(email);
        sqlx::query_as::<_, UserRecord>("select * from user where email = ? and deleted_at is null")
            .bind(encrypted_email)
            .fetch_one(&self.pool)
            .await
            .ok()
            .filter(|it| HashedPassword::from_hashed_password(it.password.clone()).verify(password))
            .map(|it| User {
                id: it.id,
                display_name: it.display_name,
                email: self.cryptor.decrypt(&it.email),
            })
    }

    async fn create(&self, email: &String, password: &String) -> Result<User, UserAlreadyExistError> {
        let encrypted_email = self.cryptor.encrypt(email);
        let hashed_password = HashedPassword::from_plain_password(password.clone()).value;

        let tx = self.pool.begin().await.unwrap();
        let user_count = sqlx::query_scalar::<_, i64>("select count(*) from user where email = ?")
            .bind(&encrypted_email)
            .fetch_one(&self.pool)
            .await
            .unwrap();

        if user_count > 0 {
            tx.commit().await.unwrap();
            return Err(UserAlreadyExistError {});
        }

        sqlx::query(
            r#"
            insert into user(display_name, email, password)
            value ('user', ?, ?)
            "#
        ).bind(&encrypted_email)
            .bind(hashed_password)
            .execute(&self.pool)
            .await
            .unwrap();

        let user_id = sqlx::query_scalar::<_, u32>("select last_insert_id()")
            .fetch_one(&self.pool)
            .await
            .unwrap();

        tx.commit().await.unwrap();

        let initial_username = &email[0..email.chars().position(|c| c == '@').unwrap()];
        Ok(User { id: user_id, display_name: initial_username.to_string(), email: email.clone() })
    }
}

#[derive(sqlx::FromRow)]
struct UserRecord {
    id: u32,
    display_name: String,
    email: Vec<u8>,
    password: String,
}
