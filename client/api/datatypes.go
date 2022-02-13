package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	v1 "github.com/manutara/manutara/api/v1"
)

type DataTypes interface {
	Create(*v1.DataType) (*v1.DataType, error)
	Update(*v1.DataType) (*v1.DataType, error)
	UpdateStatus(*v1.DataType) (*v1.DataType, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.DataType, error)
	List(opts metav1.ListOptions) (*v1.DataTypeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DataType, err error)
}
