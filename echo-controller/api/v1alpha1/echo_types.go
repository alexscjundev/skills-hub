package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	GroupVersion  = schema.GroupVersion{Group: "playground.aljun.dev", Version: "v1alpha1"}
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	AddToScheme   = SchemeBuilder.AddToScheme
)

type EchoSpec struct {
	Message string `json:"message,omitempty"`
}

type EchoStatus struct {
	ObservedMessage string `json:"observedMessage,omitempty"`
	LastReconciled  string `json:"lastReconciled,omitempty"`
}

type Echo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EchoSpec   `json:"spec,omitempty"`
	Status EchoStatus `json:"status,omitempty"`
}

type EchoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Echo `json:"items"`
}

func (in *Echo) DeepCopyInto(out *Echo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

func (in *Echo) DeepCopy() *Echo {
	if in == nil {
		return nil
	}
	out := new(Echo)
	in.DeepCopyInto(out)
	return out
}

func (in *Echo) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

func (in *EchoList) DeepCopyInto(out *EchoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Echo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *EchoList) DeepCopy() *EchoList {
	if in == nil {
		return nil
	}
	out := new(EchoList)
	in.DeepCopyInto(out)
	return out
}

func (in *EchoList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

func init() {
	SchemeBuilder.Register(&Echo{}, &EchoList{})
}
