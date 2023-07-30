use async_trait::async_trait;

use crate::entity::domain::user::User;
use crate::entity::repository::user_repository::UserRepository;

struct DummyUserRepository;

#[async_trait]
impl UserRepository for DummyUserRepository {
    async fn find_by_id(&self, user_id: u32) -> User {
        todo!()
    }
}
