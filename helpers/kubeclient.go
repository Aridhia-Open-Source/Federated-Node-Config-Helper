package helpers

import (
	v1 "k8s.io/api/core/v1"
)

// KubeClient defines the abstract contract.
type KubeClient interface {
	CreateSecret(namespace string, secret *v1.Secret) (*v1.Secret, error)
	CreateConfigMap(namespace string, cm *v1.ConfigMap) (*v1.ConfigMap, error)
	CreateNamespace(ns *v1.Namespace) (*v1.Namespace, error)
}
