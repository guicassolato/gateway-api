/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apisv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// GatewayInfrastructureApplyConfiguration represents a declarative configuration of the GatewayInfrastructure type for use
// with apply.
type GatewayInfrastructureApplyConfiguration struct {
	Labels        map[apisv1.LabelKey]apisv1.LabelValue           `json:"labels,omitempty"`
	Annotations   map[apisv1.AnnotationKey]apisv1.AnnotationValue `json:"annotations,omitempty"`
	ParametersRef *LocalParametersReferenceApplyConfiguration     `json:"parametersRef,omitempty"`
}

// GatewayInfrastructureApplyConfiguration constructs a declarative configuration of the GatewayInfrastructure type for use with
// apply.
func GatewayInfrastructure() *GatewayInfrastructureApplyConfiguration {
	return &GatewayInfrastructureApplyConfiguration{}
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *GatewayInfrastructureApplyConfiguration) WithLabels(entries map[apisv1.LabelKey]apisv1.LabelValue) *GatewayInfrastructureApplyConfiguration {
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[apisv1.LabelKey]apisv1.LabelValue, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *GatewayInfrastructureApplyConfiguration) WithAnnotations(entries map[apisv1.AnnotationKey]apisv1.AnnotationValue) *GatewayInfrastructureApplyConfiguration {
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[apisv1.AnnotationKey]apisv1.AnnotationValue, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithParametersRef sets the ParametersRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ParametersRef field is set to the value of the last call.
func (b *GatewayInfrastructureApplyConfiguration) WithParametersRef(value *LocalParametersReferenceApplyConfiguration) *GatewayInfrastructureApplyConfiguration {
	b.ParametersRef = value
	return b
}
