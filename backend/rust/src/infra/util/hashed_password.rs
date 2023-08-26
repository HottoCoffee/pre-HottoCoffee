pub struct HashedPassword {
    pub value: String,
}

const COST: u32 = 5;

impl HashedPassword {
    pub fn from_hashed_password(value: String) -> HashedPassword {
        HashedPassword { value }
    }

    pub fn from_plain_password(value: String) -> HashedPassword {
        let hashed_value = bcrypt::hash(value, COST).unwrap();
        HashedPassword { value: hashed_value }
    }

    pub fn verify(&self, plain_password: &String) -> bool {
        bcrypt::verify(plain_password, self.value.as_ref()).unwrap()
    }
}
