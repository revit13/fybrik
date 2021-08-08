// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AssetDetails is a list of assets used in the fybrikapplication. In addition to assets declared in
// fybrikapplication, AssetDetails list also contains assets that are allocated by the control-plane
// in order to serve fybrikapplication
type AssetDetails struct {

	// AssetID identifies the asset to be used for accessing the data
	// +required
	AssetID string `json:"assetID"`

	// AdvertisedAssetID links this asset to asset from fybrikapplication and is used by user facing services
	// +optional
	AdvertisedAssetID string `json:"advertisedAssetID,omitempty"`

	// Cluster name where the asset resides
	// +required
	ClusterName string `json:"cluster"`

	// +required
	DataStore DataStore `json:"assetDetails"`
}

type Service struct {
	// AssetID which this flow process
	// +required
	AssetID string `json:"assetID"`
	// Interface indicates the protocol and format expected by the data user
	// +required
	Interface InterfaceDetails `json:"interface"`
}

type StepAPI struct {
	// +required
	Service Service `json:"service"`
}

// StepSource is the source of this step: it could be assetID
// or the sink of another step
type StepSource struct {
	// AssetID identifies the source asset of this step
	// +required
	AssetID string `json:"assetID"`

	// API indicates to the application how to access the module for the given capability
	// +optional
	API *StepAPI `json:"api,omitempty"`
}

// StepSink holds information to where the target data will be written:
// it could be assetID of an asset specified in fybrikapplication or an an asset created
// by fybrik control-plane
type StepSink struct {
	// AssetID identifies the target asset of this step
	// +required
	AssetID string `json:"assetID"`
}

// StepParameters contains the parameters to the module
// that is deployed in this step
type StepParameters struct {
	// +optional
	Source *StepSource `json:"source,omitempty"`

	// +optional
	Sink *StepSink `json:"sink,omitempty"`

	// API indicates to the application how to access the module for the given capability
	// +optional
	API *StepAPI `json:"api,omitempty"`

	// Actions are the data transformations that the module supports
	// +optional
	Actions []SupportedAction `json:"action,omitempty"`
}

// Step contains details of a data flow step
type Step struct {
	// Name of the step
	// +required
	Name string `json:"name"`

	// Name of the cluster this step is executed on
	// +required
	Cluster string `json:"cluster"`

	// The template name of this step.
	// The full details of the template can be extracted from Plotter.spec.templates
	// list field.
	// +required
	Template string `json:"template"`

	// +optional
	Parameters *StepParameters `json:"parameters,omitempty"`
}

// SequentialSteps defines a series of sequential data flow steps
// +kubebuilder:validation:Type=array
type SequentialSteps struct {
	// +required
	Steps []Step `json:"-" protobuf:"bytes,1,rep,name=steps"`
}

// Subflows is a list of data flows which are originated from the same data asset
// but are triggered differently (e.g., one upon init
// trigger and one upon workload trigger)
type SubFlow struct {
	// Name of the SubFlow
	// +required
	Name string `json:"name"`

	// Type of the flow (e.g. read)
	// +required
	FlowType string `json:"flowType"`

	// Triggers
	// +required
	Triggers []CapabilityScope `json:"triggers"`

	// Steps defines a series of sequential/parallel data flow steps
	// +required
	Steps []SequentialSteps `json:"steps" protobuf:"bytes,11,opt,name=steps"`
}

// Flows is the list of data flows driven from fybrikapplication:
// Each element in the list holds the flow of the data requested in fybrikapplication.
type Flow struct {
	// Name of the flow
	// +required
	Name string `json:"name"`

	// Type of the flow (e.g. read)
	// +required
	FlowType string `json:"flowType"`

	// AssetID of this data flow
	// +required
	AssetID string `json:"assetID"`

	// +required
	SubFlows []SubFlow `json:"subFlows"`
}

// Basic info needed for chart deployment
type ChartBasicInfo struct {
	// Name of the chart
	// +required
	Name string `json:"name"`
}

