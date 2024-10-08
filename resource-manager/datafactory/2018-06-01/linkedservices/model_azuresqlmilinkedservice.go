package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LinkedService = AzureSqlMILinkedService{}

type AzureSqlMILinkedService struct {
	TypeProperties AzureSqlMILinkedServiceTypeProperties `json:"typeProperties"`

	// Fields inherited from LinkedService
	Annotations *[]interface{}                     `json:"annotations,omitempty"`
	ConnectVia  *IntegrationRuntimeReference       `json:"connectVia,omitempty"`
	Description *string                            `json:"description,omitempty"`
	Parameters  *map[string]ParameterSpecification `json:"parameters,omitempty"`
	Version     *string                            `json:"version,omitempty"`
}

var _ json.Marshaler = AzureSqlMILinkedService{}

func (s AzureSqlMILinkedService) MarshalJSON() ([]byte, error) {
	type wrapper AzureSqlMILinkedService
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureSqlMILinkedService: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureSqlMILinkedService: %+v", err)
	}
	decoded["type"] = "AzureSqlMI"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureSqlMILinkedService: %+v", err)
	}

	return encoded, nil
}
