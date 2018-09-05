// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.IsolationGroups != nil {
		in, out := &in.IsolationGroups, &out.IsolationGroups
		*out = make([]IsolationGroup, len(*in))
		copy(*out, *in)
	}
	out.Resources = in.Resources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IsolationGroup) DeepCopyInto(out *IsolationGroup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IsolationGroup.
func (in *IsolationGroup) DeepCopy() *IsolationGroup {
	if in == nil {
		return nil
	}
	out := new(IsolationGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M3DBCluster) DeepCopyInto(out *M3DBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M3DBCluster.
func (in *M3DBCluster) DeepCopy() *M3DBCluster {
	if in == nil {
		return nil
	}
	out := new(M3DBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M3DBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M3DBClusterList) DeepCopyInto(out *M3DBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]M3DBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M3DBClusterList.
func (in *M3DBClusterList) DeepCopy() *M3DBClusterList {
	if in == nil {
		return nil
	}
	out := new(M3DBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M3DBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M3DBStatus) DeepCopyInto(out *M3DBStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M3DBStatus.
func (in *M3DBStatus) DeepCopy() *M3DBStatus {
	if in == nil {
		return nil
	}
	out := new(M3DBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryCPU) DeepCopyInto(out *MemoryCPU) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryCPU.
func (in *MemoryCPU) DeepCopy() *MemoryCPU {
	if in == nil {
		return nil
	}
	out := new(MemoryCPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.Requests = in.Requests
	out.Limits = in.Limits
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}
