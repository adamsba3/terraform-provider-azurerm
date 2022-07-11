package capacitypools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolPatchProperties struct {
	QosType *QosType `json:"qosType,omitempty"`
	Size    *int64   `json:"size,omitempty"`
}
