// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// LicenseTokenLister helps list LicenseTokens.
// All objects returned here must be treated as read-only.
type LicenseTokenLister interface {
	// List lists all LicenseTokens in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.LicenseToken, err error)
	// Get retrieves the LicenseToken from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.LicenseToken, error)
	LicenseTokenListerExpansion
}

// licenseTokenLister implements the LicenseTokenLister interface.
type licenseTokenLister struct {
	listers.ResourceIndexer[*v1.LicenseToken]
}

// NewLicenseTokenLister returns a new LicenseTokenLister.
func NewLicenseTokenLister(indexer cache.Indexer) LicenseTokenLister {
	return &licenseTokenLister{listers.New[*v1.LicenseToken](indexer, v1.Resource("licensetoken"))}
}