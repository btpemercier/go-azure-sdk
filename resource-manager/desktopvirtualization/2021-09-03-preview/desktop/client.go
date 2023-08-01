package desktop

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DesktopClient struct {
	Client *resourcemanager.Client
}

func NewDesktopClientWithBaseURI(sdkApi sdkEnv.Api) (*DesktopClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "desktop", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DesktopClient: %+v", err)
	}

	return &DesktopClient{
		Client: client,
	}, nil
}
