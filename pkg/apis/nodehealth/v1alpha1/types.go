package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeHealth is a specification for a NodeHealth resource
type NodeHealth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeHealthSpec   `json:"spec"`
	Status NodeHealthStatus `json:"status"`
}

type Ping struct {
	PingFrequency *int32 `json:"pingFrequency"`
	PingTimes     *int32 `json:"pingTime"`
}

// NodeHealthSpec is the spec for a NodeHealth resource
type NodeHealthSpec struct {
	Replicas     *int32 `json:"replicas"`
	NodeName     string `json:"nodeName"`
	Cluster      string `json:"cluster"`
	CheckSetting Ping   `json:"ping"`
}

type NodeStatus struct {
	LastHealthStatus       string    `json:"lastHealthStatus"`
	LastCheckTimeStamp     time.Time `json:"lastCheckTimeStamp"`
	LastHealthyTimeStamp   time.Time `json:"lastHealthyTimeStamp"`
	LastUnHealthyTimeStamp time.Time `json:"lastUnHealthyTimeStamp"`
}

// NodeHealthStatus is the status for a NodeHealth resource
type NodeHealthStatus struct {
	AvailableReplicas int32        `json:"availableReplicas"`
	HealthStatus      []NodeStatus `json:"healthStatus"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeHealthList is a list of NodeHealth resources
type NodeHealthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NodeHealth `json:"items"`
}
