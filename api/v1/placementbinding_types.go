// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Subject reference
type Subject struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	APIGroup string `json:"apiGroup"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Kind string `json:"kind"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
}

// PlacementBindingStatus defines the observed state of PlacementBinding
type PlacementBindingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

//+kubebuilder:object:root=true

// PlacementBinding is the Schema for the placementbindings API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=placementbindings,scope=Namespaced
// +kubebuilder:resource:path=placementbindings,shortName=pb
type PlacementBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:Required
	PlacementRef Subject `json:"placementRef"`
	// +kubebuilder:validation:Required
	Subjects []Subject              `json:"subjects"`
	Status   PlacementBindingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PlacementBindingList contains a list of PlacementBinding
type PlacementBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlacementBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PlacementBinding{}, &PlacementBindingList{})
}
