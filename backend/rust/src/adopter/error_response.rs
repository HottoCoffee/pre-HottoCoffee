use serde::Serialize;

#[derive(Serialize)]
pub struct ErrorResponse {
    status: u32,
    message: String,
}

impl ErrorResponse {
    pub fn new(status: u32, message: String) -> ErrorResponse {
        ErrorResponse { status, message }
    }
}
