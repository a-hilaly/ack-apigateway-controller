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

package rest_api

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ec2apitypes "github.com/aws-controllers-k8s/ec2-controller/apis/v1alpha1"
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"

	svcapitypes "github.com/aws-controllers-k8s/apigateway-controller/apis/v1alpha1"
)

// +kubebuilder:rbac:groups=ec2.services.k8s.aws,resources=vpcendpoints,verbs=get;list
// +kubebuilder:rbac:groups=ec2.services.k8s.aws,resources=vpcendpoints/status,verbs=get;list

// ClearResolvedReferences removes any reference values that were made
// concrete in the spec. It returns a copy of the input AWSResource which
// contains the original *Ref values, but none of their respective concrete
// values.
func (rm *resourceManager) ClearResolvedReferences(res acktypes.AWSResource) acktypes.AWSResource {
	ko := rm.concreteResource(res).ko.DeepCopy()

	if ko.Spec.EndpointConfiguration != nil {
		if len(ko.Spec.EndpointConfiguration.VPCEndpointRefs) > 0 {
			ko.Spec.EndpointConfiguration.VPCEndpointIDs = nil
		}
	}

	return &resource{ko}
}

// ResolveReferences finds if there are any Reference field(s) present
// inside AWSResource passed in the parameter and attempts to resolve those
// reference field(s) into their respective target field(s). It returns a
// copy of the input AWSResource with resolved reference(s), a boolean which
// is set to true if the resource contains any references (regardless of if
// they are resolved successfully) and an error if the passed AWSResource's
// reference field(s) could not be resolved.
func (rm *resourceManager) ResolveReferences(
	ctx context.Context,
	apiReader client.Reader,
	res acktypes.AWSResource,
) (acktypes.AWSResource, bool, error) {
	ko := rm.concreteResource(res).ko

	resourceHasReferences := false
	err := validateReferenceFields(ko)
	if fieldHasReferences, err := rm.resolveReferenceForEndpointConfiguration_VPCEndpointIDs(ctx, apiReader, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	return &resource{ko}, resourceHasReferences, err
}

// validateReferenceFields validates the reference field and corresponding
// identifier field.
func validateReferenceFields(ko *svcapitypes.RestAPI) error {

	if ko.Spec.EndpointConfiguration != nil {
		if len(ko.Spec.EndpointConfiguration.VPCEndpointRefs) > 0 && len(ko.Spec.EndpointConfiguration.VPCEndpointIDs) > 0 {
			return ackerr.ResourceReferenceAndIDNotSupportedFor("EndpointConfiguration.VPCEndpointIDs", "EndpointConfiguration.VPCEndpointRefs")
		}
	}
	return nil
}

// resolveReferenceForEndpointConfiguration_VPCEndpointIDs reads the resource referenced
// from EndpointConfiguration.VPCEndpointRefs field and sets the EndpointConfiguration.VPCEndpointIDs
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForEndpointConfiguration_VPCEndpointIDs(
	ctx context.Context,
	apiReader client.Reader,
	ko *svcapitypes.RestAPI,
) (hasReferences bool, err error) {
	if ko.Spec.EndpointConfiguration != nil {
		for _, f0iter := range ko.Spec.EndpointConfiguration.VPCEndpointRefs {
			if f0iter != nil && f0iter.From != nil {
				hasReferences = true
				arr := f0iter.From
				if arr.Name == nil || *arr.Name == "" {
					return hasReferences, fmt.Errorf("provided resource reference is nil or empty: EndpointConfiguration.VPCEndpointRefs")
				}
				namespace := ko.ObjectMeta.GetNamespace()
				if arr.Namespace != nil && *arr.Namespace != "" {
					namespace = *arr.Namespace
				}
				obj := &ec2apitypes.VPCEndpoint{}
				if err := getReferencedResourceState_VPCEndpoint(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
					return hasReferences, err
				}
				if ko.Spec.EndpointConfiguration.VPCEndpointIDs == nil {
					ko.Spec.EndpointConfiguration.VPCEndpointIDs = make([]*string, 0, 1)
				}
				ko.Spec.EndpointConfiguration.VPCEndpointIDs = append(ko.Spec.EndpointConfiguration.VPCEndpointIDs, (*string)(obj.Status.VPCEndpointID))
			}
		}
	}

	return hasReferences, nil
}

// getReferencedResourceState_VPCEndpoint looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_VPCEndpoint(
	ctx context.Context,
	apiReader client.Reader,
	obj *ec2apitypes.VPCEndpoint,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"VPCEndpoint",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"VPCEndpoint",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"VPCEndpoint",
			namespace, name)
	}
	if obj.Status.VPCEndpointID == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"VPCEndpoint",
			namespace, name,
			"Status.VPCEndpointID")
	}
	return nil
}
