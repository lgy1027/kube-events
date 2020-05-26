/*
Copyright 2020 The KubeSphere Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubeEventsRuleSpec defines the desired state of KubeEventsRule
type KubeEventsRuleSpec struct {
	Rules []Rule `json:"rules,omitempty"`
}

// KubeEventsRuleStatus defines the observed state of KubeEventsRule
type KubeEventsRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// KubeEventsRule is the Schema for the kubeeventsrules API
type KubeEventsRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeEventsRuleSpec   `json:"spec"`
	Status KubeEventsRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeEventsRuleList contains a list of KubeEventsRule
type KubeEventsRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeEventsRule `json:"items"`
}

// Rule describes a notification or alert rule
type Rule struct {
	// Name is simple name of rule
	Name string `json:"name,omitempty"`
	// Condition is a string similar with the where part of sql (please use double quotation to mark a string).
	// For example: `event.type="Warning" and event.involvedObject.kind="Pod" and event.reason="FailedMount"`
	Condition string `json:"condition,omitempty"`
	// Labels
	Labels map[string]string `json:"labels,omitempty"`
	// Values of Annotations can use format string with the fields of the event.
	// For example: `{"message": "%event.message"}`
	Annotations map[string]string `json:"annotations,omitempty"`
	// Enable is whether to enable the rule
	Enable bool `json:"enable,omitempty"`
	// Type represents that the rule is for notification or alert.
	// Available values are `notification` and `alert`
	Type RuleType `json:"type,omitempty"`
}

type RuleType string

const (
	// RuleTypeNotification represents that the rule will used to generate notifications
	// based on the original event objects.
	RuleTypeNotification = "notification"
	// RuleTypeAlert represents that the rule will be used to generate alert messages
	// that conform to the alertmanager protocol.
	RuleTypeAlert = "alert"
)

func init() {
	SchemeBuilder.Register(&KubeEventsRule{}, &KubeEventsRuleList{})
}
