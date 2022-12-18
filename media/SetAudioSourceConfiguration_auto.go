// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_SetAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioSourceConfigurationResponse.
func Call_SetAudioSourceConfiguration(ctx context.Context, dev *device.Device, request SetAudioSourceConfiguration) (SetAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioSourceConfigurationResponse SetAudioSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioSourceConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetAudioSourceConfiguration")
		return reply.Body.SetAudioSourceConfigurationResponse, errors.Annotate(err, "reply")
	}
}
