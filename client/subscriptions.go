package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/manutara/manutara/api/v1"
)

type SubscriptionsClient struct {
	namespace string
}

func (*SubscriptionsClient) Create(*v1.Subscription) (*v1.Subscription, error) {
	panic("implement me")
}

func (*SubscriptionsClient) Update(*v1.Subscription) (*v1.Subscription, error) {
	panic("implement me")
}

func (*SubscriptionsClient) UpdateStatus(*v1.Subscription) (*v1.Subscription, error) {
	panic("implement me")
}

func (*SubscriptionsClient) Delete(name string, options *metav1.DeleteOptions) error {
	panic("implement me")
}

func (*SubscriptionsClient) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	panic("implement me")
}

func (*SubscriptionsClient) Get(name string, options metav1.GetOptions) (*v1.Subscription, error) {
	panic("implement me")
}

func (*SubscriptionsClient) List(opts metav1.ListOptions) (*v1.SubscriptionList, error) {
	panic("implement me")
}

func (*SubscriptionsClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	panic("implement me")
}

func (*SubscriptionsClient) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Subscription, err error) {
	panic("implement me")
}

