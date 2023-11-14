package appservicecertificateorders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RenewCertificateOrderRequestProperties struct {
	Csr                  *string `json:"csr,omitempty"`
	IsPrivateKeyExternal *bool   `json:"isPrivateKeyExternal,omitempty"`
	KeySize              *int64  `json:"keySize,omitempty"`
}
