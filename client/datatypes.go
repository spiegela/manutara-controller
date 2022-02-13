package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/manutara/manutara/api/v1"
)

type DataTypesClient struct {
	namespace string
}

func (*DataTypesClient) Create() (*v1.DataType, error) {
	panic("implement me")
}

func (*DataTypesClient) Update(*v1.DataType) (*v1.DataType, error) {
	panic("implement me")
}

func (*DataTypesClient) UpdateStatus(*v1.DataType) (*v1.DataType, error) {
	panic("implement me")
}

func (*DataTypesClient) Delete(name string, options *metav1.DeleteOptions) error {
	panic("implement me")
}

func (*DataTypesClient) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	panic("implement me")
}

func (*DataTypesClient) Get(name string, options metav1.GetOptions) (*v1.DataType, error) {
	panic("implement me")
}

func (*DataTypesClient) List(opts metav1.ListOptions) (*v1.DataTypeList, error) {
	panic("implement me")
}

func (*DataTypesClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	panic("implement me")
}

func (*DataTypesClient) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DataType, err error) {
	panic("implement me")
}

