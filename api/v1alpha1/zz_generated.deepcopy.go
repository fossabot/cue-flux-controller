//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossNamespaceSourceReference) DeepCopyInto(out *CrossNamespaceSourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossNamespaceSourceReference.
func (in *CrossNamespaceSourceReference) DeepCopy() *CrossNamespaceSourceReference {
	if in == nil {
		return nil
	}
	out := new(CrossNamespaceSourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CueInstance) DeepCopyInto(out *CueInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CueInstance.
func (in *CueInstance) DeepCopy() *CueInstance {
	if in == nil {
		return nil
	}
	out := new(CueInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CueInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CueInstanceList) DeepCopyInto(out *CueInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CueInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CueInstanceList.
func (in *CueInstanceList) DeepCopy() *CueInstanceList {
	if in == nil {
		return nil
	}
	out := new(CueInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CueInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CueInstanceSpec) DeepCopyInto(out *CueInstanceSpec) {
	*out = *in
	in.Interval.DeepCopyInto(&out.Interval)
	out.SourceRef = in.SourceRef
	if in.Inject != nil {
		in, out := &in.Inject, &out.Inject
		*out = make([]InjectItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CueInstanceSpec.
func (in *CueInstanceSpec) DeepCopy() *CueInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(CueInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CueInstanceStatus) DeepCopyInto(out *CueInstanceStatus) {
	*out = *in
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Inventory != nil {
		in, out := &in.Inventory, &out.Inventory
		*out = new(ResourceInventory)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CueInstanceStatus.
func (in *CueInstanceStatus) DeepCopy() *CueInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(CueInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InjectItem) DeepCopyInto(out *InjectItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InjectItem.
func (in *InjectItem) DeepCopy() *InjectItem {
	if in == nil {
		return nil
	}
	out := new(InjectItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInventory) DeepCopyInto(out *ResourceInventory) {
	*out = *in
	if in.Entries != nil {
		in, out := &in.Entries, &out.Entries
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInventory.
func (in *ResourceInventory) DeepCopy() *ResourceInventory {
	if in == nil {
		return nil
	}
	out := new(ResourceInventory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRef.
func (in *ResourceRef) DeepCopy() *ResourceRef {
	if in == nil {
		return nil
	}
	out := new(ResourceRef)
	in.DeepCopyInto(out)
	return out
}
