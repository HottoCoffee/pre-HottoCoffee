use async_trait::async_trait;

use crate::entity::domain::user::User;
use crate::entity::repository::user_repository::UserRepository;

pub struct DummyUserRepository;

#[async_trait]
impl UserRepository for DummyUserRepository {
    async fn find_by_email_and_password(&self, email: String, password: String) -> Option<User> {
        if email == "hoge@fuga.org" {
            Some(User { id: 1, display_name: "dummy".to_string(), email })
        } else {
            None
        }
    }
}
