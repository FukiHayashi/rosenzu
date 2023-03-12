// Code generated by goa v3.11.1, DO NOT EDIT.
//
// rosenzu HTTP server types
//
// Command:
// $ goa gen rosenzu/design

package server

import (
	rosenzu "rosenzu/gen/rosenzu"
	rosenzuviews "rosenzu/gen/rosenzu/views"
)

// FindResponseBody is the type of the "rosenzu" service "find" endpoint HTTP
// response body.
type FindResponseBody struct {
	// 路線名
	Name              *string                                `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Elements          ElementResponseBodyCollection          `form:"elements,omitempty" json:"elements,omitempty" xml:"elements,omitempty"`
	Relations         RelationResponseBodyCollection         `form:"relations,omitempty" json:"relations,omitempty" xml:"relations,omitempty"`
	OperationalPoints OperationalpointResponseBodyCollection `form:"operationalPoints,omitempty" json:"operationalPoints,omitempty" xml:"operationalPoints,omitempty"`
}

// ElementResponseBodyCollection is used to define fields on response body
// types.
type ElementResponseBodyCollection []*ElementResponseBody

// ElementResponseBody is used to define fields on response body types.
type ElementResponseBody struct {
	// id
	ID          *uint                            `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Coordinates CoordinateResponseBodyCollection `form:"coordinates,omitempty" json:"coordinates,omitempty" xml:"coordinates,omitempty"`
}

// CoordinateResponseBodyCollection is used to define fields on response body
// types.
type CoordinateResponseBodyCollection []*CoordinateResponseBody

// CoordinateResponseBody is used to define fields on response body types.
type CoordinateResponseBody struct {
	// 緯度
	Longitude *float32 `form:"longitude,omitempty" json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 経度
	Latitude *float32 `form:"latitude,omitempty" json:"latitude,omitempty" xml:"latitude,omitempty"`
}

// RelationResponseBodyCollection is used to define fields on response body
// types.
type RelationResponseBodyCollection []*RelationResponseBody

// RelationResponseBody is used to define fields on response body types.
type RelationResponseBody struct {
	// elementA
	ElementA *uint `form:"elementA,omitempty" json:"elementA,omitempty" xml:"elementA,omitempty"`
	// elementB
	ElementB *uint `form:"elementB,omitempty" json:"elementB,omitempty" xml:"elementB,omitempty"`
}

// OperationalpointResponseBodyCollection is used to define fields on response
// body types.
type OperationalpointResponseBodyCollection []*OperationalpointResponseBody

// OperationalpointResponseBody is used to define fields on response body types.
type OperationalpointResponseBody struct {
	// バス停名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// elementID
	ElementID *uint `form:"elementID,omitempty" json:"elementID,omitempty" xml:"elementID,omitempty"`
}

// NewFindResponseBody builds the HTTP response body from the result of the
// "find" endpoint of the "rosenzu" service.
func NewFindResponseBody(res *rosenzuviews.LineView) *FindResponseBody {
	body := &FindResponseBody{
		Name: res.Name,
	}
	if res.Elements != nil {
		body.Elements = make([]*ElementResponseBody, len(res.Elements))
		for i, val := range res.Elements {
			body.Elements[i] = marshalRosenzuviewsElementViewToElementResponseBody(val)
		}
	}
	if res.Relations != nil {
		body.Relations = make([]*RelationResponseBody, len(res.Relations))
		for i, val := range res.Relations {
			body.Relations[i] = marshalRosenzuviewsRelationViewToRelationResponseBody(val)
		}
	}
	if res.OperationalPoints != nil {
		body.OperationalPoints = make([]*OperationalpointResponseBody, len(res.OperationalPoints))
		for i, val := range res.OperationalPoints {
			body.OperationalPoints[i] = marshalRosenzuviewsOperationalpointViewToOperationalpointResponseBody(val)
		}
	}
	return body
}

// NewFindPayload builds a rosenzu service find endpoint payload.
func NewFindPayload(name string) *rosenzu.FindPayload {
	v := &rosenzu.FindPayload{}
	v.Name = name

	return v
}
