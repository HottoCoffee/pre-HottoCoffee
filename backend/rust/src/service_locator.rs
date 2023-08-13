use sqlx::MySqlPool;

use crate::infra::repository::user_repository_db_client::UserDbClient;
use crate::infra::util::cryptor::Cryptor;

#[derive(Clone)]
pub struct ServiceLocator {
    pub pool: MySqlPool,
    pub user_db_client: UserDbClient,
}

const CRYPT_KEY: String = String::from("pass");

impl ServiceLocator {
    pub async fn new(db_url: &str) -> ServiceLocator { // TODO: receive config as param
        let pool = MySqlPool::connect(db_url) // TODO: from env ver
            .await
            .expect("failed to connect DB");

        let cryptor = Cryptor::new(CRYPT_KEY);
        let user_repository = UserDbClient::new(pool.clone(), cryptor);
        ServiceLocator { pool, user_db_client: user_repository }
    }
}
