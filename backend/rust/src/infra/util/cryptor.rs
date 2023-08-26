use aes_gcm::aead::{Aead, NewAead};
use aes_gcm::Aes256Gcm;
use sha2::{Digest, Sha256};
use sha2::digest::generic_array::GenericArray;
use typenum::{U12, U32};

pub struct Cryptor {
    secret: GenericArray<u8, U32>,
}

impl Cryptor {
    pub fn new(key: String) -> Cryptor {
        let secret = {
            let mut hasher = Sha256::new();
            hasher.update(key.as_bytes());
            let result = hasher.finalize();
            GenericArray::clone_from_slice(&result)
        };
        Cryptor { secret }
    }

    pub fn encrypt(&self, str: &String) -> Vec<u8> {
        let cipher = Aes256Gcm::new(GenericArray::from_slice(&self.secret));
        let fixed_nonce: GenericArray<u8, U12> = GenericArray::from([0; 12]);

        cipher.encrypt(&fixed_nonce, str.as_bytes()).map(|mut v| {
            v.extend(&fixed_nonce);
            v
        }).unwrap()
    }

    pub fn decrypt(&self, bytes: &Vec<u8>) -> String {
        let cipher = Aes256Gcm::new(GenericArray::from_slice(&self.secret));
        let fixed_nonce: GenericArray<u8, U12> = GenericArray::from([0; 12]);

        String::from_utf8(cipher.decrypt(&fixed_nonce, &bytes[..bytes.len() - 12])
            .unwrap()
        ).unwrap()
    }
}