// ModuleBasicInfo is a subset of M4DModule Custom Resource.  It contains partial information
// to instantiate resource of type M4DModule. The complete information to instantiate the module
// can be constructed from other properties in the plotter e.g., the flows
type ModuleBasicInfo struct {
	// Name of the template
	// +required
	Name string `json:"name"`

	// Kind of k8s resource
	// +required
	Kind string `json:"kind"`

	// +required
	Chart ChartBasicInfo `json:"chart"`
}

// Template contains basic information about the required modules to serve the fybrikapplication
type Template struct {
	// Template name
	Name string `json:"name"`
	// Modules is a list of dependent modules. e.g., if a plugin module is used
	// then the service module is used in should appear in the same modules list.
	// +required
	Modules []ModuleBasicInfo `json:"modules"`
}

// PlotterSpec defines the desired state of Plotter, which is applied in a multi-clustered environment. Plotter installs the runtime environment
// (as blueprints running on remote clusters) which provides the Data Scientist's application with secure and governed access to the data requested in the
// FybrikApplication.
type PlotterSpec struct {

	// Selector enables to connect the resource to the application
	// Application labels should match the labels in the selector.
	// For some flows the selector may not be used.
	// +optional
	Selector Selector `json:"appSelector,omitempty"`

	// +required
	Assets []AssetDetails `json:"assets"`

	// +required
	Flows []Flow `json:"flows"`

	// +required
	Templates []Template `json:"templates"`
}

// SubFlowsStatus holds status information about a subFlow
type SubFlowsStatus struct {
	// SubFlow name
	// +required
	Name string `json:"name"`

	// ObservedState includes information about this subflow
	// It includes readiness and error indications, as well as user instructions
	// +optional
	ObservedState ObservedState `json:"status,omitempty"`
}

// AssetStatus includes information to be reported back to the FybrikApplication resource
// It holds the status and access information for each asset in fybrikapplication
type AssetStatus struct {
	// Asset name
	// +required
	Name string `json:"name"`

	// Endpoint
	// +required
	Endpoint string `json:"endpoint"`

	// Port
	// +required
	Port int32 `json:"port"`

	// ObservedState includes information to be reported back to the FybrikApplication resource
	// It includes readiness and error indications, as well as user instructions
	// +optional
	ObservedState ObservedState `json:"status,omitempty"`

	// +optional
	Errors []string `json:"errors,omitempty,omitempty"`
}

// FlowStatus includes information to be reported back to the FybrikApplication resource
// It holds the status per data flow
type FlowStatus struct {
	// Template name
	// +required
	Name string `json:"name"`

	// ObservedState includes information about the current flow
	// It includes readiness and error indications, as well as user instructions
	// +optional
	ObservedState ObservedState `json:"status,omitempty"`

	// +required
	SubFlows []SubFlowsStatus `json:"subFlows"`
}

// PlotterStatus defines the observed state of Plotter
// This includes readiness, error message, and indicators received from blueprint
// resources owned by the Plotter for cleanup and status monitoring
type PlotterStatus struct {
	// ObservedState includes information to be reported back to the FybrikApplication resource
	// It includes readiness and error indications, as well as user instructions
	// +optional
	ObservedState ObservedState `json:"observedState,omitempty"`

	// ObservedGeneration is taken from the Plotter metadata.  This is used to determine during reconcile
	// whether reconcile was called because the desired state changed, or whether status of the allocated blueprints should be checked.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// +required
	Flows []FlowStatus `json:"flows"`

	// +required
	Assets []AssetStatus `json:"assets"`

	// Conditions represent the possible error and failure conditions
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.observedState.ready`
// +kubebuilder:printcolumn:name="ReadySince",type=string,JSONPath=`.status.readyTimestamp`

// Plotter is the Schema for the plotters API
type Plotter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlotterSpec   `json:"spec,omitempty"`
	Status PlotterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlotterList contains a list of Plotter resources
type PlotterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Plotter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Plotter{}, &PlotterList{})
}
