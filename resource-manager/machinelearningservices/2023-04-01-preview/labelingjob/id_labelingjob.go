package labelingjob

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = LabelingJobId{}

// LabelingJobId is a struct representing the Resource ID for a Labeling Job
type LabelingJobId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	LabelingJobName   string
}

// NewLabelingJobID returns a new LabelingJobId struct
func NewLabelingJobID(subscriptionId string, resourceGroupName string, workspaceName string, labelingJobName string) LabelingJobId {
	return LabelingJobId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		LabelingJobName:   labelingJobName,
	}
}

// ParseLabelingJobID parses 'input' into a LabelingJobId
func ParseLabelingJobID(input string) (*LabelingJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(LabelingJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LabelingJobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.LabelingJobName, ok = parsed.Parsed["labelingJobName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "labelingJobName", *parsed)
	}

	return &id, nil
}

// ParseLabelingJobIDInsensitively parses 'input' case-insensitively into a LabelingJobId
// note: this method should only be used for API response data and not user input
func ParseLabelingJobIDInsensitively(input string) (*LabelingJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(LabelingJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LabelingJobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.LabelingJobName, ok = parsed.Parsed["labelingJobName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "labelingJobName", *parsed)
	}

	return &id, nil
}

// ValidateLabelingJobID checks that 'input' can be parsed as a Labeling Job ID
func ValidateLabelingJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLabelingJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Labeling Job ID
func (id LabelingJobId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/labelingJobs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.LabelingJobName)
}

// Segments returns a slice of Resource ID Segments which comprise this Labeling Job ID
func (id LabelingJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticLabelingJobs", "labelingJobs", "labelingJobs"),
		resourceids.UserSpecifiedSegment("labelingJobName", "labelingJobValue"),
	}
}

// String returns a human-readable description of this Labeling Job ID
func (id LabelingJobId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Labeling Job Name: %q", id.LabelingJobName),
	}
	return fmt.Sprintf("Labeling Job (%s)", strings.Join(components, "\n"))
}
