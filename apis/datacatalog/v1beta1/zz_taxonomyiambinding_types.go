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

type TaxonomyIAMBindingConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TaxonomyIAMBindingConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TaxonomyIAMBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type TaxonomyIAMBindingInitParameters struct {
	Condition *TaxonomyIAMBindingConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Taxonomy *string `json:"taxonomy,omitempty" tf:"taxonomy,omitempty"`
}

type TaxonomyIAMBindingObservation struct {
	Condition *TaxonomyIAMBindingConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Taxonomy *string `json:"taxonomy,omitempty" tf:"taxonomy,omitempty"`
}

type TaxonomyIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition *TaxonomyIAMBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Taxonomy *string `json:"taxonomy,omitempty" tf:"taxonomy,omitempty"`
}

// TaxonomyIAMBindingSpec defines the desired state of TaxonomyIAMBinding
type TaxonomyIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TaxonomyIAMBindingParameters `json:"forProvider"`
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
	InitProvider TaxonomyIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// TaxonomyIAMBindingStatus defines the observed state of TaxonomyIAMBinding.
type TaxonomyIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TaxonomyIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TaxonomyIAMBinding is the Schema for the TaxonomyIAMBindings API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TaxonomyIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.taxonomy) || (has(self.initProvider) && has(self.initProvider.taxonomy))",message="spec.forProvider.taxonomy is a required parameter"
	Spec   TaxonomyIAMBindingSpec   `json:"spec"`
	Status TaxonomyIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TaxonomyIAMBindingList contains a list of TaxonomyIAMBindings
type TaxonomyIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TaxonomyIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	TaxonomyIAMBinding_Kind             = "TaxonomyIAMBinding"
	TaxonomyIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TaxonomyIAMBinding_Kind}.String()
	TaxonomyIAMBinding_KindAPIVersion   = TaxonomyIAMBinding_Kind + "." + CRDGroupVersion.String()
	TaxonomyIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(TaxonomyIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&TaxonomyIAMBinding{}, &TaxonomyIAMBindingList{})
}
