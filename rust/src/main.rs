mod api;

use actix_web::{App, HttpServer};
use rusty_tarantool::tarantool::{ClientConfig};

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let tarantool_client =
        ClientConfig::new("127.0.0.1:3301", "tokens", "tokens")
            .set_timeout_time_ms(2000)
            .set_reconnect_time_ms(2000)
            .build();

    HttpServer::new(move || {
        App::new()
            .data(tarantool_client.clone())
            .service(api::new_token_handler)
            .service(api::revoke_token_handler)
            .service(api::check_token_handler)
            .service(api::stub_handler)
    })
        .workers(8)
        .bind("127.0.0.1:8080")?
        .run()
        .await
}
