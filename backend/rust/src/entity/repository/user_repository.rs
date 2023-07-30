use async_trait::async_trait;

use crate::entity::domain::user::User;

#[async_trait]
pub trait UserRepository {
    async fn find_by_id(&self, user_id: u32) -> User;
}
