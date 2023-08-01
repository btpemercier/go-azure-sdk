package kubeenvironments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubeEnvironmentsClient struct {
	Client *resourcemanager.Client
}

func NewKubeEnvironmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*KubeEnvironmentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "kubeenvironments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KubeEnvironmentsClient: %+v", err)
	}

	return &KubeEnvironmentsClient{
		Client: client,
	}, nil
}
