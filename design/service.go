package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("rosenzu", func() {
	Description("路線データ")
	Error("Unauthorized")
	Error("BadRequest")

	Method("find", func() {
		// ペイロードの定義
		Payload(func() {
			Attribute("id", String, func() {
				Description("id")
			})
			Required("id")
		})
		// 返却値
		Result(Int)
		// HTTPトランスポートの定義
		HTTP(func() {
			GET("/line/{id}")
			Response(StatusOK)
		})
	})
})
