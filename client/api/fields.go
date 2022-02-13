package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	v1 "github.com/manutara/manutara/api/v1"
)

// Fields is a client interface for managing GraphQL mutation resources
type Fields interface {
	Create(*v1.Field) (*v1.Field, error)
	Update(*v1.Field) (*v1.Field, error)
	UpdateStatus(*v1.Field) (*v1.Field, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Field, error)
	List(opts metav1.ListOptions) (*v1.FieldList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Field, err error)
}
