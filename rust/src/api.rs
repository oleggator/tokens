extern crate rusty_tarantool;
extern crate serde_json;

use rusty_tarantool::tarantool::{Client, ExecWithParamaters};
use std::io;
use actix_web::{get, web, HttpResponse};
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, PartialEq, Serialize)]
pub struct TokenReq {
    token: String,
}

#[derive(Debug, Deserialize, PartialEq, Serialize)]
pub struct NewToken {
    token: String,
    expiration: u32,
}

#[get("/new")]
pub async fn new_token_handler(tarantool_client: web::Data<Client>) -> io::Result<HttpResponse> {
    let response =
        tarantool_client.prepare_fn_call("get_new_token")
            .execute()
            .await?;
    let new_token: NewToken = response.decode_single()?;
    Ok(HttpResponse::Ok().json(new_token))
}

#[get("/revoke")]
pub async fn revoke_token_handler(
    params: web::Query<TokenReq>,
    tarantool_client: web::Data<Client>,
) -> io::Result<HttpResponse> {
    let response =
        tarantool_client.prepare_fn_call("revoke_token")
            .bind_ref(&params.token)?
            .execute()
            .await?;
    let ok: bool = response.decode_single()?;
    Ok(HttpResponse::Ok().json(if ok { "ok" } else { "not ok" }))
}

#[get("/check")]
pub async fn check_token_handler(
    params: web::Query<TokenReq>,
    tarantool_client: web::Data<Client>,
) -> io::Result<HttpResponse> {
    let response =
        tarantool_client.call_fn1("check_token", &params.token).await?;
    let ok: bool = response.decode_single()?;
    Ok(HttpResponse::Ok().json(if ok { "ok" } else { "not ok" }))
}

#[get("/stub")]
pub async fn stub_handler() -> io::Result<HttpResponse> {
    Ok(HttpResponse::Ok().json("ok"))
}
