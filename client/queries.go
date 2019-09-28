package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/manutara/manutara/api/v1"
)

type QueriesClient struct {
	namespace string
}

func (*QueriesClient) Create(*v1.Query) (*v1.Query, error) {
	panic("implement me")
}

func (*QueriesClient) Update(*v1.Query) (*v1.Query, error) {
	panic("implement me")
}

func (*QueriesClient) UpdateStatus(*v1.Query) (*v1.Query, error) {
	panic("implement me")
}

func (*QueriesClient) Delete(name string, options *metav1.DeleteOptions) error {
	panic("implement me")
}

func (*QueriesClient) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	panic("implement me")
}

func (*QueriesClient) Get(name string, options metav1.GetOptions) (*v1.Query, error) {
	panic("implement me")
}

func (*QueriesClient) List(opts metav1.ListOptions) (*v1.QueryList, error) {
	panic("implement me")
}

func (*QueriesClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	panic("implement me")
}

func (*QueriesClient) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Query, err error) {
	panic("implement me")
}

