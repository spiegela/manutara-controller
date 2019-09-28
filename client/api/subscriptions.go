package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	v1 "github.com/manutara/manutara/api/v1"
)

// Subscriptions is a client interface for managing GraphQL subscription resources
type Subscriptions interface {
	Create(*v1.Subscription) (*v1.Subscription, error)
	Update(*v1.Subscription) (*v1.Subscription, error)
	UpdateStatus(*v1.Subscription) (*v1.Subscription, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Subscription, error)
	List(opts metav1.ListOptions) (*v1.SubscriptionList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Subscription, err error)
}
