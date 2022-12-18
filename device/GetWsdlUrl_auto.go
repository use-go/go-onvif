// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetWsdlUrl forwards the call to dev.CallMethod() then parses the payload of the reply as a GetWsdlUrlResponse.
func Call_GetWsdlUrl(ctx context.Context, dev *Device, request GetWsdlUrl) (GetWsdlUrlResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetWsdlUrlResponse GetWsdlUrlResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetWsdlUrlResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetWsdlUrl")
		return reply.Body.GetWsdlUrlResponse, errors.Annotate(err, "reply")
	}
}
