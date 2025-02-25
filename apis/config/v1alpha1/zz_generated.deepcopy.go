//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Feature) DeepCopyInto(out *Feature) {
	*out = *in
	out.Status = in.Status
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Feature.
func (in *Feature) DeepCopy() *Feature {
	if in == nil {
		return nil
	}
	out := new(Feature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Feature) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGate) DeepCopyInto(out *FeatureGate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGate.
func (in *FeatureGate) DeepCopy() *FeatureGate {
	if in == nil {
		return nil
	}
	out := new(FeatureGate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureGate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGateList) DeepCopyInto(out *FeatureGateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FeatureGate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGateList.
func (in *FeatureGateList) DeepCopy() *FeatureGateList {
	if in == nil {
		return nil
	}
	out := new(FeatureGateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureGateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGateSpec) DeepCopyInto(out *FeatureGateSpec) {
	*out = *in
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]FeatureReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGateSpec.
func (in *FeatureGateSpec) DeepCopy() *FeatureGateSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureGateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGateStatus) DeepCopyInto(out *FeatureGateStatus) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActivatedFeatures != nil {
		in, out := &in.ActivatedFeatures, &out.ActivatedFeatures
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeactivatedFeatures != nil {
		in, out := &in.DeactivatedFeatures, &out.DeactivatedFeatures
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UnavailableFeatures != nil {
		in, out := &in.UnavailableFeatures, &out.UnavailableFeatures
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGateStatus.
func (in *FeatureGateStatus) DeepCopy() *FeatureGateStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureGateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureList) DeepCopyInto(out *FeatureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Feature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureList.
func (in *FeatureList) DeepCopy() *FeatureList {
	if in == nil {
		return nil
	}
	out := new(FeatureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureReference) DeepCopyInto(out *FeatureReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureReference.
func (in *FeatureReference) DeepCopy() *FeatureReference {
	if in == nil {
		return nil
	}
	out := new(FeatureReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSpec) DeepCopyInto(out *FeatureSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSpec.
func (in *FeatureSpec) DeepCopy() *FeatureSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStatus) DeepCopyInto(out *FeatureStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStatus.
func (in *FeatureStatus) DeepCopy() *FeatureStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureStatus)
	in.DeepCopyInto(out)
	return out
}
