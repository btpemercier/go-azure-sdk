package dataconnectors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnector interface {
}

// RawDataConnectorImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataConnectorImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalDataConnectorImplementation(input []byte) (DataConnector, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataConnector into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureActiveDirectory") {
		var out AADDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AADDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureAdvancedThreatProtection") {
		var out AATPDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AATPDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureSecurityCenter") {
		var out ASCDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ASCDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AmazonWebServicesCloudTrail") {
		var out AwsCloudTrailDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsCloudTrailDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AmazonWebServicesS3") {
		var out AwsS3DataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsS3DataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "APIPolling") {
		var out CodelessApiPollingDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CodelessApiPollingDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GenericUI") {
		var out CodelessUiDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CodelessUiDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Dynamics365") {
		var out Dynamics365DataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Dynamics365DataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MicrosoftCloudAppSecurity") {
		var out MCASDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MCASDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MicrosoftDefenderAdvancedThreatProtection") {
		var out MDATPDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MDATPDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MicrosoftThreatIntelligence") {
		var out MSTIDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MSTIDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MicrosoftThreatProtection") {
		var out MTPDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MTPDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OfficeATP") {
		var out OfficeATPDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfficeATPDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Office365") {
		var out OfficeDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfficeDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OfficeIRM") {
		var out OfficeIRMDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfficeIRMDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ThreatIntelligence") {
		var out TIDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TIDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ThreatIntelligenceTaxii") {
		var out TiTaxiiDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TiTaxiiDataConnector: %+v", err)
		}
		return out, nil
	}

	out := RawDataConnectorImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
