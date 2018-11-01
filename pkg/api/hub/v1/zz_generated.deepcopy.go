// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environs) DeepCopyInto(out *Environs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environs.
func (in *Environs) DeepCopy() *Environs {
	if in == nil {
		return nil
	}
	out := new(Environs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hub) DeepCopyInto(out *Hub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.View.DeepCopyInto(&out.View)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hub.
func (in *Hub) DeepCopy() *Hub {
	if in == nil {
		return nil
	}
	out := new(Hub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubList) DeepCopyInto(out *HubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubList.
func (in *HubList) DeepCopy() *HubList {
	if in == nil {
		return nil
	}
	out := new(HubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubSpec) DeepCopyInto(out *HubSpec) {
	*out = *in
	if in.Environs != nil {
		in, out := &in.Environs, &out.Environs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImageTagMap != nil {
		in, out := &in.ImageTagMap, &out.ImageTagMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubSpec.
func (in *HubSpec) DeepCopy() *HubSpec {
	if in == nil {
		return nil
	}
	out := new(HubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubStatus) DeepCopyInto(out *HubStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubStatus.
func (in *HubStatus) DeepCopy() *HubStatus {
	if in == nil {
		return nil
	}
	out := new(HubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubView) DeepCopyInto(out *HubView) {
	*out = *in
	if in.Clones != nil {
		in, out := &in.Clones, &out.Clones
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CertificateNames != nil {
		in, out := &in.CertificateNames, &out.CertificateNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Environs != nil {
		in, out := &in.Environs, &out.Environs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContainerTags != nil {
		in, out := &in.ContainerTags, &out.ContainerTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubView.
func (in *HubView) DeepCopy() *HubView {
	if in == nil {
		return nil
	}
	out := new(HubView)
	in.DeepCopyInto(out)
	return out
}
