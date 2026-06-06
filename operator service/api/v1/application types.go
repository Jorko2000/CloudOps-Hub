package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ApplicationSpec struct {
	Image    string `json:"image"`
	Replicas int32  `json:"replicas"`
}

type ApplicationStatus struct {
	Status string `json:"status,omitempty"`
}

type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}
