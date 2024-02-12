package containerappssourcecontrols

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsSourceControlsClient struct {
	Client *resourcemanager.Client
}

func NewContainerAppsSourceControlsClientWithBaseURI(sdkApi sdkEnv.Api) (*ContainerAppsSourceControlsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "containerappssourcecontrols", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContainerAppsSourceControlsClient: %+v", err)
	}

	return &ContainerAppsSourceControlsClient{
		Client: client,
	}, nil
}
