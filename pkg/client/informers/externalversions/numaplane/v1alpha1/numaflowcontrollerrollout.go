/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	numaplanev1alpha1 "github.com/numaproj/numaplane/pkg/apis/numaplane/v1alpha1"
	versioned "github.com/numaproj/numaplane/pkg/client/clientset/versioned"
	internalinterfaces "github.com/numaproj/numaplane/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/numaproj/numaplane/pkg/client/listers/numaplane/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NumaflowControllerRolloutInformer provides access to a shared informer and lister for
// NumaflowControllerRollouts.
type NumaflowControllerRolloutInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NumaflowControllerRolloutLister
}

type numaflowControllerRolloutInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNumaflowControllerRolloutInformer constructs a new informer for NumaflowControllerRollout type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNumaflowControllerRolloutInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNumaflowControllerRolloutInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNumaflowControllerRolloutInformer constructs a new informer for NumaflowControllerRollout type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNumaflowControllerRolloutInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NumaplaneV1alpha1().NumaflowControllerRollouts(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NumaplaneV1alpha1().NumaflowControllerRollouts(namespace).Watch(context.TODO(), options)
			},
		},
		&numaplanev1alpha1.NumaflowControllerRollout{},
		resyncPeriod,
		indexers,
	)
}

func (f *numaflowControllerRolloutInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNumaflowControllerRolloutInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *numaflowControllerRolloutInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&numaplanev1alpha1.NumaflowControllerRollout{}, f.defaultInformer)
}

func (f *numaflowControllerRolloutInformer) Lister() v1alpha1.NumaflowControllerRolloutLister {
	return v1alpha1.NewNumaflowControllerRolloutLister(f.Informer().GetIndexer())
}