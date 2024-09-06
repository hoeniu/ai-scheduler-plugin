// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	resourcev1alpha1 "ai.plugin/scheduler/apis/resource/v1alpha1"
	versioned "ai.plugin/scheduler/pkg/client/clientset/versioned"
	internalinterfaces "ai.plugin/scheduler/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "ai.plugin/scheduler/pkg/client/listers/resource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MachineInfoInformer provides access to a shared informer and lister for
// MachineInfos.
type MachineInfoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineInfoLister
}

type machineInfoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMachineInfoInformer constructs a new informer for MachineInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineInfoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMachineInfoInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMachineInfoInformer constructs a new informer for MachineInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMachineInfoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().MachineInfos(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().MachineInfos(namespace).Watch(context.TODO(), options)
			},
		},
		&resourcev1alpha1.MachineInfo{},
		resyncPeriod,
		indexers,
	)
}

func (f *machineInfoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMachineInfoInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *machineInfoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resourcev1alpha1.MachineInfo{}, f.defaultInformer)
}

func (f *machineInfoInformer) Lister() v1alpha1.MachineInfoLister {
	return v1alpha1.NewMachineInfoLister(f.Informer().GetIndexer())
}
