package multicluster

import (
	"github.com/ibm/the-mesh-for-data/manager/apis/app/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
        networkingv1 "k8s.io/api/networking/v1"
)

type ClusterLister interface {
	GetClusters() ([]Cluster, error)
}

type ClusterManager interface {
	ClusterLister
	GetBlueprint(cluster string, namespace string, name string) (*v1alpha1.Blueprint, error)
	CreateBlueprint(cluster string, blueprint *v1alpha1.Blueprint, policies ...*networkingv1.NetworkPolicy) error
	UpdateBlueprint(cluster string, blueprint *v1alpha1.Blueprint) error
	DeleteBlueprint(cluster string, namespace string, name string) error
}

type ClusterMetadata struct {
	Region string
	Zone   string
}

type Cluster struct {
	Name     string
	Metadata ClusterMetadata
}

// Decode json into runtime.Object, which is a pointer (such as &corev1.ConfigMapList)
func Decode(json string, scheme *runtime.Scheme, object runtime.Object) error {
	decoder := serializer.NewCodecFactory(scheme).UniversalDecoder()
	err := runtime.DecodeInto(decoder, []byte(json), object)
	if err != nil {
		return err
	}
	return nil
}
