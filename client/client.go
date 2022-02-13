package client

import (
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"

	"github.com/manutara/manutara/api/v1"
	"github.com/manutara/manutara/client/api"
)

// Client represents underlying CRD client struct
// nolint
type Client struct {
	restClient rest.Interface
}

// NewForConfig returns Interface to access Notifiers()
func NewForConfig(c *rest.Config) (*Client, error) {
	err := v1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}
	config := *c
	config.ContentConfig.GroupVersion = &v1.GroupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{restClient: client}, nil
}

// Fields returns a client for mutation resources
func Fields(ns string) api.Fields {
	return &FieldsClient{
		namespace: ns,
	}
}

// DataTypes returns a client for mutation resources
func DataTypes(ns string) api.DataTypes {
	return &DataTypesClient{
		namespace: ns,
	}
}
