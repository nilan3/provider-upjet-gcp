// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PolicyTagIAMMemberConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type PolicyTagIAMMemberConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type PolicyTagIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type PolicyTagIAMMemberInitParameters struct {
	Condition *PolicyTagIAMMemberConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	PolicyTag *string `json:"policyTag,omitempty" tf:"policy_tag,omitempty"`
}

type PolicyTagIAMMemberObservation struct {
	Condition *PolicyTagIAMMemberConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	PolicyTag *string `json:"policyTag,omitempty" tf:"policy_tag,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type PolicyTagIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition *PolicyTagIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyTag *string `json:"policyTag,omitempty" tf:"policy_tag,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// PolicyTagIAMMemberSpec defines the desired state of PolicyTagIAMMember
type PolicyTagIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyTagIAMMemberParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PolicyTagIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// PolicyTagIAMMemberStatus defines the observed state of PolicyTagIAMMember.
type PolicyTagIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyTagIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PolicyTagIAMMember is the Schema for the PolicyTagIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type PolicyTagIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyTag) || (has(self.initProvider) && has(self.initProvider.policyTag))",message="spec.forProvider.policyTag is a required parameter"
	Spec   PolicyTagIAMMemberSpec   `json:"spec"`
	Status PolicyTagIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyTagIAMMemberList contains a list of PolicyTagIAMMembers
type PolicyTagIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyTagIAMMember `json:"items"`
}

// Repository type metadata.
var (
	PolicyTagIAMMember_Kind             = "PolicyTagIAMMember"
	PolicyTagIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyTagIAMMember_Kind}.String()
	PolicyTagIAMMember_KindAPIVersion   = PolicyTagIAMMember_Kind + "." + CRDGroupVersion.String()
	PolicyTagIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(PolicyTagIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyTagIAMMember{}, &PolicyTagIAMMemberList{})
}
