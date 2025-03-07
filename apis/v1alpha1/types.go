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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// API stage name of the associated API stage in a usage plan.
type APIStage struct {
	APIID *string `json:"apiID,omitempty"`
	Stage *string `json:"stage,omitempty"`
}

// Access log settings, including the access log format and access log destination
// ARN.
type AccessLogSettings struct {
	DestinationARN *string `json:"destinationARN,omitempty"`
	Format         *string `json:"format,omitempty"`
}

// Configuration settings of a canary deployment.
type CanarySettings struct {
	DeploymentID           *string            `json:"deploymentID,omitempty"`
	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty"`
	UseStageCache          *bool              `json:"useStageCache,omitempty"`
}

// The input configuration for a canary deployment.
type DeploymentCanarySettings struct {
	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty"`
	UseStageCache          *bool              `json:"useStageCache,omitempty"`
}

// Specifies the target API entity to which the documentation applies.
type DocumentationPartLocation struct {
	Method *string `json:"method,omitempty"`
	Name   *string `json:"name,omitempty"`
	Path   *string `json:"path,omitempty"`
}

// The endpoint configuration to indicate the types of endpoints an API (RestApi)
// or its custom domain name (DomainName) has.
type EndpointConfiguration struct {
	Types          []*string `json:"types,omitempty"`
	VPCEndpointIDs []*string `json:"vpcEndpointIDs,omitempty"`
	// Reference field for VPCEndpointIDs
	VPCEndpointRefs []*ackv1alpha1.AWSResourceReferenceWrapper `json:"vpcEndpointRefs,omitempty"`
}

// Specifies the method setting properties.
type MethodSetting struct {
	CacheDataEncrypted                  *bool   `json:"cacheDataEncrypted,omitempty"`
	CachingEnabled                      *bool   `json:"cachingEnabled,omitempty"`
	DataTraceEnabled                    *bool   `json:"dataTraceEnabled,omitempty"`
	LoggingLevel                        *string `json:"loggingLevel,omitempty"`
	MetricsEnabled                      *bool   `json:"metricsEnabled,omitempty"`
	RequireAuthorizationForCacheControl *bool   `json:"requireAuthorizationForCacheControl,omitempty"`
}

// Represents a summary of a Method resource, given a particular date and time.
type MethodSnapshot struct {
	APIKeyRequired    *bool   `json:"apiKeyRequired,omitempty"`
	AuthorizationType *string `json:"authorizationType,omitempty"`
}

// The mutual TLS authentication configuration for a custom domain name. If
// specified, API Gateway performs two-way authentication between the client
// and the server. Clients must present a trusted certificate to access your
// API.
type MutualTLSAuthentication struct {
	TruststoreURI      *string   `json:"truststoreURI,omitempty"`
	TruststoreVersion  *string   `json:"truststoreVersion,omitempty"`
	TruststoreWarnings []*string `json:"truststoreWarnings,omitempty"`
}

// The mutual TLS authentication configuration for a custom domain name. If
// specified, API Gateway performs two-way authentication between the client
// and the server. Clients must present a trusted certificate to access your
// API.
type MutualTLSAuthenticationInput struct {
	TruststoreURI     *string `json:"truststoreURI,omitempty"`
	TruststoreVersion *string `json:"truststoreVersion,omitempty"`
}

// For more information about supported patch operations, see Patch Operations
// (https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).
type PatchOperation struct {
	From  *string `json:"from,omitempty"`
	Op    *string `json:"op,omitempty"`
	Path  *string `json:"path,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Represents a REST API.
type RestAPI_SDK struct {
	APIKeySource              *string      `json:"apiKeySource,omitempty"`
	BinaryMediaTypes          []*string    `json:"binaryMediaTypes,omitempty"`
	CreatedDate               *metav1.Time `json:"createdDate,omitempty"`
	Description               *string      `json:"description,omitempty"`
	DisableExecuteAPIEndpoint *bool        `json:"disableExecuteAPIEndpoint,omitempty"`
	// The endpoint configuration to indicate the types of endpoints an API (RestApi)
	// or its custom domain name (DomainName) has.
	EndpointConfiguration  *EndpointConfiguration `json:"endpointConfiguration,omitempty"`
	ID                     *string                `json:"id,omitempty"`
	MinimumCompressionSize *int64                 `json:"minimumCompressionSize,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	Policy                 *string                `json:"policy,omitempty"`
	RootResourceID         *string                `json:"rootResourceID,omitempty"`
	Tags                   map[string]*string     `json:"tags,omitempty"`
	Version                *string                `json:"version,omitempty"`
	Warnings               []*string              `json:"warnings,omitempty"`
}

// A configuration property of an SDK type.
type SDKConfigurationProperty struct {
	DefaultValue *string `json:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty"`
	FriendlyName *string `json:"friendlyName,omitempty"`
	Name         *string `json:"name,omitempty"`
	Required     *bool   `json:"required,omitempty"`
}

// A reference to a unique stage identified in the format {restApiId}/{stage}.
type StageKey struct {
	RestAPIID *string `json:"restAPIID,omitempty"`
	StageName *string `json:"stageName,omitempty"`
}

// Specifies the TLS configuration for an integration.
type TLSConfig struct {
	InsecureSkipVerification *bool `json:"insecureSkipVerification,omitempty"`
}
