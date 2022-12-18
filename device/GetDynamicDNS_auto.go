// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetDynamicDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDynamicDNSResponse.
func Call_GetDynamicDNS(ctx context.Context, dev *Device, request GetDynamicDNS) (GetDynamicDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDynamicDNSResponse GetDynamicDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDynamicDNSResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetDynamicDNS")
		return reply.Body.GetDynamicDNSResponse, errors.Annotate(err, "reply")
	}
}