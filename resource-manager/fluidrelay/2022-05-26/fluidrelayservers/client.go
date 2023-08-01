package fluidrelayservers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FluidRelayServersClient struct {
	Client *resourcemanager.Client
}

func NewFluidRelayServersClientWithBaseURI(sdkApi sdkEnv.Api) (*FluidRelayServersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "fluidrelayservers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FluidRelayServersClient: %+v", err)
	}

	return &FluidRelayServersClient{
		Client: client,
	}, nil
}
