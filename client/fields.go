package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/manutara/manutara/api/v1"
)

type FieldsClient struct {
	namespace string
}

func (*FieldsClient) Create(*v1.Field) (*v1.Field, error) {
	panic("implement me")
}

func (*FieldsClient) Update(*v1.Field) (*v1.Field, error) {
	panic("implement me")
}

func (*FieldsClient) UpdateStatus(*v1.Field) (*v1.Field, error) {
	panic("implement me")
}

func (*FieldsClient) Delete(name string, options *metav1.DeleteOptions) error {
	panic("implement me")
}

func (*FieldsClient) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	panic("implement me")
}

func (*FieldsClient) Get(name string, options metav1.GetOptions) (*v1.Field, error) {
	panic("implement me")
}

func (*FieldsClient) List(opts metav1.ListOptions) (*v1.FieldList, error) {
	panic("implement me")
}

func (*FieldsClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	panic("implement me")
}

func (*FieldsClient) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Field, err error) {
	panic("implement me")
}
