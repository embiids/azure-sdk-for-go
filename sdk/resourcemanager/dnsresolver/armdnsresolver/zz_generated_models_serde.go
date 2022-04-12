//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DNSForwardingRuleset.
func (d DNSForwardingRuleset) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", d.Etag)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DNSForwardingRulesetListResult.
func (d DNSForwardingRulesetListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DNSForwardingRulesetPatch.
func (d DNSForwardingRulesetPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DNSForwardingRulesetProperties.
func (d DNSForwardingRulesetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dnsResolverOutboundEndpoints", d.DNSResolverOutboundEndpoints)
	populate(objectMap, "provisioningState", d.ProvisioningState)
	populate(objectMap, "resourceGuid", d.ResourceGUID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DNSResolver.
func (d DNSResolver) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", d.Etag)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ForwardingRuleListResult.
func (f ForwardingRuleListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", f.NextLink)
	populate(objectMap, "value", f.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ForwardingRulePatch.
func (f ForwardingRulePatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", f.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ForwardingRulePatchProperties.
func (f ForwardingRulePatchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "forwardingRuleState", f.ForwardingRuleState)
	populate(objectMap, "metadata", f.Metadata)
	populate(objectMap, "targetDnsServers", f.TargetDNSServers)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ForwardingRuleProperties.
func (f ForwardingRuleProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "domainName", f.DomainName)
	populate(objectMap, "forwardingRuleState", f.ForwardingRuleState)
	populate(objectMap, "metadata", f.Metadata)
	populate(objectMap, "provisioningState", f.ProvisioningState)
	populate(objectMap, "targetDnsServers", f.TargetDNSServers)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InboundEndpoint.
func (i InboundEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", i.Etag)
	populate(objectMap, "id", i.ID)
	populate(objectMap, "location", i.Location)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "systemData", i.SystemData)
	populate(objectMap, "tags", i.Tags)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InboundEndpointListResult.
func (i InboundEndpointListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", i.NextLink)
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InboundEndpointPatch.
func (i InboundEndpointPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", i.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InboundEndpointProperties.
func (i InboundEndpointProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "ipConfigurations", i.IPConfigurations)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "resourceGuid", i.ResourceGUID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListResult.
func (l ListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundEndpoint.
func (o OutboundEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", o.Etag)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "location", o.Location)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "properties", o.Properties)
	populate(objectMap, "systemData", o.SystemData)
	populate(objectMap, "tags", o.Tags)
	populate(objectMap, "type", o.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundEndpointListResult.
func (o OutboundEndpointListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundEndpointPatch.
func (o OutboundEndpointPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", o.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Patch.
func (p Patch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", p.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubResourceListResult.
func (s SubResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkDNSForwardingRulesetListResult.
func (v VirtualNetworkDNSForwardingRulesetListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkLinkListResult.
func (v VirtualNetworkLinkListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkLinkPatch.
func (v VirtualNetworkLinkPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", v.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkLinkPatchProperties.
func (v VirtualNetworkLinkPatchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "metadata", v.Metadata)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkLinkProperties.
func (v VirtualNetworkLinkProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "metadata", v.Metadata)
	populate(objectMap, "provisioningState", v.ProvisioningState)
	populate(objectMap, "virtualNetwork", v.VirtualNetwork)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
