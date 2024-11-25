// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// RunnerLister helps list Runners.
// All objects returned here must be treated as read-only.
type RunnerLister interface {
	// List lists all Runners in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Runner, err error)
	// Get retrieves the Runner from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Runner, error)
	RunnerListerExpansion
}

// runnerLister implements the RunnerLister interface.
type runnerLister struct {
	listers.ResourceIndexer[*v1.Runner]
}

// NewRunnerLister returns a new RunnerLister.
func NewRunnerLister(indexer cache.Indexer) RunnerLister {
	return &runnerLister{listers.New[*v1.Runner](indexer, v1.Resource("runner"))}
}