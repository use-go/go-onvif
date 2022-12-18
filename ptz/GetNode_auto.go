// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_GetNode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNodeResponse.
func Call_GetNode(ctx context.Context, dev *device.Device, request GetNode) (GetNodeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNodeResponse GetNodeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNodeResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetNode")
		return reply.Body.GetNodeResponse, errors.Annotate(err, "reply")
	}
}
