package storage

import (
	pub "github.com/go-ap/activitypub"
)

type Store interface {
	WriteStore
	ReadStore
}

// ReadStore
type ReadStore interface {
	// Load returns an Item or an ItemCollection from an IRI
	Load(pub.IRI) (pub.Item, error)
}

// Store saves ActivityStreams objects.
type WriteStore interface {
	// Save saves the incoming ActivityStreams Object, and returns it together with any properties
	// populated by the method's side effects. (eg, Published property can point to the current time, etc).
	Save(pub.Item) (pub.Item, error)
	// Delete deletes the incoming ActivityStreams Object, and returns the resulting Tombstone.
	Delete(pub.Item) (pub.Item, error)
}

type CollectionStore interface {
	// Create creates the "col" collection.
	Create(col pub.CollectionInterface) (pub.CollectionInterface, error)
	// AddTo adds "it" element to the "col" collection.
	AddTo(col pub.IRI, it pub.Item) error
	// RemoveFrom removes "it" item from "col" collection
	RemoveFrom(col pub.IRI, it pub.Item) error
}
