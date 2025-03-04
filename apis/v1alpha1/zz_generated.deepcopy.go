//go:build !ignore_autogenerated

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIStage) DeepCopyInto(out *APIStage) {
	*out = *in
	if in.APIID != nil {
		in, out := &in.APIID, &out.APIID
		*out = new(string)
		**out = **in
	}
	if in.Stage != nil {
		in, out := &in.Stage, &out.Stage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIStage.
func (in *APIStage) DeepCopy() *APIStage {
	if in == nil {
		return nil
	}
	out := new(APIStage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessLogSettings) DeepCopyInto(out *AccessLogSettings) {
	*out = *in
	if in.DestinationARN != nil {
		in, out := &in.DestinationARN, &out.DestinationARN
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessLogSettings.
func (in *AccessLogSettings) DeepCopy() *AccessLogSettings {
	if in == nil {
		return nil
	}
	out := new(AccessLogSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySettings) DeepCopyInto(out *CanarySettings) {
	*out = *in
	if in.DeploymentID != nil {
		in, out := &in.DeploymentID, &out.DeploymentID
		*out = new(string)
		**out = **in
	}
	if in.StageVariableOverrides != nil {
		in, out := &in.StageVariableOverrides, &out.StageVariableOverrides
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.UseStageCache != nil {
		in, out := &in.UseStageCache, &out.UseStageCache
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySettings.
func (in *CanarySettings) DeepCopy() *CanarySettings {
	if in == nil {
		return nil
	}
	out := new(CanarySettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentCanarySettings) DeepCopyInto(out *DeploymentCanarySettings) {
	*out = *in
	if in.StageVariableOverrides != nil {
		in, out := &in.StageVariableOverrides, &out.StageVariableOverrides
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.UseStageCache != nil {
		in, out := &in.UseStageCache, &out.UseStageCache
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentCanarySettings.
func (in *DeploymentCanarySettings) DeepCopy() *DeploymentCanarySettings {
	if in == nil {
		return nil
	}
	out := new(DeploymentCanarySettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocumentationPartLocation) DeepCopyInto(out *DocumentationPartLocation) {
	*out = *in
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocumentationPartLocation.
func (in *DocumentationPartLocation) DeepCopy() *DocumentationPartLocation {
	if in == nil {
		return nil
	}
	out := new(DocumentationPartLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointConfiguration) DeepCopyInto(out *EndpointConfiguration) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VPCEndpointIDs != nil {
		in, out := &in.VPCEndpointIDs, &out.VPCEndpointIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VPCEndpointRefs != nil {
		in, out := &in.VPCEndpointRefs, &out.VPCEndpointRefs
		*out = make([]*corev1alpha1.AWSResourceReferenceWrapper, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.AWSResourceReferenceWrapper)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointConfiguration.
func (in *EndpointConfiguration) DeepCopy() *EndpointConfiguration {
	if in == nil {
		return nil
	}
	out := new(EndpointConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MethodSetting) DeepCopyInto(out *MethodSetting) {
	*out = *in
	if in.CacheDataEncrypted != nil {
		in, out := &in.CacheDataEncrypted, &out.CacheDataEncrypted
		*out = new(bool)
		**out = **in
	}
	if in.CachingEnabled != nil {
		in, out := &in.CachingEnabled, &out.CachingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DataTraceEnabled != nil {
		in, out := &in.DataTraceEnabled, &out.DataTraceEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LoggingLevel != nil {
		in, out := &in.LoggingLevel, &out.LoggingLevel
		*out = new(string)
		**out = **in
	}
	if in.MetricsEnabled != nil {
		in, out := &in.MetricsEnabled, &out.MetricsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RequireAuthorizationForCacheControl != nil {
		in, out := &in.RequireAuthorizationForCacheControl, &out.RequireAuthorizationForCacheControl
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MethodSetting.
func (in *MethodSetting) DeepCopy() *MethodSetting {
	if in == nil {
		return nil
	}
	out := new(MethodSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MethodSnapshot) DeepCopyInto(out *MethodSnapshot) {
	*out = *in
	if in.APIKeyRequired != nil {
		in, out := &in.APIKeyRequired, &out.APIKeyRequired
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizationType != nil {
		in, out := &in.AuthorizationType, &out.AuthorizationType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MethodSnapshot.
func (in *MethodSnapshot) DeepCopy() *MethodSnapshot {
	if in == nil {
		return nil
	}
	out := new(MethodSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutualTLSAuthentication) DeepCopyInto(out *MutualTLSAuthentication) {
	*out = *in
	if in.TruststoreURI != nil {
		in, out := &in.TruststoreURI, &out.TruststoreURI
		*out = new(string)
		**out = **in
	}
	if in.TruststoreVersion != nil {
		in, out := &in.TruststoreVersion, &out.TruststoreVersion
		*out = new(string)
		**out = **in
	}
	if in.TruststoreWarnings != nil {
		in, out := &in.TruststoreWarnings, &out.TruststoreWarnings
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutualTLSAuthentication.
func (in *MutualTLSAuthentication) DeepCopy() *MutualTLSAuthentication {
	if in == nil {
		return nil
	}
	out := new(MutualTLSAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutualTLSAuthenticationInput) DeepCopyInto(out *MutualTLSAuthenticationInput) {
	*out = *in
	if in.TruststoreURI != nil {
		in, out := &in.TruststoreURI, &out.TruststoreURI
		*out = new(string)
		**out = **in
	}
	if in.TruststoreVersion != nil {
		in, out := &in.TruststoreVersion, &out.TruststoreVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutualTLSAuthenticationInput.
func (in *MutualTLSAuthenticationInput) DeepCopy() *MutualTLSAuthenticationInput {
	if in == nil {
		return nil
	}
	out := new(MutualTLSAuthenticationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchOperation) DeepCopyInto(out *PatchOperation) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.Op != nil {
		in, out := &in.Op, &out.Op
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchOperation.
func (in *PatchOperation) DeepCopy() *PatchOperation {
	if in == nil {
		return nil
	}
	out := new(PatchOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestAPI) DeepCopyInto(out *RestAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestAPI.
func (in *RestAPI) DeepCopy() *RestAPI {
	if in == nil {
		return nil
	}
	out := new(RestAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestAPI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestAPIList) DeepCopyInto(out *RestAPIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RestAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestAPIList.
func (in *RestAPIList) DeepCopy() *RestAPIList {
	if in == nil {
		return nil
	}
	out := new(RestAPIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestAPIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestAPISpec) DeepCopyInto(out *RestAPISpec) {
	*out = *in
	if in.APIKeySource != nil {
		in, out := &in.APIKeySource, &out.APIKeySource
		*out = new(string)
		**out = **in
	}
	if in.BinaryMediaTypes != nil {
		in, out := &in.BinaryMediaTypes, &out.BinaryMediaTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CloneFrom != nil {
		in, out := &in.CloneFrom, &out.CloneFrom
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableExecuteAPIEndpoint != nil {
		in, out := &in.DisableExecuteAPIEndpoint, &out.DisableExecuteAPIEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.EndpointConfiguration != nil {
		in, out := &in.EndpointConfiguration, &out.EndpointConfiguration
		*out = new(EndpointConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.MinimumCompressionSize != nil {
		in, out := &in.MinimumCompressionSize, &out.MinimumCompressionSize
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestAPISpec.
func (in *RestAPISpec) DeepCopy() *RestAPISpec {
	if in == nil {
		return nil
	}
	out := new(RestAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestAPIStatus) DeepCopyInto(out *RestAPIStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = (*in).DeepCopy()
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RootResourceID != nil {
		in, out := &in.RootResourceID, &out.RootResourceID
		*out = new(string)
		**out = **in
	}
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestAPIStatus.
func (in *RestAPIStatus) DeepCopy() *RestAPIStatus {
	if in == nil {
		return nil
	}
	out := new(RestAPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestAPI_SDK) DeepCopyInto(out *RestAPI_SDK) {
	*out = *in
	if in.APIKeySource != nil {
		in, out := &in.APIKeySource, &out.APIKeySource
		*out = new(string)
		**out = **in
	}
	if in.BinaryMediaTypes != nil {
		in, out := &in.BinaryMediaTypes, &out.BinaryMediaTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = (*in).DeepCopy()
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableExecuteAPIEndpoint != nil {
		in, out := &in.DisableExecuteAPIEndpoint, &out.DisableExecuteAPIEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.EndpointConfiguration != nil {
		in, out := &in.EndpointConfiguration, &out.EndpointConfiguration
		*out = new(EndpointConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MinimumCompressionSize != nil {
		in, out := &in.MinimumCompressionSize, &out.MinimumCompressionSize
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.RootResourceID != nil {
		in, out := &in.RootResourceID, &out.RootResourceID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestAPI_SDK.
func (in *RestAPI_SDK) DeepCopy() *RestAPI_SDK {
	if in == nil {
		return nil
	}
	out := new(RestAPI_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SDKConfigurationProperty) DeepCopyInto(out *SDKConfigurationProperty) {
	*out = *in
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FriendlyName != nil {
		in, out := &in.FriendlyName, &out.FriendlyName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SDKConfigurationProperty.
func (in *SDKConfigurationProperty) DeepCopy() *SDKConfigurationProperty {
	if in == nil {
		return nil
	}
	out := new(SDKConfigurationProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageKey) DeepCopyInto(out *StageKey) {
	*out = *in
	if in.RestAPIID != nil {
		in, out := &in.RestAPIID, &out.RestAPIID
		*out = new(string)
		**out = **in
	}
	if in.StageName != nil {
		in, out := &in.StageName, &out.StageName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageKey.
func (in *StageKey) DeepCopy() *StageKey {
	if in == nil {
		return nil
	}
	out := new(StageKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	if in.InsecureSkipVerification != nil {
		in, out := &in.InsecureSkipVerification, &out.InsecureSkipVerification
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}
