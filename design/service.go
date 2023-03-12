package design

import (
	"rosenzu/design/response"

	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = Service("rosenzu", func() {
	Description("路線データ")
	Error("Unauthorized")
	Error("BadRequest")
	cors.Origin("*", func() {
		cors.Methods("GET")
		cors.Headers("Accept", "Content-Type")
		cors.Expose("Content-Type", "Origin")
		cors.Credentials()
	})

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
