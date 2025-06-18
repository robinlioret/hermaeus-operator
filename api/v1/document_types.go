/*
Copyright 2025.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Destination specification.
type Destination struct {
	// Reference to the git repository.
	GitRepositoryRef corev1.ObjectReference `json:"gitRepositoryRef"`

	// Path in the git repository.
	Path string `json:"path"`
}

// DocumentSpec defines the desired state of Document.
type DocumentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Where the document is supposed to be stored. Ignored if empty.
	// +optional
	Destination Destination `json:"destination,omitempty"`

	// Textual content of the document file.
	Content string `json:"content"`
}

// DocumentStatus defines the observed state of Document.
type DocumentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Document is the Schema for the documents API.
type Document struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocumentSpec   `json:"spec,omitempty"`
	Status DocumentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentList contains a list of Document.
type DocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Document `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Document{}, &DocumentList{})
}
