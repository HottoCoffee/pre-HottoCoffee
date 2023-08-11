use async_trait::async_trait;
use derive_new::new;
use sqlx::MySqlPool;

use crate::entity::domain::user::User;
use crate::entity::repository::user_repository::UserRepository;
use crate::infra::util::cryptor::Cryptor;
use crate::infra::util::hashed_password::HashedPassword;

#[derive(new)]
pub struct UserDbClient {
    pool: MySqlPool,
    cryptor: Cryptor,
}

#[async_trait]
impl UserRepository for UserDbClient {
    async fn find_by_email_and_password(&self, email: &String, password: &String) -> Option<User> {
        let encrypted_email = self.cryptor.encrypt(email);
        sqlx::query_as::<_, UserRecord>("select * from user where email = ?")
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
}

#[derive(sqlx::FromRow)]
struct UserRecord {
    id: u32,
    display_name: String,
    email: Vec<u8>,
    password: String,
}
