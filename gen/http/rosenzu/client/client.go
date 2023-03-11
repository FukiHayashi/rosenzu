// Code generated by goa v3.11.1, DO NOT EDIT.
//
// rosenzu client HTTP transport
//
// Command:
// $ goa gen rosenzu/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the rosenzu service endpoint HTTP clients.
type Client struct {
	// Find Doer is the HTTP client used to make requests to the find endpoint.
	FindDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the rosenzu service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		FindDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Find returns an endpoint that makes HTTP requests to the rosenzu service
// find server.
func (c *Client) Find() goa.Endpoint {
	var (
		decodeResponse = DecodeFindResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildFindRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.FindDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("rosenzu", "find", err)
		}
		return decodeResponse(resp)
	}
}