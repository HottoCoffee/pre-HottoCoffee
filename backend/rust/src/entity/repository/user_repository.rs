use async_trait::async_trait;

use crate::entity::domain::user::User;

#[async_trait]
pub trait UserRepository {
    async fn find_by_email_and_password(&self, email: &String, password: &String) -> Option<User>;
    async fn create(&self, email: &String, password: &String) -> Result<User, UserAlreadyExistError>;
}

pub struct UserAlreadyExistError {}
