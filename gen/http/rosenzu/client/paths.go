// Code generated by goa v3.11.1, DO NOT EDIT.
//
// HTTP request path constructors for the rosenzu service.
//
// Command:
// $ goa gen rosenzu/design

package client

import (
	"fmt"
)

// FindRosenzuPath returns the URL path to the rosenzu service find HTTP endpoint.
func FindRosenzuPath(name string) string {
	return fmt.Sprintf("/line/%v", name)
}
