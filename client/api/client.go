package api

// Client returns an interface for the manutara clientset
type Client interface {
	// Mutations returns a REST client for managing mutation resources
	Mutations(ns string) Mutations

	// Queries returns a REST client for managing query resources
	Queries(ns string) Queries

	// Subscription returns a REST client for managing subscription resources
	Subscriptions(ns string) Subscriptions

	// Services returns a REST client for managing service resources
	Services(ns string) Services

	// DataTypes returns a REST client for managing datatype resources
	DataTypes(ns string) DataTypes

	// DataStores returns a REST client for managing datastore resources
	DataStores(ns string) DataStores

	// DataStoreBindings returns a REST client for managing datastore binding resources
	DataStoreBindings(ns string) DataStoreBindings
}
