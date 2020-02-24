/*
Copyright 2020 ShotaKitazawa.

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
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ManyTimesJobSpec defines the desired state of ManyTimesJob
type ManyTimesJobSpec struct {
	// +kubebuilder:validation:Default=false

	// Default false, true when job finished.
	// This field is requirement for apply to same name Job resource.
	// +optional
	Done bool `json:"done,omitempty"`

	// Specifies the job that will be created when executing a CronJob.
	JobTemplate batchv1beta1.JobTemplateSpec `json:"jobTemplate"`

	// +kubebuilder:validation:Minimum=0

	// The number of successful finished jobs to retain.
	// This is a pointer to distinguish between explicit zero and not specified.
	// +optional
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// +kubebuilder:validation:Minimum=0

	// The number of failed finished jobs to retain.
	// This is a pointer to distinguish between explicit zero and not specified.
	// +optional
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}

// ManyTimesJobStatus defines the observed state of ManyTimesJob
type ManyTimesJobStatus struct {

	// A list of pointers to currently running jobs.
	// +optional
	Active []corev1.ObjectReference `json:"active,omitempty"`
}

// +kubebuilder:object:root=true

// ManyTimesJob is the Schema for the manytimesjobs API
type ManyTimesJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManyTimesJobSpec   `json:"spec,omitempty"`
	Status ManyTimesJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManyTimesJobList contains a list of ManyTimesJob
type ManyTimesJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManyTimesJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManyTimesJob{}, &ManyTimesJobList{})
}
