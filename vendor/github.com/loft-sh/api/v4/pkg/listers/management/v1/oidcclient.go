// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// OIDCClientLister helps list OIDCClients.
// All objects returned here must be treated as read-only.
type OIDCClientLister interface {
	// List lists all OIDCClients in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.OIDCClient, err error)
	// Get retrieves the OIDCClient from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.OIDCClient, error)
	OIDCClientListerExpansion
}

// oIDCClientLister implements the OIDCClientLister interface.
type oIDCClientLister struct {
	listers.ResourceIndexer[*v1.OIDCClient]
}

// NewOIDCClientLister returns a new OIDCClientLister.
func NewOIDCClientLister(indexer cache.Indexer) OIDCClientLister {
	return &oIDCClientLister{listers.New[*v1.OIDCClient](indexer, v1.Resource("oidcclient"))}
}