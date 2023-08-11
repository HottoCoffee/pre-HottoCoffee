use std::string::ToString;

use chrono::{Duration, Local};
use jsonwebtoken::{Algorithm, EncodingKey, Header};
use serde::{Deserialize, Serialize};

const ISSUER: &str = "HottoCoffee";
const JWT_SECRET: &str = "JWT"; // TODO: from env var

#[derive(Clone, Deserialize, Serialize)]
pub struct Jwt {
    user_id: u32,
    iss: String,
    exp: i64,
}

impl Jwt {
    pub fn issue(user_id: u32) -> String {
        let header = Header::new(Algorithm::HS512);
        let claims = Jwt::new(user_id);

        let key = EncodingKey::from_secret(JWT_SECRET.as_ref());
        jsonwebtoken::encode(&header, &claims, &key).unwrap()
    }

    fn new(user_id: u32) -> Jwt {
        Jwt {
            user_id,
            iss: ISSUER.to_string(),
            exp: (Local::now() + Duration::minutes(30)).timestamp_nanos(),
        }
    }
}
