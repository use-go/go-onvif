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

// Call_GetMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMetadataConfigurationResponse.
func Call_GetMetadataConfiguration(ctx context.Context, dev *device.Device, request GetMetadataConfiguration) (GetMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationResponse GetMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetadataConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetMetadataConfiguration")
		return reply.Body.GetMetadataConfigurationResponse, errors.Annotate(err, "reply")
	}
}