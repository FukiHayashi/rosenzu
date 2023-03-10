package design

import (
	. "goa.design/goa/v3/dsl"
)

// API 定義
var _ = API("rosenzu", func() {
	// API の説明（タイトルと説明）
	Title("rosenzu Example")
	Description("路線図を書くためのデータ")
	Version("v1")
	// サーバ定義
	Server("rosenzu", func() {
		Host("localhost", func() {
			URI("http://localhost:8000/api/v1") // HTTP REST API
		})
	})
})
