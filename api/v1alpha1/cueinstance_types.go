/*
Copyright 2022.

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
	"time"

	"github.com/fluxcd/pkg/apis/meta"
	"github.com/fluxcd/pkg/runtime/dependency"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type ValidationMode string

const (
	// IgnorePolicy will ignore validation errors
	IgnorePolicy ValidationMode = "Ignore"
	// AuditPolicy will ignore validation failures and generate an event
	AuditPolicy ValidationMode = "Audit"
	// DropPolicy will drop objects which are invalid but continue to reconcile valid objects
	DropPolicy ValidationMode = "Drop"
	// FailPolicy will fail the entire reconcile if any validation errors are encountered
	FailPolicy ValidationMode = "Fail"
)

const (
	CueInstanceKind           = "CueInstance"
	CueInstanceFinalizer      = "finalizers.fluxcd.io"
	MaxConditionMessageLength = 20000
	DisabledValue             = "disabled"
)

// CueInstanceSpec defines the desired state of CueInstance
type CueInstanceSpec struct {
	// The interval at which the instance will be reconciled.
	// +required
	Interval metav1.Duration `json:"interval"`

	// A reference to a Flux Source from which an artifact will be downloaded
	// and the CUE instance built.
	// +required
	SourceRef CrossNamespaceSourceReference `json:"sourceRef"`

	// The module root of the CUE instance.
	// +optional
	Root string `json:"root,omitempty"`

	// The path at which the CUE instance will be built from.
	// +optional
	Path string `json:"path,omitempty"`

	// The CUE package to use for the CUE instance. This is useful when applying
	// a CUE schema to plain yaml files.
	// +optional
	Package string `json:"package,omitempty"`

	// Tags that will be injected into the CUE instance.
	// +optional
	Tags []TagVar `json:"tags,omitempty"`

	// TagVars that will be available to the CUE instance.
	// +optional
	TagVars []TagVar `json:"tagVars,omitempty"`

	// The CUE expression(s) to execute.
	// +optional
	Exprs []string `json:"expressions,omitempty"`

	// Dependencies that must be ready before the CUE instance is reconciled.
	// +optional
	DependsOn []dependency.CrossNamespaceDependencyReference `json:"dependsOn,omitempty"`

	// Prune enables garbage collection.
	// +required
	Prune bool `json:"prune"`

	// The interval at which to retry a previously failed reconciliation.
	// When not specified, the controller uses the CueInstanceSpec.Interval
	// value to retry failures.
	// +optional
	RetryInterval *metav1.Duration `json:"retryInterval,omitempty"`

	// Timeout for validation, apply and health checking operations.
	// Defaults to 'Interval' duration.
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`

	// This flag tells the controller to suspend subsequent cue executions,
	// it does not apply to already started executions. Defaults to false.
	// +optional
	Suspend bool `json:"suspend,omitempty"`

	// The name of the Kubernetes service account to impersonate
	// when reconciling this CueInstance.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// The KubeConfig for reconciling the CueInstance on a remote cluster.
	// When specified, KubeConfig takes precedence over ServiceAccountName.
	// +optional
	KubeConfig *KubeConfig `json:"kubeConfig,omitempty"`

	// Force instructs the controller to recreate resources
	// when patching fails due to an immutable field change.
	// +kubebuilder:default:=false
	// +optional
	Force bool `json:"force,omitempty"`

	// TODO(maybe): this could be an array of validations
	// in which case the policy may need to apply to all resources
	// would allow for greater flexibility
	// +optional
	Validate *Validation `json:"validate,omitempty"`
}

// TagVar is a tag variable with a required name and optional value
type TagVar struct {
	// +required
	Name string `json:"name"`

	// +optional
	Value string `json:"value,omitempty"`
}

type Validation struct {
	// +kubebuilder:default:="Audit"
	// +optional
	Mode ValidationMode `json:"mode,omitempty"`

	// +required
	Schema string `json:"schema"`

	// +kubebuilder:default:="yaml"
	// +optional
	Type string `json:"type,omitempty"`
}

// GetTimeout returns the timeout
func (in CueInstance) GetTimeout() time.Duration {
	duration := in.Spec.Interval.Duration - 30*time.Second
	if in.Spec.Timeout != nil {
		duration = in.Spec.Timeout.Duration
	}
	if duration < 30*time.Second {
		return 30 * time.Second
	}
	return duration
}

// KubeConfig references a Kubernetes secret that contains a kubeconfig file.
type KubeConfig struct {
	// SecretRef holds the name to a secret that contains a 'value' key with
	// the kubeconfig file as the value. It must be in the same namespace as
	// the CueInstance.
	// It is recommended that the kubeconfig is self-contained, and the secret
	// is regularly updated if credentials such as a cloud-access-token expire.
	// Cloud specific `cmd-path` auth helpers will not function without adding
	// binaries and credentials to the Pod that is responsible for reconciling
	// the CueInstance.
	// +required
	SecretRef meta.LocalObjectReference `json:"secretRef,omitempty"`
}

// GetRetryInterval returns the retry interval
func (in CueInstance) GetRetryInterval() time.Duration {
	if in.Spec.RetryInterval != nil {
		return in.Spec.RetryInterval.Duration
	}
	return in.Spec.Interval.Duration
}

// GetDependsOn returns the list of dependencies across-namespaces.
func (in CueInstance) GetDependsOn() (types.NamespacedName, []dependency.CrossNamespaceDependencyReference) {
	return types.NamespacedName{
		Namespace: in.Namespace,
		Name:      in.Name,
	}, in.Spec.DependsOn
}

// GetStatusConditions returns a pointer to the Status.Conditions slice.
func (in *CueInstance) GetStatusConditions() *[]metav1.Condition {
	return &in.Status.Conditions
}

// CueInstanceProgressing resets the conditions of the given CueInstance to a single
// ReadyCondition with status ConditionUnknown.
func CueInstanceProgressing(k CueInstance, message string) CueInstance {
	meta.SetResourceCondition(&k, meta.ReadyCondition, metav1.ConditionUnknown, meta.ProgressingReason, message)
	return k
}

// SetCueInstanceReadiness sets the ReadyCondition, ObservedGeneration, and LastAttemptedRevision, on the CueInstance.
func SetCueInstanceReadiness(k *CueInstance, status metav1.ConditionStatus, reason, message string, revision string) {
	meta.SetResourceCondition(k, meta.ReadyCondition, status, reason, trimString(message, MaxConditionMessageLength))
	k.Status.ObservedGeneration = k.Generation
	k.Status.LastAttemptedRevision = revision
}

// CueInstanceNotReady registers a failed apply attempt of the given CueInstance.
func CueInstanceNotReady(k CueInstance, revision, reason, message string) CueInstance {
	SetCueInstanceReadiness(&k, metav1.ConditionFalse, reason, trimString(message, MaxConditionMessageLength), revision)
	if revision != "" {
		k.Status.LastAttemptedRevision = revision
	}
	return k
}

// CueInstanceNotReadyInventory registers a failed apply attempt of the given CueInstance.
func CueInstanceNotReadyInventory(k CueInstance, inventory *ResourceInventory, revision, reason, message string) CueInstance {
	SetCueInstanceReadiness(&k, metav1.ConditionFalse, reason, trimString(message, MaxConditionMessageLength), revision)
	if revision != "" {
		k.Status.LastAttemptedRevision = revision
	}
	k.Status.Inventory = inventory
	return k
}

// CueInstanceReadyInventory registers a successful apply attempt of the given CueInstance.
func CueInstanceReadyInventory(k CueInstance, inventory *ResourceInventory, revision, reason, message string) CueInstance {
	SetCueInstanceReadiness(&k, metav1.ConditionTrue, reason, trimString(message, MaxConditionMessageLength), revision)
	k.Status.Inventory = inventory
	k.Status.LastAppliedRevision = revision
	return k
}

// CueInstanceStatus defines the observed state of CueInstance
type CueInstanceStatus struct {
	meta.ReconcileRequestStatus `json:",inline"`

	// ObservedGeneration is the last reconciled generation.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The last successfully applied revision.
	// The revision format for Git sources is <branch|tag>/<commit-sha>.
	// +optional
	LastAppliedRevision string `json:"lastAppliedRevision,omitempty"`

	// LastAttemptedRevision is the revision of the last reconciliation attempt.
	// +optional
	LastAttemptedRevision string `json:"lastAttemptedRevision,omitempty"`

	// Inventory contains the list of Kubernetes resource object references that have been successfully applied.
	// +optional
	Inventory *ResourceInventory `json:"inventory,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CueInstance is the Schema for the cueinstances API
type CueInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CueInstanceSpec   `json:"spec,omitempty"`
	Status CueInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CueInstanceList contains a list of CueInstance
type CueInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CueInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CueInstance{}, &CueInstanceList{})
}

func trimString(str string, limit int) string {
	if len(str) <= limit {
		return str
	}

	return str[0:limit] + "..."
}
