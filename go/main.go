package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/tarantool/go-tarantool"
	"github.com/valyala/fasthttp"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type API struct {
	tntConn *tarantool.Connection
}

type NewToken struct {
	Token      string
	Expiration int64
}

func (newToken *NewToken) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	if _, err = d.DecodeArrayLen(); err != nil {
		return err
	}
	if newToken.Token, err = d.DecodeString(); err != nil {
		return err
	}
	if newToken.Expiration, err = d.DecodeInt64(); err != nil {
		return err
	}

	return nil
}

func (api *API) NewTokenHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	var result [][]NewToken
	if err := api.tntConn.Call17Typed("libprocedures.get_new_token", []interface{}{}, &result); err != nil {
		log.Println(err)
		ctx.SetStatusCode(500)
		_, _ = fmt.Fprintf(ctx, "%s", err)
		return
	}

	token := result[0][0]
	_, _ = fmt.Fprintf(ctx, `{"token: "%s", "expiration": %d}`, token.Token, token.Expiration)
}

func (api *API) RevokeTokenHandler(ctx *fasthttp.RequestCtx) {
	tokenUUID := string(ctx.QueryArgs().Peek("token"))
	var result []bool
	if err := api.tntConn.Call17Typed("revoke_token", []string{tokenUUID}, &result); err != nil {
		ctx.SetStatusCode(500)
		_, _ = fmt.Fprintf(ctx, "%s", err)
		return
	}
	if ok := result[0]; ok {
		ctx.WriteString("ok")
	} else {
		ctx.WriteString("not ok")
	}
}

func (api *API) CheckTokenHandler(ctx *fasthttp.RequestCtx) {
	tokenUUID := string(ctx.QueryArgs().Peek("token"))
	var result []bool
	if err := api.tntConn.Call17Typed("check_token", []string{tokenUUID}, &result); err != nil {
		ctx.SetStatusCode(500)
		_, _ = fmt.Fprintf(ctx, "%s", err)
		return
	}
	if ok := result[0]; ok {
		ctx.WriteString("ok")
	} else {
		ctx.WriteString("not ok")
	}
}

func (api *API) StubHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("ok")
}

func main() {
	conn, err := tarantool.Connect("localhost:3301", tarantool.Opts{
		Concurrency: 32,
		User:        "tokens",
		Pass:        "tokens",
	})
	if err != nil {
		log.Fatalln(err)
	}

	api := API{
		tntConn: conn,
	}

	r := router.New()
	r.GET("/new", api.NewTokenHandler)
	r.GET("/stub", api.StubHandler)
	r.GET("/revoke", api.RevokeTokenHandler)
	r.GET("/check", api.CheckTokenHandler)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
