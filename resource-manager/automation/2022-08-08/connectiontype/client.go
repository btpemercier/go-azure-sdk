package connectiontype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionTypeClient struct {
	Client *resourcemanager.Client
}

func NewConnectionTypeClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectionTypeClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "connectiontype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectionTypeClient: %+v", err)
	}

	return &ConnectionTypeClient{
		Client: client,
	}, nil
}
