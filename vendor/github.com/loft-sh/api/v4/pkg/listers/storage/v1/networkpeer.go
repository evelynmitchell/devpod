// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// NetworkPeerLister helps list NetworkPeers.
// All objects returned here must be treated as read-only.
type NetworkPeerLister interface {
	// List lists all NetworkPeers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NetworkPeer, err error)
	// Get retrieves the NetworkPeer from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.NetworkPeer, error)
	NetworkPeerListerExpansion
}

// networkPeerLister implements the NetworkPeerLister interface.
type networkPeerLister struct {
	listers.ResourceIndexer[*v1.NetworkPeer]
}

// NewNetworkPeerLister returns a new NetworkPeerLister.
func NewNetworkPeerLister(indexer cache.Indexer) NetworkPeerLister {
	return &networkPeerLister{listers.New[*v1.NetworkPeer](indexer, v1.Resource("networkpeer"))}
}