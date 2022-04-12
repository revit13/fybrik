//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"fybrik.io/fybrik/pkg/model/datacatalog"
	"fybrik.io/fybrik/pkg/model/taxonomy"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDetails) DeepCopyInto(out *ApplicationDetails) {
	*out = *in
	in.WorkloadSelector.DeepCopyInto(&out.WorkloadSelector)
	in.Context.DeepCopyInto(&out.Context)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDetails.
func (in *ApplicationDetails) DeepCopy() *ApplicationDetails {
	if in == nil {
		return nil
	}
	out := new(ApplicationDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetContext) DeepCopyInto(out *AssetContext) {
	*out = *in
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]*DataStore, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DataStore)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]taxonomy.Action, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetContext.
func (in *AssetContext) DeepCopy() *AssetContext {
	if in == nil {
		return nil
	}
	out := new(AssetContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetDetails) DeepCopyInto(out *AssetDetails) {
	*out = *in
	in.DataStore.DeepCopyInto(&out.DataStore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetDetails.
func (in *AssetDetails) DeepCopy() *AssetDetails {
	if in == nil {
		return nil
	}
	out := new(AssetDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetState) DeepCopyInto(out *AssetState) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	in.Endpoint.DeepCopyInto(&out.Endpoint)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetState.
func (in *AssetState) DeepCopy() *AssetState {
	if in == nil {
		return nil
	}
	out := new(AssetState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Blueprint) DeepCopyInto(out *Blueprint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Blueprint.
func (in *Blueprint) DeepCopy() *Blueprint {
	if in == nil {
		return nil
	}
	out := new(Blueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Blueprint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintList) DeepCopyInto(out *BlueprintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Blueprint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintList.
func (in *BlueprintList) DeepCopy() *BlueprintList {
	if in == nil {
		return nil
	}
	out := new(BlueprintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlueprintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintModule) DeepCopyInto(out *BlueprintModule) {
	*out = *in
	in.Chart.DeepCopyInto(&out.Chart)
	in.Arguments.DeepCopyInto(&out.Arguments)
	if in.AssetIDs != nil {
		in, out := &in.AssetIDs, &out.AssetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintModule.
func (in *BlueprintModule) DeepCopy() *BlueprintModule {
	if in == nil {
		return nil
	}
	out := new(BlueprintModule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintSpec) DeepCopyInto(out *BlueprintSpec) {
	*out = *in
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make(map[string]BlueprintModule, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(ApplicationDetails)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintSpec.
func (in *BlueprintSpec) DeepCopy() *BlueprintSpec {
	if in == nil {
		return nil
	}
	out := new(BlueprintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintStatus) DeepCopyInto(out *BlueprintStatus) {
	*out = *in
	out.ObservedState = in.ObservedState
	if in.ModulesState != nil {
		in, out := &in.ModulesState, &out.ModulesState
		*out = make(map[string]ObservedState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Releases != nil {
		in, out := &in.Releases, &out.Releases
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintStatus.
func (in *BlueprintStatus) DeepCopy() *BlueprintStatus {
	if in == nil {
		return nil
	}
	out := new(BlueprintStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpec) DeepCopyInto(out *ChartSpec) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpec.
func (in *ChartSpec) DeepCopy() *ChartSpec {
	if in == nil {
		return nil
	}
	out := new(ChartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataContext) DeepCopyInto(out *DataContext) {
	*out = *in
	in.Requirements.DeepCopyInto(&out.Requirements)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataContext.
func (in *DataContext) DeepCopy() *DataContext {
	if in == nil {
		return nil
	}
	out := new(DataContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataFlowStep) DeepCopyInto(out *DataFlowStep) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(StepParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataFlowStep.
func (in *DataFlowStep) DeepCopy() *DataFlowStep {
	if in == nil {
		return nil
	}
	out := new(DataFlowStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataRequirements) DeepCopyInto(out *DataRequirements) {
	*out = *in
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(taxonomy.Interface)
		**out = **in
	}
	out.FlowParams = in.FlowParams
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataRequirements.
func (in *DataRequirements) DeepCopy() *DataRequirements {
	if in == nil {
		return nil
	}
	out := new(DataRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataStore) DeepCopyInto(out *DataStore) {
	*out = *in
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		*out = make(map[string]Vault, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Connection.DeepCopyInto(&out.Connection)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStore.
func (in *DataStore) DeepCopy() *DataStore {
	if in == nil {
		return nil
	}
	out := new(DataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetDetails) DeepCopyInto(out *DatasetDetails) {
	*out = *in
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(DataStore)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetDetails.
func (in *DatasetDetails) DeepCopy() *DatasetDetails {
	if in == nil {
		return nil
	}
	out := new(DatasetDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependency) DeepCopyInto(out *Dependency) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependency.
func (in *Dependency) DeepCopy() *Dependency {
	if in == nil {
		return nil
	}
	out := new(Dependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flow) DeepCopyInto(out *Flow) {
	*out = *in
	if in.SubFlows != nil {
		in, out := &in.SubFlows, &out.SubFlows
		*out = make([]SubFlow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flow.
func (in *Flow) DeepCopy() *Flow {
	if in == nil {
		return nil
	}
	out := new(Flow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowRequirements) DeepCopyInto(out *FlowRequirements) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowRequirements.
func (in *FlowRequirements) DeepCopy() *FlowRequirements {
	if in == nil {
		return nil
	}
	out := new(FlowRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowStatus) DeepCopyInto(out *FlowStatus) {
	*out = *in
	out.ObservedState = in.ObservedState
	if in.SubFlows != nil {
		in, out := &in.SubFlows, &out.SubFlows
		*out = make(map[string]ObservedState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowStatus.
func (in *FlowStatus) DeepCopy() *FlowStatus {
	if in == nil {
		return nil
	}
	out := new(FlowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikApplication) DeepCopyInto(out *FybrikApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikApplication.
func (in *FybrikApplication) DeepCopy() *FybrikApplication {
	if in == nil {
		return nil
	}
	out := new(FybrikApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikApplicationList) DeepCopyInto(out *FybrikApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FybrikApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikApplicationList.
func (in *FybrikApplicationList) DeepCopy() *FybrikApplicationList {
	if in == nil {
		return nil
	}
	out := new(FybrikApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikApplicationSpec) DeepCopyInto(out *FybrikApplicationSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.AppInfo.DeepCopyInto(&out.AppInfo)
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataContext, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikApplicationSpec.
func (in *FybrikApplicationSpec) DeepCopy() *FybrikApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(FybrikApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikApplicationStatus) DeepCopyInto(out *FybrikApplicationStatus) {
	*out = *in
	if in.AssetStates != nil {
		in, out := &in.AssetStates, &out.AssetStates
		*out = make(map[string]AssetState, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Generated != nil {
		in, out := &in.Generated, &out.Generated
		*out = new(ResourceReference)
		**out = **in
	}
	if in.ProvisionedStorage != nil {
		in, out := &in.ProvisionedStorage, &out.ProvisionedStorage
		*out = make(map[string]DatasetDetails, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikApplicationStatus.
func (in *FybrikApplicationStatus) DeepCopy() *FybrikApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(FybrikApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikModule) DeepCopyInto(out *FybrikModule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikModule.
func (in *FybrikModule) DeepCopy() *FybrikModule {
	if in == nil {
		return nil
	}
	out := new(FybrikModule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikModule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikModuleList) DeepCopyInto(out *FybrikModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FybrikModule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikModuleList.
func (in *FybrikModuleList) DeepCopy() *FybrikModuleList {
	if in == nil {
		return nil
	}
	out := new(FybrikModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikModuleSpec) DeepCopyInto(out *FybrikModuleSpec) {
	*out = *in
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]Dependency, len(*in))
		copy(*out, *in)
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]ModuleCapability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Chart.DeepCopyInto(&out.Chart)
	if in.StatusIndicators != nil {
		in, out := &in.StatusIndicators, &out.StatusIndicators
		*out = make([]ResourceStatusIndicator, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikModuleSpec.
func (in *FybrikModuleSpec) DeepCopy() *FybrikModuleSpec {
	if in == nil {
		return nil
	}
	out := new(FybrikModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikModuleStatus) DeepCopyInto(out *FybrikModuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikModuleStatus.
func (in *FybrikModuleStatus) DeepCopy() *FybrikModuleStatus {
	if in == nil {
		return nil
	}
	out := new(FybrikModuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikStorageAccount) DeepCopyInto(out *FybrikStorageAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikStorageAccount.
func (in *FybrikStorageAccount) DeepCopy() *FybrikStorageAccount {
	if in == nil {
		return nil
	}
	out := new(FybrikStorageAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikStorageAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikStorageAccountList) DeepCopyInto(out *FybrikStorageAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FybrikStorageAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikStorageAccountList.
func (in *FybrikStorageAccountList) DeepCopy() *FybrikStorageAccountList {
	if in == nil {
		return nil
	}
	out := new(FybrikStorageAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikStorageAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikStorageAccountSpec) DeepCopyInto(out *FybrikStorageAccountSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikStorageAccountSpec.
func (in *FybrikStorageAccountSpec) DeepCopy() *FybrikStorageAccountSpec {
	if in == nil {
		return nil
	}
	out := new(FybrikStorageAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikStorageAccountStatus) DeepCopyInto(out *FybrikStorageAccountStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikStorageAccountStatus.
func (in *FybrikStorageAccountStatus) DeepCopy() *FybrikStorageAccountStatus {
	if in == nil {
		return nil
	}
	out := new(FybrikStorageAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceDetails) DeepCopyInto(out *InterfaceDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceDetails.
func (in *InterfaceDetails) DeepCopy() *InterfaceDetails {
	if in == nil {
		return nil
	}
	out := new(InterfaceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaBlueprint) DeepCopyInto(out *MetaBlueprint) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaBlueprint.
func (in *MetaBlueprint) DeepCopy() *MetaBlueprint {
	if in == nil {
		return nil
	}
	out := new(MetaBlueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleArguments) DeepCopyInto(out *ModuleArguments) {
	*out = *in
	if in.Assets != nil {
		in, out := &in.Assets, &out.Assets
		*out = make([]AssetContext, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleArguments.
func (in *ModuleArguments) DeepCopy() *ModuleArguments {
	if in == nil {
		return nil
	}
	out := new(ModuleArguments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleCapability) DeepCopyInto(out *ModuleCapability) {
	*out = *in
	if in.SupportedInterfaces != nil {
		in, out := &in.SupportedInterfaces, &out.SupportedInterfaces
		*out = make([]ModuleInOut, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(datacatalog.ResourceDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]ModuleSupportedAction, len(*in))
		copy(*out, *in)
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]Plugin, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleCapability.
func (in *ModuleCapability) DeepCopy() *ModuleCapability {
	if in == nil {
		return nil
	}
	out := new(ModuleCapability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleInOut) DeepCopyInto(out *ModuleInOut) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(taxonomy.Interface)
		**out = **in
	}
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(taxonomy.Interface)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleInOut.
func (in *ModuleInOut) DeepCopy() *ModuleInOut {
	if in == nil {
		return nil
	}
	out := new(ModuleInOut)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleInfo) DeepCopyInto(out *ModuleInfo) {
	*out = *in
	in.Chart.DeepCopyInto(&out.Chart)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleInfo.
func (in *ModuleInfo) DeepCopy() *ModuleInfo {
	if in == nil {
		return nil
	}
	out := new(ModuleInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleSupportedAction) DeepCopyInto(out *ModuleSupportedAction) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleSupportedAction.
func (in *ModuleSupportedAction) DeepCopy() *ModuleSupportedAction {
	if in == nil {
		return nil
	}
	out := new(ModuleSupportedAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservedState) DeepCopyInto(out *ObservedState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservedState.
func (in *ObservedState) DeepCopy() *ObservedState {
	if in == nil {
		return nil
	}
	out := new(ObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plotter) DeepCopyInto(out *Plotter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plotter.
func (in *Plotter) DeepCopy() *Plotter {
	if in == nil {
		return nil
	}
	out := new(Plotter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Plotter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlotterList) DeepCopyInto(out *PlotterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Plotter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlotterList.
func (in *PlotterList) DeepCopy() *PlotterList {
	if in == nil {
		return nil
	}
	out := new(PlotterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlotterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlotterSpec) DeepCopyInto(out *PlotterSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.AppInfo.DeepCopyInto(&out.AppInfo)
	if in.Assets != nil {
		in, out := &in.Assets, &out.Assets
		*out = make(map[string]AssetDetails, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Flows != nil {
		in, out := &in.Flows, &out.Flows
		*out = make([]Flow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]Template, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlotterSpec.
func (in *PlotterSpec) DeepCopy() *PlotterSpec {
	if in == nil {
		return nil
	}
	out := new(PlotterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlotterStatus) DeepCopyInto(out *PlotterStatus) {
	*out = *in
	out.ObservedState = in.ObservedState
	if in.Flows != nil {
		in, out := &in.Flows, &out.Flows
		*out = make(map[string]FlowStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Assets != nil {
		in, out := &in.Assets, &out.Assets
		*out = make(map[string]ObservedState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Blueprints != nil {
		in, out := &in.Blueprints, &out.Blueprints
		*out = make(map[string]MetaBlueprint, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	if in.ReadyTimestamp != nil {
		in, out := &in.ReadyTimestamp, &out.ReadyTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlotterStatus.
func (in *PlotterStatus) DeepCopy() *PlotterStatus {
	if in == nil {
		return nil
	}
	out := new(PlotterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReference) DeepCopyInto(out *ResourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReference.
func (in *ResourceReference) DeepCopy() *ResourceReference {
	if in == nil {
		return nil
	}
	out := new(ResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatusIndicator) DeepCopyInto(out *ResourceStatusIndicator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatusIndicator.
func (in *ResourceStatusIndicator) DeepCopy() *ResourceStatusIndicator {
	if in == nil {
		return nil
	}
	out := new(ResourceStatusIndicator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	in.WorkloadSelector.DeepCopyInto(&out.WorkloadSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepArgument) DeepCopyInto(out *StepArgument) {
	*out = *in
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(datacatalog.ResourceDetails)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepArgument.
func (in *StepArgument) DeepCopy() *StepArgument {
	if in == nil {
		return nil
	}
	out := new(StepArgument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepParameters) DeepCopyInto(out *StepParameters) {
	*out = *in
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]*StepArgument, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StepArgument)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(datacatalog.ResourceDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]taxonomy.Action, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepParameters.
func (in *StepParameters) DeepCopy() *StepParameters {
	if in == nil {
		return nil
	}
	out := new(StepParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubFlow) DeepCopyInto(out *SubFlow) {
	*out = *in
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]SubFlowTrigger, len(*in))
		copy(*out, *in)
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([][]DataFlowStep, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]DataFlowStep, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubFlow.
func (in *SubFlow) DeepCopy() *SubFlow {
	if in == nil {
		return nil
	}
	out := new(SubFlow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]ModuleInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault) DeepCopyInto(out *Vault) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault.
func (in *Vault) DeepCopy() *Vault {
	if in == nil {
		return nil
	}
	out := new(Vault)
	in.DeepCopyInto(out)
	return out
}
