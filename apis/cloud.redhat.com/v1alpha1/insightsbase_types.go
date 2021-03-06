/*
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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=none;operator
type WebProvider string

type WebConfig struct {
	Port      int32       `json:"port,omitempty"`
	ApiPrefix string      `json:"apiPrefix,omitempty"`
	Provider  WebProvider `json:"provider"`
}

// +kubebuilder:validation:Enum=none;operator
type MetricsProvider string

type MetricsConfig struct {
	Port     int32           `json:"port,omitempty"`
	Path     string          `json:"path,omitempty"`
	Provider MetricsProvider `json:"provider"`
}

// TODO: Other potential provider: saas

// +kubebuilder:validation:Enum=operator;app-interface;local
type KafkaProvider string

type KafkaConfig struct {
	ClusterName string        `json:"clusterName"`
	Namespace   string        `json:"namespace"`
	Provider    KafkaProvider `json:"provider"`
	Suffix      string        `json:"suffix,omitempty"`
}

// TODO: Other potential providers: RDS and Operator (e.g. CrunchyDB)

// +kubebuilder:validation:Enum=app-interface;local
type DatabaseProvider string

type DatabaseConfig struct {
	Provider DatabaseProvider `json:"provider"`
	Image    string           `json:"image"`
}

// TODO: Other potential providers: splunk, kafka

type LoggingProviders []string

type LoggingConfig struct {
	Providers LoggingProviders `json:"providers"`
}

// TODO: Other potential provider: ceph, S3

// +kubebuilder:validation:Enum=minio;app-interface
type ObjectStoreProvider string

type ObjectStoreConfig struct {
	Provider ObjectStoreProvider `json:"provider"`
	Suffix   string              `json:"suffix,omitempty"`
}

// +kubebuilder:validation:Enum=redis;app-interface
type InMemoryProvider string

type InMemoryDBConfig struct {
	Provider InMemoryProvider `json:"provider"`
}

// InsightsBaseSpec defines the desired state of InsightsBase
type InsightsBaseSpec struct {
	Web         WebConfig         `json:"web,omitempty"`
	Metrics     MetricsConfig     `json:"metrics,omitempty"`
	Kafka       KafkaConfig       `json:"kafka"`
	Database    DatabaseConfig    `json:"db,omitempty"`
	Logging     LoggingConfig     `json:"logging"`
	ObjectStore ObjectStoreConfig `json:"objectStore"`
	InMemoryDB  InMemoryDBConfig  `json:"inMemoryDb"`
}

type MinioStatus struct {
	Credentials core.SecretReference `json:"credentials"`
	Endpoint    string               `json:"endpoint"`
}

type ObjectStoreStatus struct {
	Minio   MinioStatus `json:"minio,omitempty"`
	Buckets []string    `json:"buckets"`
}

// InsightsBaseStatus defines the observed state of InsightsBase
type InsightsBaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ObjectStore ObjectStoreStatus `json:"objectStore"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// InsightsBase is the Schema for the insightsbases API
type InsightsBase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InsightsBaseSpec   `json:"spec,omitempty"`
	Status InsightsBaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InsightsBaseList contains a list of InsightsBase
type InsightsBaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InsightsBase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InsightsBase{}, &InsightsBaseList{})
}

func (i *InsightsBase) GetLabels() map[string]string {
	return map[string]string{"app": i.ObjectMeta.Name}
}

func (i *InsightsBase) MakeOwnerReference() metav1.OwnerReference {
	return metav1.OwnerReference{
		APIVersion: i.APIVersion,
		Kind:       i.Kind,
		Name:       i.ObjectMeta.Name,
		UID:        i.ObjectMeta.UID,
	}
}
