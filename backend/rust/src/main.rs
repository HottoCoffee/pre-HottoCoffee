use sqlx::MySqlPool;

#[tokio::main]
async fn main() {
    let pool = MySqlPool::connect(&"mysql://root:root@0.0.0.0:3306/hottocoffee")
        .await
        .unwrap();

    sqlx::migrate!("./migrations")
        .run(&pool)
        .await
        .unwrap();
}
