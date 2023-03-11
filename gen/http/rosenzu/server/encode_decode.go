// Code generated by goa v3.11.1, DO NOT EDIT.
//
// rosenzu HTTP server encoders and decoders
//
// Command:
// $ goa gen rosenzu/design

package server

import (
	"context"
	"net/http"
	rosenzuviews "rosenzu/gen/rosenzu/views"

	goahttp "goa.design/goa/v3/http"
)

// EncodeFindResponse returns an encoder for responses returned by the rosenzu
// find endpoint.
func EncodeFindResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*rosenzuviews.Line)
		enc := encoder(ctx, w)
		body := NewFindResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeFindRequest returns a decoder for requests sent to the rosenzu find
// endpoint.
func DecodeFindRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			name string

			params = mux.Vars(r)
		)
		name = params["name"]
		payload := NewFindPayload(name)

		return payload, nil
	}
}

// marshalRosenzuviewsElementViewToElementResponseBody builds a value of type
// *ElementResponseBody from a value of type *rosenzuviews.ElementView.
func marshalRosenzuviewsElementViewToElementResponseBody(v *rosenzuviews.ElementView) *ElementResponseBody {
	if v == nil {
		return nil
	}
	res := &ElementResponseBody{}
	if v.Coordinates != nil {
		res.Coordinates = make([]*CoordinateResponseBody, len(v.Coordinates))
		for i, val := range v.Coordinates {
			res.Coordinates[i] = marshalRosenzuviewsCoordinateViewToCoordinateResponseBody(val)
		}
	}

	return res
}

// marshalRosenzuviewsCoordinateViewToCoordinateResponseBody builds a value of
// type *CoordinateResponseBody from a value of type
// *rosenzuviews.CoordinateView.
func marshalRosenzuviewsCoordinateViewToCoordinateResponseBody(v *rosenzuviews.CoordinateView) *CoordinateResponseBody {
	if v == nil {
		return nil
	}
	res := &CoordinateResponseBody{
		Longitude: v.Longitude,
		Latitude:  v.Latitude,
	}

	return res
}

// marshalRosenzuviewsRelationViewToRelationResponseBody builds a value of type
// *RelationResponseBody from a value of type *rosenzuviews.RelationView.
func marshalRosenzuviewsRelationViewToRelationResponseBody(v *rosenzuviews.RelationView) *RelationResponseBody {
	if v == nil {
		return nil
	}
	res := &RelationResponseBody{
		ElementA: v.ElementA,
		ElementB: v.ElementB,
	}

	return res
}

// marshalRosenzuviewsOperationalpointViewToOperationalpointResponseBody builds
// a value of type *OperationalpointResponseBody from a value of type
// *rosenzuviews.OperationalpointView.
func marshalRosenzuviewsOperationalpointViewToOperationalpointResponseBody(v *rosenzuviews.OperationalpointView) *OperationalpointResponseBody {
	if v == nil {
		return nil
	}
	res := &OperationalpointResponseBody{
		Name:      v.Name,
		ElementID: v.ElementID,
	}

	return res
}
