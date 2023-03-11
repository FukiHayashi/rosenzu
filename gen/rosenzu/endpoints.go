// Code generated by goa v3.11.1, DO NOT EDIT.
//
// rosenzu endpoints
//
// Command:
// $ goa gen rosenzu/design

package rosenzu

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "rosenzu" service endpoints.
type Endpoints struct {
	Find goa.Endpoint
}

// NewEndpoints wraps the methods of the "rosenzu" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Find: NewFindEndpoint(s),
	}
}

// Use applies the given middleware to all the "rosenzu" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Find = m(e.Find)
}

// NewFindEndpoint returns an endpoint function that calls the method "find" of
// service "rosenzu".
func NewFindEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FindPayload)
		res, err := s.Find(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedLine(res, "default")
		return vres, nil
	}
}