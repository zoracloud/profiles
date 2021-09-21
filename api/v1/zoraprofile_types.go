/*
Copyright 2021.

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
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// plugin is for customize actions on different platform
type Plugin struct {
	metav1.TypeMeta `json:",inline"`
	Spec            *runtime.RawExtension `json:"spec,omitempty"`
}

type ProfileCondition struct {
	Type    string `json:"type,omitempty"`
	Status  string `json:"status,omitempty" description:"status of the condition, one of True, False, Unknown"`
	Message string `json:"message,omitempty"`
}

// ZoraProfileSpec defines the desired state of ZoraProfile
type ZoraProfileSpec struct {
	// The profile owner
	Owner   rbacv1.Subject `json:"owner,omitempty"`
	Plugins []Plugin       `json:"plugins,omitempty"`

	// Resourcequota that will be applied to target namespace
	ResourceQuotaSpec v1.ResourceQuotaSpec `json:"resourceQuotaSpec,omitempty"`
}

const (
	ProfileSucceed = "Successful"
	ProfileFailed  = "Failed"
	ProfileUnknown = "Unknown"
)

// ZoraProfileStatus defines the observed state of ZoraProfile
type ZoraProfileStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []ProfileCondition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ZoraProfile is the Schema for the zoraprofiles API
type ZoraProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZoraProfileSpec   `json:"spec,omitempty"`
	Status ZoraProfileStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ZoraProfileList contains a list of ZoraProfile
type ZoraProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoraProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZoraProfile{}, &ZoraProfileList{})
}
