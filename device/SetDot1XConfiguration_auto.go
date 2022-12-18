// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_SetDot1XConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetDot1XConfigurationResponse.
func Call_SetDot1XConfiguration(ctx context.Context, dev *Device, request SetDot1XConfiguration) (SetDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDot1XConfigurationResponse SetDot1XConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetDot1XConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetDot1XConfiguration")
		return reply.Body.SetDot1XConfigurationResponse, errors.Annotate(err, "reply")
	}
}
