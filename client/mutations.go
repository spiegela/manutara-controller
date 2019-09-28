package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/manutara/manutara/api/v1"
)

type MutationsClient struct {
	namespace string
}

func (*MutationsClient) Create(*v1.Mutation) (*v1.Mutation, error) {
	panic("implement me")
}

func (*MutationsClient) Update(*v1.Mutation) (*v1.Mutation, error) {
	panic("implement me")
}

func (*MutationsClient) UpdateStatus(*v1.Mutation) (*v1.Mutation, error) {
	panic("implement me")
}

func (*MutationsClient) Delete(name string, options *metav1.DeleteOptions) error {
	panic("implement me")
}

func (*MutationsClient) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	panic("implement me")
}

func (*MutationsClient) Get(name string, options metav1.GetOptions) (*v1.Mutation, error) {
	panic("implement me")
}

func (*MutationsClient) List(opts metav1.ListOptions) (*v1.MutationList, error) {
	panic("implement me")
}

func (*MutationsClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	panic("implement me")
}

func (*MutationsClient) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Mutation, err error) {
	panic("implement me")
}
