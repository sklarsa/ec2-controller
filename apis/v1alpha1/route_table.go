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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RouteTableSpec defines the desired state of RouteTable.
//
// Describes a route table.
type RouteTableSpec struct {
	// The tags to assign to the route table.
	TagSpecifications []*TagSpecification `json:"tagSpecifications,omitempty"`
	// The ID of the VPC.
	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcID"`
}

// RouteTableStatus defines the observed state of RouteTable
type RouteTableStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The associations between the route table and one or more subnets or a gateway.
	// +kubebuilder:validation:Optional
	Associations []*RouteTableAssociation `json:"associations,omitempty"`
	// The ID of the AWS account that owns the route table.
	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerID,omitempty"`
	// Any virtual private gateway (VGW) propagating routes.
	// +kubebuilder:validation:Optional
	PropagatingVGWs []*PropagatingVGW `json:"propagatingVGWs,omitempty"`
	// The ID of the route table.
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableID,omitempty"`
	// The routes in the route table.
	// +kubebuilder:validation:Optional
	Routes []*Route `json:"routes,omitempty"`
	// Any tags assigned to the route table.
	// +kubebuilder:validation:Optional
	Tags []*Tag `json:"tags,omitempty"`
}

// RouteTable is the Schema for the RouteTables API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec,omitempty"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

// RouteTableList contains a list of RouteTable
// +kubebuilder:object:root=true
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
