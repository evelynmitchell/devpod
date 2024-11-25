// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// TranslateVClusterResourceNameLister helps list TranslateVClusterResourceNames.
// All objects returned here must be treated as read-only.
type TranslateVClusterResourceNameLister interface {
	// List lists all TranslateVClusterResourceNames in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.TranslateVClusterResourceName, err error)
	// Get retrieves the TranslateVClusterResourceName from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.TranslateVClusterResourceName, error)
	TranslateVClusterResourceNameListerExpansion
}

// translateVClusterResourceNameLister implements the TranslateVClusterResourceNameLister interface.
type translateVClusterResourceNameLister struct {
	listers.ResourceIndexer[*v1.TranslateVClusterResourceName]
}

// NewTranslateVClusterResourceNameLister returns a new TranslateVClusterResourceNameLister.
func NewTranslateVClusterResourceNameLister(indexer cache.Indexer) TranslateVClusterResourceNameLister {
	return &translateVClusterResourceNameLister{listers.New[*v1.TranslateVClusterResourceName](indexer, v1.Resource("translatevclusterresourcename"))}
}