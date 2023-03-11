package design

import (
	"rosenzu/design/response"

	. "goa.design/goa/v3/dsl"
)

var _ = Service("rosenzu", func() {
	Description("路線データ")
	Error("Unauthorized")
	Error("BadRequest")

	Method("find", func() {
		// ペイロードの定義
		Payload(func() {
			Attribute("name", String, func() {
				Description("路線名")
			})
			Required("name")
		})
		// 返却値
		Result(response.Line)
		// HTTPトランスポートの定義
		HTTP(func() {
			GET("/line/{name}")
			Response(StatusOK)
		})
	})
})
