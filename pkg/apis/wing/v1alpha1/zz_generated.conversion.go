// Copyright Jetstack Ltd. See LICENSE for details.
// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	wing "github.com/jetstack/tarmak/pkg/apis/wing"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Instance_To_wing_Instance,
		Convert_wing_Instance_To_v1alpha1_Instance,
		Convert_v1alpha1_InstanceList_To_wing_InstanceList,
		Convert_wing_InstanceList_To_v1alpha1_InstanceList,
		Convert_v1alpha1_InstanceSpec_To_wing_InstanceSpec,
		Convert_wing_InstanceSpec_To_v1alpha1_InstanceSpec,
		Convert_v1alpha1_InstanceSpecManifest_To_wing_InstanceSpecManifest,
		Convert_wing_InstanceSpecManifest_To_v1alpha1_InstanceSpecManifest,
		Convert_v1alpha1_InstanceStatus_To_wing_InstanceStatus,
		Convert_wing_InstanceStatus_To_v1alpha1_InstanceStatus,
		Convert_v1alpha1_InstanceStatusManifest_To_wing_InstanceStatusManifest,
		Convert_wing_InstanceStatusManifest_To_v1alpha1_InstanceStatusManifest,
	)
}

func autoConvert_v1alpha1_Instance_To_wing_Instance(in *Instance, out *wing.Instance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.InstanceID = in.InstanceID
	out.InstancePool = in.InstancePool
	out.Spec = (*wing.InstanceSpec)(unsafe.Pointer(in.Spec))
	out.Status = (*wing.InstanceStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_v1alpha1_Instance_To_wing_Instance is an autogenerated conversion function.
func Convert_v1alpha1_Instance_To_wing_Instance(in *Instance, out *wing.Instance, s conversion.Scope) error {
	return autoConvert_v1alpha1_Instance_To_wing_Instance(in, out, s)
}

func autoConvert_wing_Instance_To_v1alpha1_Instance(in *wing.Instance, out *Instance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.InstanceID = in.InstanceID
	out.InstancePool = in.InstancePool
	out.Spec = (*InstanceSpec)(unsafe.Pointer(in.Spec))
	out.Status = (*InstanceStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_wing_Instance_To_v1alpha1_Instance is an autogenerated conversion function.
func Convert_wing_Instance_To_v1alpha1_Instance(in *wing.Instance, out *Instance, s conversion.Scope) error {
	return autoConvert_wing_Instance_To_v1alpha1_Instance(in, out, s)
}

func autoConvert_v1alpha1_InstanceList_To_wing_InstanceList(in *InstanceList, out *wing.InstanceList, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Items = *(*[]wing.Instance)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_InstanceList_To_wing_InstanceList is an autogenerated conversion function.
func Convert_v1alpha1_InstanceList_To_wing_InstanceList(in *InstanceList, out *wing.InstanceList, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceList_To_wing_InstanceList(in, out, s)
}

func autoConvert_wing_InstanceList_To_v1alpha1_InstanceList(in *wing.InstanceList, out *InstanceList, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Items = *(*[]Instance)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_wing_InstanceList_To_v1alpha1_InstanceList is an autogenerated conversion function.
func Convert_wing_InstanceList_To_v1alpha1_InstanceList(in *wing.InstanceList, out *InstanceList, s conversion.Scope) error {
	return autoConvert_wing_InstanceList_To_v1alpha1_InstanceList(in, out, s)
}

func autoConvert_v1alpha1_InstanceSpec_To_wing_InstanceSpec(in *InstanceSpec, out *wing.InstanceSpec, s conversion.Scope) error {
	out.Converge = (*wing.InstanceSpecManifest)(unsafe.Pointer(in.Converge))
	out.DryRun = (*wing.InstanceSpecManifest)(unsafe.Pointer(in.DryRun))
	return nil
}

// Convert_v1alpha1_InstanceSpec_To_wing_InstanceSpec is an autogenerated conversion function.
func Convert_v1alpha1_InstanceSpec_To_wing_InstanceSpec(in *InstanceSpec, out *wing.InstanceSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceSpec_To_wing_InstanceSpec(in, out, s)
}

func autoConvert_wing_InstanceSpec_To_v1alpha1_InstanceSpec(in *wing.InstanceSpec, out *InstanceSpec, s conversion.Scope) error {
	out.Converge = (*InstanceSpecManifest)(unsafe.Pointer(in.Converge))
	out.DryRun = (*InstanceSpecManifest)(unsafe.Pointer(in.DryRun))
	return nil
}

// Convert_wing_InstanceSpec_To_v1alpha1_InstanceSpec is an autogenerated conversion function.
func Convert_wing_InstanceSpec_To_v1alpha1_InstanceSpec(in *wing.InstanceSpec, out *InstanceSpec, s conversion.Scope) error {
	return autoConvert_wing_InstanceSpec_To_v1alpha1_InstanceSpec(in, out, s)
}

func autoConvert_v1alpha1_InstanceSpecManifest_To_wing_InstanceSpecManifest(in *InstanceSpecManifest, out *wing.InstanceSpecManifest, s conversion.Scope) error {
	out.Path = in.Path
	out.Hash = in.Hash
	out.RequestTimestamp = in.RequestTimestamp
	return nil
}

// Convert_v1alpha1_InstanceSpecManifest_To_wing_InstanceSpecManifest is an autogenerated conversion function.
func Convert_v1alpha1_InstanceSpecManifest_To_wing_InstanceSpecManifest(in *InstanceSpecManifest, out *wing.InstanceSpecManifest, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceSpecManifest_To_wing_InstanceSpecManifest(in, out, s)
}

func autoConvert_wing_InstanceSpecManifest_To_v1alpha1_InstanceSpecManifest(in *wing.InstanceSpecManifest, out *InstanceSpecManifest, s conversion.Scope) error {
	out.Path = in.Path
	out.Hash = in.Hash
	out.RequestTimestamp = in.RequestTimestamp
	return nil
}

// Convert_wing_InstanceSpecManifest_To_v1alpha1_InstanceSpecManifest is an autogenerated conversion function.
func Convert_wing_InstanceSpecManifest_To_v1alpha1_InstanceSpecManifest(in *wing.InstanceSpecManifest, out *InstanceSpecManifest, s conversion.Scope) error {
	return autoConvert_wing_InstanceSpecManifest_To_v1alpha1_InstanceSpecManifest(in, out, s)
}

func autoConvert_v1alpha1_InstanceStatus_To_wing_InstanceStatus(in *InstanceStatus, out *wing.InstanceStatus, s conversion.Scope) error {
	out.Converge = (*wing.InstanceStatusManifest)(unsafe.Pointer(in.Converge))
	out.DryRun = (*wing.InstanceStatusManifest)(unsafe.Pointer(in.DryRun))
	return nil
}

// Convert_v1alpha1_InstanceStatus_To_wing_InstanceStatus is an autogenerated conversion function.
func Convert_v1alpha1_InstanceStatus_To_wing_InstanceStatus(in *InstanceStatus, out *wing.InstanceStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceStatus_To_wing_InstanceStatus(in, out, s)
}

func autoConvert_wing_InstanceStatus_To_v1alpha1_InstanceStatus(in *wing.InstanceStatus, out *InstanceStatus, s conversion.Scope) error {
	out.Converge = (*InstanceStatusManifest)(unsafe.Pointer(in.Converge))
	out.DryRun = (*InstanceStatusManifest)(unsafe.Pointer(in.DryRun))
	return nil
}

// Convert_wing_InstanceStatus_To_v1alpha1_InstanceStatus is an autogenerated conversion function.
func Convert_wing_InstanceStatus_To_v1alpha1_InstanceStatus(in *wing.InstanceStatus, out *InstanceStatus, s conversion.Scope) error {
	return autoConvert_wing_InstanceStatus_To_v1alpha1_InstanceStatus(in, out, s)
}

func autoConvert_v1alpha1_InstanceStatusManifest_To_wing_InstanceStatusManifest(in *InstanceStatusManifest, out *wing.InstanceStatusManifest, s conversion.Scope) error {
	out.State = wing.InstanceManifestState(in.State)
	out.Hash = in.Hash
	out.LastUpdateTimestamp = in.LastUpdateTimestamp
	out.Messages = *(*[]string)(unsafe.Pointer(&in.Messages))
	out.ExitCodes = *(*[]int)(unsafe.Pointer(&in.ExitCodes))
	return nil
}

// Convert_v1alpha1_InstanceStatusManifest_To_wing_InstanceStatusManifest is an autogenerated conversion function.
func Convert_v1alpha1_InstanceStatusManifest_To_wing_InstanceStatusManifest(in *InstanceStatusManifest, out *wing.InstanceStatusManifest, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceStatusManifest_To_wing_InstanceStatusManifest(in, out, s)
}

func autoConvert_wing_InstanceStatusManifest_To_v1alpha1_InstanceStatusManifest(in *wing.InstanceStatusManifest, out *InstanceStatusManifest, s conversion.Scope) error {
	out.State = InstanceManifestState(in.State)
	out.Hash = in.Hash
	out.LastUpdateTimestamp = in.LastUpdateTimestamp
	out.Messages = *(*[]string)(unsafe.Pointer(&in.Messages))
	out.ExitCodes = *(*[]int)(unsafe.Pointer(&in.ExitCodes))
	return nil
}

// Convert_wing_InstanceStatusManifest_To_v1alpha1_InstanceStatusManifest is an autogenerated conversion function.
func Convert_wing_InstanceStatusManifest_To_v1alpha1_InstanceStatusManifest(in *wing.InstanceStatusManifest, out *InstanceStatusManifest, s conversion.Scope) error {
	return autoConvert_wing_InstanceStatusManifest_To_v1alpha1_InstanceStatusManifest(in, out, s)
}
