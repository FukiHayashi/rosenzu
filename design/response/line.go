package response

import (
	. "goa.design/goa/v3/dsl"
)

var Line = ResultType("application/vnd.line+json", func() {
	Attributes(func() {
		Attribute("name", String, "路線名", func() {
			Example("路線名")
		})
		Attribute("elements", CollectionOf(Element))
		Attribute("relations", CollectionOf(Relation))
		Attribute("operationalPoints", CollectionOf(OpelationalPoint))
	})
	View("default", func() {
		Attribute("name")
		Attribute("elements")
		Attribute("relations")
		Attribute("operationalPoints")
	})
})

var Element = ResultType("application/vnd.element+json", func() {
	Attributes(func() {
		Attribute("id", Int, "id", func() {
			Example(1)
		})
		Attribute("coordinates", CollectionOf(Coordinate))
	})
	View("default", func() {
		Attribute("coordinates")
	})
})

var Coordinate = ResultType("application/vnd.coordinate+json", func() {
	Attributes(func() {
		Attribute("longitude", Float32, "緯度", func() {
			Example(37.446707)
		})
		Attribute("latitude", Float32, "経度", func() {
			Example(138.854723)
		})
	})
	View("dafault", func() {
		Attribute("longitude")
		Attribute("latitude")
	})
})

var Relation = ResultType("application/vnd.relation+json", func() {
	Attributes(func() {
		Attribute("elementA", Int, "elementA", func() {
			Example(1)
		})
		Attribute("elementB", Int, "elementB", func() {
			Example(2)
		})
	})
	View("dafault", func() {
		Attribute("elementA")
		Attribute("elementB")
	})
})

var OpelationalPoint = ResultType("application/vnd.operationalpoint+json", func() {
	Attributes(func() {
		Attribute("name", String, "バス停名", func() {
			Example("バス停名")
		})
		Attribute("elementID", Int, "elementID", func() {
			Example(1)
		})
	})
	View("default", func() {
		Attribute("name")
		Attribute("elementID")
	})
})
