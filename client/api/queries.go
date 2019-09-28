package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	v1 "github.com/manutara/manutara/api/v1"
)

// Queries is a client interface for managing GraphQL query resources
type Queries interface {
	Create(*v1.Query) (*v1.Query, error)
	Update(*v1.Query) (*v1.Query, error)
	UpdateStatus(*v1.Query) (*v1.Query, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Query, error)
	List(opts metav1.ListOptions) (*v1.QueryList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Query, err error)
}
