// Code generated by goa v3.11.1, DO NOT EDIT.
//
// rosenzu HTTP client CLI support package
//
// Command:
// $ goa gen rosenzu/design

package client

import (
	rosenzu "rosenzu/gen/rosenzu"
)

// BuildFindPayload builds the payload for the rosenzu find endpoint from CLI
// flags.
func BuildFindPayload(rosenzuFindName string) (*rosenzu.FindPayload, error) {
	var name string
	{
		name = rosenzuFindName
	}
	v := &rosenzu.FindPayload{}
	v.Name = name

	return v, nil
}